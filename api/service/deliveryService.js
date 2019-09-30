const orderBy = require("lodash/orderBy");
const findIndex = require("lodash/findIndex");

function findDeliveryPartner(data, partnerDetails) {
  try {
    const contentSize = parseInt(data["contentSize"]);

    var partnerList = fetchPartnersByTheatreIdAndSlab(
      data["theatreId"],
      contentSize,
      partnerDetails
    );

    return findParnerWithMinConst(partnerList, contentSize, data["deliveryId"]);
  } catch (err) {
    console.log("Error in findDeliveryPartner", err);
  }
}

function fetchPartnersByTheatreIdAndSlab(
  theatreId,
  contentSize,
  partnerDetails
) {
  return partnerDetails.filter(each => {
    return (
      each.theatreId == theatreId &&
      contentSize >= getSlabSize(each["sizeSlab"], 0) &&
      contentSize <= getSlabSize(each["sizeSlab"], 1)
    );
  });
}

function getCapacityOfPartner(capacities, id) {
  return parseInt(
    capacities.filter(each => each["partnerId"] == id)[0]["capacity"] || "0"
  );
}

function fetchPartnersByTheatreIdAndCapacity(
  theatreId,
  contentSize,
  partnerDetails,
  capacities
) {
  return partnerDetails.filter(each => {
    return (
      each.theatreId == theatreId &&
      contentSize >= getSlabSize(each["sizeSlab"], 0) &&
      contentSize <= getSlabSize(each["sizeSlab"], 1) &&
      getCapacityOfPartner(capacities, each["partnerId"]) >= contentSize
    );
  });
}

function getSlabSize(slab, index) {
  return parseInt(slab.split("-")[index]);
}

function allocatePartnerForAllDelivery(deliveries, partners, capacities) {
  deliveries = orderBy(deliveries, ["contentSize"], ["desc"]);
  var response = [];
  deliveries.map(each => {
    const contentSize = parseInt(each["contentSize"]);
    const theatreId = each["theatreId"];
    const deliveryId =  each["deliveryId"];
    const filteredPartner = fetchPartnersByTheatreIdAndCapacity(
      theatreId,
      contentSize,
      partners,
      capacities
    );

    const result = findParnerWithMinConst(
      filteredPartner,
      contentSize,
      deliveryId
    );
    if (result["isTransferable"]) {
      const partner = result["partner"];
      const newCapacity =
        getCapacityOfPartner(capacities, partner) - contentSize;
      capacities = updateCapacity(capacities, partner, newCapacity);
    }
    response.push(result);
  });
  return response;
}

function findParnerWithMinConst(partnerList, contentSize, deliveryId) {
  let min = Number.MAX_SAFE_INTEGER;
  let partnerName = "";
  partnerList.map(each => {
    const currentMin = Math.min(
      min,
      Math.max(
        contentSize * parseInt(each["costPerGB"]),
        parseInt(each["minimumCost"])
      )
    );
    if (min != currentMin) {
      partnerName = each["partnerId"];
      min = currentMin;
    }
  });
  const transferable = min != Number.MAX_SAFE_INTEGER;
  const result = {
    deliveryId: deliveryId,
    isTransferable: transferable,
    partner: partnerName,
    cost: transferable ? min : ""
  };
  return result;
}

function updateCapacity(data, key, newCapacity) {
  var index = findIndex(data, { partnerId: key });
  data.splice(index, 1, { partnerId: key, capacity: newCapacity });
  return data;
}

module.exports = { findDeliveryPartner, allocatePartnerForAllDelivery };
