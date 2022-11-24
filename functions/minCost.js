module.exports.getMinCost = (minCost, costperGB, contentSize) => {
  let calculatedCost = contentSize * costperGB;
  return calculatedCost > minCost ? calculatedCost : minCost;
};
