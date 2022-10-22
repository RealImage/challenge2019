const csv = require('csv-parser')
const fs = require('fs')

let partnersArray = []
let capacitiesArray = []
let inputArray = []

//Parsing the data from partners.csv
async function partnersArr() {
    return new Promise((resolve, reject) => {
        fs.createReadStream('./parsedFiles/partners.csv')
            .pipe(csv({ headers: false, skipLines: 1 }))
            .on('data', (data) => {
                try {
                    const minRange = data['1'].split('-')[0].trim();
                    const maxRange = data['1'].split('-')[1].trim();
                    partnersArray.push({
                        "theatreId": data[0].trim(),
                        "minSlab": parseInt(minRange),
                        "maxSlab": parseInt(maxRange),
                        "minCost": parseInt(data[2].trim()),
                        "costPerGb": parseInt(data[3].trim()),
                        "partnerId": data[4].trim(),
                    })
                }
                catch (err) {
                    console.error("Error parsing csv file : " + err)
                    reject(err)
                }
            })
            .on('end', () => {
                resolve(partnersArray)
            });
        partnersArray = []
    })
}

//Parsing the data from capacities.csv
async function capacitiesArr() {
    return new Promise((resolve, reject) => {
        fs.createReadStream('./parsedFiles/capacities.csv')
            .pipe(csv({ headers: false, skipLines: 1 }))
            .on('data', (data) => {
                try {
                    capacitiesArray.push({
                        "partnerId": data[0].trim(),
                        "capacityInGb": parseInt(data[1])
                    })
                }
                catch (err) {
                    console.error("Error parsing csv file : " + err)
                    reject(err)
                }
            })
            .on('end', () => {
                resolve(capacitiesArray)
            });
        capacitiesArray = []
    })
}

//Parsing the data from input.csv
async function inputArr() {
    return new Promise((resolve, reject) => {
        fs.createReadStream('./parsedFiles/input.csv')
            .pipe(csv({ headers: false }))
            .on('data', (data) => {
                try {
                    inputArray.push({
                        "deliveryId": data[0].trim(),
                        "deliverySize": parseInt(data[1]),
                        "theatreId": data[2].trim()
                    })
                }
                catch (err) {
                    console.error("Error parsing csv file : " + err)
                    reject(err)
                }
            })
            .on('end', () => {
                resolve(inputArray)
            });
        inputArray = []
    })
}


module.exports = {
    partnersArr,
    capacitiesArr,
    inputArr
}