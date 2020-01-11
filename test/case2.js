const partners = require("../data/partners");
const capacities = require("../data/capacities");
const input = require("../input/input");

// const partners = require("../data/partnerstest");
// const capacities = require("../data/capacitiestest");
// const input = require("../input/test4");

const writeSecondResult = require("../functions/writeFile").writeSecondResult;
const getMinCost = require("../functions/minCost").getMinCost;

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

    let capacityIndex = 0;
    let partnerIndex = 0;
    let assumedMinCost = 0;

    let minCostIndex = 0;
    let minCostFound = false;

    while (!minCostFound) {
      let index = copiedCapacities.findIndex(
        cp => cp.partnerId === filteredPartners[minCostIndex].partnerId
      );

      if (index > -1 && copiedCapacities[index].capacity >= contentSize) {
        assumedMinCost = getMinCost(
          filteredPartners[minCostIndex].minCost,
          filteredPartners[minCostIndex].costGB,
          contentSize
        );

        capacityIndex = index;
        partnerIndex = minCostIndex;
        minCostFound = true;
      }

      minCostIndex++;
    }

    filteredPartners.map((partner, index) => {
      const calculatedMinCost = getMinCost(
        partner.minCost,
        partner.costGB,
        contentSize
      );

      let cpIndex = copiedCapacities.findIndex(
        cp => cp.partnerId === partner.partnerId
      );

      if (cpIndex > -1) {
        if (
          calculatedMinCost < assumedMinCost &&
          copiedCapacities[cpIndex].capacity >= contentSize
        ) {
          assumedMinCost = calculatedMinCost;
          partnerIndex = index;
          capacityIndex = cpIndex;
        }
      }
    });

    copiedCapacities[capacityIndex].capacity =
      copiedCapacities[capacityIndex].capacity - contentSize;
    formattedRow = `${delivery}, "true", ${filteredPartners[partnerIndex].partnerId} , ${assumedMinCost}`;

    resultArray.push(formattedRow);
  } else {
    formattedRow = `${delivery}, "false", "",""`;
    resultArray.push(formattedRow);
  }
});

writeSecondResult(resultArray);
