const csv = require('csv-parser')
const fs = require('fs')

let PartnerList = []
let InputList = []
let capacitiesList = []

//Parsing the Partners CSV File
const GetPartnersDetails = async () => {
    return new Promise((resolve, reject) => {
        fs.createReadStream('./inputs/partners.csv')
            .pipe(csv({ headers: false, skipLines: 1 }))
            .on('data', (data) => {
                try {
                    createPartnersList(data);
                }
                catch (err) {
                    console.error("Error parsing csv file : " + err)
                    reject(err)
                }
            })
            .on('end', () => {
                resolve(PartnerList)
            });
            PartnerList =[]
    })
}


//Spliting up the data by Key and Value for partners
const createPartnersList = (data) => {
    PartnerList.push({
        "theatreID": data[0].trim(),
        "slabMinSize": parseInt(data['1'].split('-')[0].trim()),
        "slabMaxSize": parseInt(data['1'].split('-')[1].trim()),
        "minimumCost": parseInt(data[2].trim()),
        "costPerGB": parseInt(data[3].trim()),
        "partnerID": data[4].trim(),
    })
}


//Parsing the Input CSV File 
const getInputDetails = async () => {
    return new Promise((resolve, reject) => {
        fs.createReadStream('./inputs/input.csv')
            .pipe(csv({headers:false}))
            .on('data',(data) => {
                try {
                    createInputList(data)
                }
                catch (err) {
                    console.error("Error parsing csv file : " + err)
                    reject(err)
                }
            })
            .on('end', () => {
                resolve(InputList)
            });
            InputList = []
    })
}


//Spliting up the data by Key and Value for input
const createInputList = (data) => {
    InputList.push({
        "deliveryID":data[0].trim(),
        "deliverySize": parseInt(data[1].trim()),
        "theatreID": data[2].trim(),
    })
}

//Parsing the capacities CSV File 
const getCapacitiesDetails = async () => {
    return new Promise((resolve, reject) => {
        fs.createReadStream('./inputs/capacities.csv')
            .pipe(csv({headers:false, skipLines: 1}))
            .on('data',(data) => {
                try {
                    createCapacitiesList(data)
                }
                catch (err) {
                    console.error("Error parsing csv file : " + err)
                    reject(err)
                }
            })
            .on('end', () => {
                resolve(capacitiesList)
            });
            capacitiesList = []
    })
}


//Spliting up the data by Key and Value for capacities
const createCapacitiesList = (data) => {
    capacitiesList.push({
        "partnerID":data[0].trim(),
        "capacityInGB": parseInt(data[1].trim()),
    })
}



//Exporting all the Functions
module.exports = {
    GetPartnersDetails,
    getInputDetails,
    getCapacitiesDetails
}

