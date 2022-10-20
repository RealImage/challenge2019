const fs = require('fs');
const csv = require('csv-parser');

/*
Used to parse partner.csv and get the required output format
Input: partner.csv file
Output: Object<Array[Object]>
eg: 
theater = {
    "T1": [{..},{..}]
    "T2": [{..},{..}]
    ...
}
*/
let theater = {}
function parsePartnersCSV() {
    return new Promise((resolve, reject) => {
        fs.createReadStream(__dirname + '/partners.csv')
            .pipe(csv({ headers: false, skipLines: 1 }))
            .on('data', function (data) {
                try {
                    if (theater[data['0'].trim()]) {
                        let slabSizeArray = (data['1']).split("-");
                        theater[data['0'].trim()].push(
                            {
                                "slabSizeMin": parseInt(slabSizeArray[0].trim()),
                                "slabSizeMax": parseInt(slabSizeArray[1].trim()),
                                "minimumCost": parseInt(data[2].trim()),
                                "costPerGB": parseInt(data[3].trim()),
                                "partnerID": data[4].trim()
                            }
                        )
                    } else {
                        let slabSizeArray = (data['1']).split("-");
                        theater[data['0'].trim()] = [{
                            "slabSizeMin": parseInt(slabSizeArray[0].trim()),
                            "slabSizeMax": parseInt(slabSizeArray[1].trim()),
                            "minimumCost": parseInt(data[2].trim()),
                            "costPerGB": parseInt(data[3].trim()),
                            "partnerID": data[4].trim()
                        }]
                    }
                }
                catch (err) {
                    console.error("Error parsing csv file : " + err)
                    reject(err)
                }
            })
            .on('end', function () {
                resolve(theater)
            });
    })
}

/*
Used to parse input.csv and get the required output format
Input: input.csv file
Output: Object<Object>
eg: 
input = {
    "D1": { slabSize: 150, theater: 'T1' }
    "D2": {..}
    ...
}
*/
let input = {}
function parseInputCSV() {
    return new Promise((resolve, reject) => {
        fs.createReadStream(__dirname + '/input.csv')
            .pipe(csv({ headers: false }))
            .on('data', function (data) {
                try {
                    input[data['0']] = {
                        "slabSize": parseInt(data[1].trim()),
                        "theater": data[2].trim(),
                    }
                }
                catch (err) {
                    console.error("Error parsing csv file : " + err)
                    reject(err)
                }
            })
            .on('end', function () {
                resolve(input)
            });
    })
}


let capacity = {}
function parseCapacityCSV() {
    return new Promise((resolve, reject) => {
        fs.createReadStream(__dirname + '/capacities.csv')
            .pipe(csv({ headers: false, skipLines: 1 }))
            .on('data', function (data) {
                try {
                    capacity[data['0'].trim()] = parseInt(data['1'].trim())
                }
                catch (err) {
                    console.error("Error parsing csv file : " + err)
                    reject(err)
                }
            })
            .on('end', function () {
                resolve(capacity)
            });
    })
}

module.exports = {
    parsePartnersCSV,
    parseInputCSV,
    parseCapacityCSV
}