const fs = require('fs');
const fastcsv = require('fast-csv');
const partner_csv = fs.createReadStream('./partners.csv');
const input_csv = fs.createReadStream('./input.csv');
const output_csv = fs.createReadStream('./output1.csv');
const ws = fs.createWriteStream('solution_for_problem_1.csv');

let partnersData = []; // initial partners data array
let actualResult = []; // actual result data array
let expectedResult = [];

console.log('PROBLEM 1 :');

partner_csv
  .pipe(fastcsv.parse({ headers: true }))
  .on('data', row => {
    // console.log("INPUT CSV DATA");
    // console.log(row);
    // pushing partner's csv data row into initialised array
    partnersData.push(row);
  })
  .on('end', function() {
    input_csv
      .pipe(fastcsv.parse({ headers: false }))
      .on('data', row => {
        console.log('INPUT CSV DATA');
        console.log(row);
        //find Minimum Cost

        minimumCost(row);
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
            console.log(expectedResult);
            fastcsv
              .write(expectedResult, {
                headers: [
                  'deliveryID',
                  'deliveryPossible',
                  'partnerId',
                  'minimumCost'
                ]
              })
              .pipe(ws);

            console.log(
              '-----------------------------------------------------------------'
            );
          });
      });
  });

function minimumCost(input) {
  let deliveryID = input[0];
  let deliverySize = input[1];
  let theatreID = input[2];

  //pre-defined final result
  let finalResult = {
    deliveryID: deliveryID,
    deliveryPossible: false,
    partnerId: '',
    minimumCost: ''
  };

  partnersData.map(res => {
    // console.log("result set :", res);

    if (
      res.Theatre.trim().toLowerCase() === theatreID.toLowerCase() &&
      isSizeSlabAvailable(res['Size Slab (in GB)'], deliverySize)
    ) {
      //set Delivery is Possible
      finalResult.deliveryPossible = true;
      let totalCost = Number(res['Cost Per GB']) * Number(deliverySize);
      //if total cost is lesser than minimum cost than select minimum cost
      res.totalCost =
        totalCost < res['Minimum cost']
          ? res['Minimum cost']
          : String(totalCost);

      if (finalResult.minimumCost) {
        if (res.totalCost < finalResult.minimumCost) {
          finalResult.partnerId = res['Partner ID'];
          finalResult.minimumCost = res.totalCost;
        }
      } else {
        finalResult.partnerId = res['Partner ID'];
        finalResult.minimumCost = res.totalCost.trim();
      }
      return res;
    }
  });
  console.log('bingo :', finalResult);
  actualResult.push(finalResult);
}

const isSizeSlabAvailable = (selectedSize, deliverySize) => {
  selectedSize = selectedSize.trim().split('-');
  let smallSize = Math.min(...selectedSize);
  let largeSize = Math.max(...selectedSize);
  return deliverySize >= smallSize && deliverySize <= largeSize;
};
