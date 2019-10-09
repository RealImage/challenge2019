const fs = require('fs');
const fastcsv = require('fast-csv');
const partner_csv = fs.createReadStream('./partners.csv');
const capacity_csv = fs.createReadStream('./capacities.csv');
const input_csv = fs.createReadStream('./input.csv');
const output_csv = fs.createReadStream('./output2.csv');
let partnersData = [];
let capacityData = [];
let actualResult = [];
let expectedResult = [];

console.log('PROBLEM 2 ::::::::::');
capacity_csv.pipe(fastcsv.parse({ headers: true })).on('data', row => {
  capacityData.push(row);
});
partner_csv
  .pipe(fastcsv.parse({ headers: true }))
  .on('data', row => {
    partnersData.push(row);
  })
  .on('end', function() {
    input_csv
      .pipe(fastcsv.parse({ headers: false }))
      .on('data', row => {
        //find Minimum Capacity
        minimumCapacity(row);
      })
      .on('end', function() {
        console.log('ACTUAL RESULT :');
        console.table(actualResult);
        output_csv
          .pipe(fastcsv.parse({ headers: false }))
          .on('data', row => {
            expectedResult.push(row);
          })
          .on('end', () => {
            console.log('EXPECTED RESULT :');
            console.table(expectedResult);
            console.log(
              '-----------------------------------------------------------------'
            );
          });
      });
  });

function minimumCapacity(input) {
  // console.log(capacityData);
  let deliveryID = input[0];
  let deliverySize = input[1];
  let theatreID = input[2];
  let selectedPartner;
  //pre-defined final result
  let finalResult = {
    deliveryID: deliveryID,
    deliveryPossible: false,
    partnerId: '',
    minimumCost: ''
  };
  partnersData.map(res => {
    if (
      res.Theatre.trim().toLowerCase() === theatreID.toLowerCase() &&
      isSizeSlabAvailable(res['Size Slab (in GB)'], deliverySize)
    ) {
      selectedPartner = capacityData.find(
        data =>
          data['Partner ID'].trim().toLowerCase() ==
          res['Partner ID'].trim().toLowerCase()
      );

      //set Delivery is Possible
      finalResult.deliveryPossible = true;
      let totalCost = Number(res['Cost Per GB']) * Number(deliverySize);
      res.totalCost =
        totalCost < res['Minimum cost']
          ? res['Minimum cost']
          : String(totalCost);
      if (Number(selectedPartner['Capacity (in GB)']) >= Number(deliverySize)) {
        if (finalResult.minimumCost) {
          if (res.totalCost < finalResult.minimumCost) {
            finalResult.partnerId = res['Partner ID'];
            finalResult.minimumCost = res.totalCost;
          }
        } else {
          finalResult.partnerId = res['Partner ID'];
          finalResult.minimumCost = res.totalCost.trim();
        }
      }

      return res;
    }
  });
  capacityData.map(res => {
    if (res['Partner ID'].trim() === finalResult.partnerId) {
      res['Capacity (in GB)'] =
        Number(res['Capacity (in GB)']) - Number(deliverySize);
    }
    return res;
  });
  // console.log(capacityData);
  actualResult.push(finalResult);
}

function isSizeSlabAvailable(selectedSize, deliverySize) {
  selectedSize = selectedSize.trim().split('-');
  let smallSize = Math.min(...selectedSize);
  let largeSize = Math.max(...selectedSize);
  return deliverySize >= smallSize && deliverySize <= largeSize;
}
