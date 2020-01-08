const partners = require("../data/partners");
const capacities = require("../data/capacities");
const input = require("../input/input");

const writeSecondResult = require("../functions/writeFile").writeSecondResult;
const getMinFare = require("../functions/minCost").getMinCost;

let copiedCapacities = [...capacities];

const sortedInput = input.sort((ip1, ip2) => ip2.contentSize - ip1.contentSize);

let resultArray = [];

sortedInput.map(theatre => {
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

    let found = false;
    let i = 0;

    while (!found) {
      let index = copiedCapacities.findIndex(
        cp => cp.partnerId === sortedPartners[i].partnerId
      );

      if (copiedCapacities[index].capacity >= contentSize) {
        copiedCapacities[index].capacity =
          copiedCapacities[index].capacity - contentSize;

        found = true;

        formattedRow = `${delivery}, "true", ${
          sortedPartners[i].partnerId
        } , ${getMinFare(
          sortedPartners[i].minCost,
          sortedPartners[i].costGB,
          contentSize
        )}`;

        resultArray.push(formattedRow);
      }

      i++;
    }
  } else {
    formattedRow = `${delivery}, "false", "",""`;
    resultArray.push(formattedRow);
  }
});

writeSecondResult(resultArray);
