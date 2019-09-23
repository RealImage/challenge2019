const fs = require('fs');
const fastcsv = require('fast-csv');
const partnerCsv = fs.createReadStream('./partners.csv');
const inputCsv = fs.createReadStream('./input.csv');
const output1Csv = fs.createReadStream('./output1.csv');
let partnersData = [];
let actualResult = [];
let expectedResult = [];
console.log('-----------------------------------------------------------------');
console.log('PROBLEM 1 :');
partnerCsv
    .pipe(fastcsv.parse({ headers: true }))
    .on('data', row => {
        partnersData.push(row);
    })
    .on('end', function () {
        inputCsv
            .pipe(fastcsv.parse({ headers: false }))
            .on('data', row => {
                //find Minimum Cost
                minimumCost(row);
            })
            .on('end', function () {
                console.log('ACTUAL RESULT :');
                console.table(actualResult);
                output1Csv
                    .pipe(fastcsv.parse({headers: false}))
                    .on('data', row => {
                        expectedResult.push(row);
                    })
                    .on('end', () => {
                        console.log('EXPECTED RESULT :');
                        console.table(expectedResult);
                        console.log('-----------------------------------------------------------------');
                    })
            });
    });
/**
 * 
 * @param {Array} input 
 *  1. Filter by theatre id
 *  2. Filter deliverySize(deliverySize available within the range) isSizeSlabAvailable function
 *  3. Calculate Total cost per GB
 *  4. Filter cost that below minimum amount
 *  5. Pick Lower cost from the result
 */

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
        if (
            res.Theatre.trim().toLowerCase() === theatreID.toLowerCase() &&
            isSizeSlabAvailable(res['Size Slab (in GB)'], deliverySize)
        ) {
            //set Delivery is Possible
            finalResult.deliveryPossible = true;
            let totalCost = Number(res['Cost Per GB']) * Number(deliverySize);
            //if total cost is lesser than minimum cost than select minimum cost
            res.totalCost = totalCost < res['Minimum cost'] ? res['Minimum cost'] : String(totalCost);

            if (finalResult.minimumCost) {
                if(res.totalCost < finalResult.minimumCost) {
                    finalResult.partnerId =  res['Partner ID'];
                    finalResult.minimumCost = res.totalCost;
                }
               
            } else {
                finalResult.partnerId = res['Partner ID'];
                finalResult.minimumCost = res.totalCost.trim();
            }
            return res;
        }
    });
    actualResult.push(finalResult);

}

/**
 * 
 * @param {String} selectedSize 
 * @param {String} deliverySize 
 * Helper function
 */

function isSizeSlabAvailable(selectedSize, deliverySize) {
    selectedSize = selectedSize.trim().split('-');
    let smallSize = Math.min(...selectedSize);
    let largeSize = Math.max(...selectedSize);
    return (deliverySize >= smallSize && deliverySize <= largeSize);
}