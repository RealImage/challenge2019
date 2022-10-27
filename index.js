const fs = require('fs')

const csv = require('csv-parser');
// import the functions problem1Solution from solution.js
const solutions = require('./solutions.js');
let tempData = [];

// call the main function to execute the solution
main()
async function main() {
    try {
        console.log('inside main')
        let partnerData = await parsePartnerCSV();
        console.log("partnerData holds data from partner.csv \n", partnerData)
        let inputData = await parseInputCSV();
        console.log("inputData  holds data from input.csv ", inputData)

        // problem1Solution is called to execute solution for problem1 with partnerData and inputData as input
        solutions.problem1Solution(partnerData, inputData)


    }
    catch (error) {
        console.log(error)
    }

}

// parsePartnerCSV function is to read data from partners.csv
function parsePartnerCSV() {
    return new Promise((resolve, reject) => {
        fs.createReadStream(__dirname + '/partners.csv')
            .pipe(csv({ headers: false, skipLines: 1 }))
            .on('data', function (data) {
                try {

                    partnerTemp = {
                        "theaterId": data[0].trim(),
                        "minSlabSize": parseInt(data[1].split('-')[0].trim()),
                        "maxSlabSize": parseInt(data[1].split('-')[1].trim()),
                        "minimumCost": parseInt(data[2].trim()),
                        "costPerGb": parseInt(data[3].trim()),
                        "partnerId": data[4].trim()
                    }
                    tempData.push(partnerTemp)
                    // console.log(partnerTemp)
                }
                catch (err) {
                    console.error("Error parsing csv file : " + err)
                    reject(err)
                }
            })
            .on('end', function () {

                resolve(tempData)
                tempData = [];
            });
    })
}


// parsePartnerCSV function is to read data from input.csv
function parseInputCSV() {
    return new Promise((resolve, reject) => {
        fs.createReadStream(__dirname + '/input.csv')
            .pipe(csv({ headers: false }))
            .on('data', function (data) {
                try {

                    inputTemp = {
                        "deliveryId": data[0].trim(),
                        "gbSize": parseInt(data[1].trim()),
                        "theaterId": data[2].trim(),
                    }
                    tempData.push(inputTemp)
                    // console.log(partnerTemp)
                }
                catch (err) {
                    console.error("Error parsing csv file : " + err)
                    reject(err)
                }
            })
            .on('end', function () {

                resolve(tempData)
                tempData = [];
            });
    })
}