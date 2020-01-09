const partners = require("../data/partners");
const input = require("../input/input");

// const partners = require("../data/partnerstest");
// const input = require("../input/test1");

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

    let assumedMinCost = getMinCost(
      filteredPartners[0].minCost,
      filteredPartners[0].costGB,
      contentSize
    );

    let partnerIndex = 0;

    filteredPartners.map((partner, index) => {
      const calculatedMinCost = getMinCost(
        partner.minCost,
        partner.costGB,
        contentSize
      );

      if (assumedMinCost > calculatedMinCost) {
        assumedMinCost = calculatedMinCost;
        partnerIndex = index;
      }
    });

    formattedRow = `${delivery}, "true", ${filteredPartners[partnerIndex].partnerId}, ${assumedMinCost}`;

    resultArray.push(formattedRow);
  } else {
    formattedRow = `${delivery}, "false", "",""`;
    resultArray.push(formattedRow);
  }
});

writeFirstResult(resultArray);
