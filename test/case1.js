const partners = require("../data/partners");
const input = require("../input/input");

const writeFirstResult = require("../functions/writeFile").writeFirstResult;
const getMinCost = require("../functions/minCost").getMinCost;

let resultArray = [];

input.map(theatre => {
  let { theatreId, contentSize, delivery } = theatre;

  let possiblePartners = partners.filter(
    partner => partner.theatre === theatreId
  );

  let isDeliveryPossible = possiblePartners.some(
    partner =>
      partner.sizeslab[0] <= contentSize && contentSize <= partner.sizeslab[1]
  );

  let formattedRow = "";

  if (isDeliveryPossible) {
    let filteredPartners = possiblePartners.filter(
      partner =>
        partner.sizeslab[0] <= contentSize && contentSize <= partner.sizeslab[1]
    );

    let sortedPartners = filteredPartners.sort(
      (p1, p2) => p1.costGB - p2.costGB
    );

    formattedRow = `${delivery}, "true", ${
      sortedPartners[0].partnerId
    }, ${getMinCost(
      sortedPartners[0].minCost,
      sortedPartners[0].costGB,
      contentSize
    )}`;

    resultArray.push(formattedRow);
  } else {
    formattedRow = `${delivery}, "false", "",""`;
    resultArray.push(formattedRow);
  }
});

writeFirstResult(resultArray);
