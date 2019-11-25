/*
Solution for problem statement 2
*/
const csv = require('csv-parser');
const fs = require('fs');
const _ = require('lodash');
const jsonexport = require('jsonexport');

const inputData = require('./inputData.json');
var partnerData = [], outputData = [], capacityData = [], deliveryComparison = [];
try{
fs.createReadStream('capacities.csv')
    .pipe(csv())
    .on('data', (data) => {
        Object.keys(data).map(k => data[k] = data[k].trim()); //  trim the values
        capacityData.push(data);
    })
    .on('end', () => {
        var capacityGroupedData = _.mapValues(_.groupBy(capacityData, 'Partner ID'),
            clist => clist.map(capacityData => _.omit(capacityData, 'Partner ID')));

        fs.createReadStream('partners.csv')
            .pipe(csv())
            .on('data', (data) => {
                Object.keys(data).map(k => data[k] = data[k].trim()); //  trim the values
                partnerData.push(data);
            })
            .on('end', () => {

                var groupedTheatreData = _.mapValues(_.groupBy(partnerData, 'Theatre'),
                    clist => clist.map(partnerData => _.omit(partnerData, 'Theatre'))); // to group all the theatre data sets


                for (var i in inputData) {
                    var theatreData = groupedTheatreData[inputData[i]['theatre']];
                    var deliveryPossible = false;
                    var deliveryCost = [];
                    for (var j in theatreData) { // to find the size slab to calculate cost
                        var partnerDataFlag = rangeCheck(theatreData[j], inputData[i]['totalSize']);
                        if (partnerDataFlag == true) {
                            deliveryPossible = true;
                            var totalCost = theatreData[j]['Cost Per GB'] * inputData[i]['totalSize'];
                            if (totalCost < theatreData[j]['Minimum cost']) { // total cost on computation should be greater than minimum cost if not minimum cost need to be considered
                                totalCost = theatreData[j]['Minimum cost']
                            }
                            deliveryCost.push({
                                "deliveryPartner": inputData[i]['deliveryPartner'],
                                "deliveryPossible": deliveryPossible,
                                "totalCost": totalCost.toString(),
                                "partnerID": theatreData[j]['Partner ID'],
                                "totalSize": inputData[i]['totalSize']
                            })
                        }
                    }
                    if (deliveryPossible == true) {
                        deliveryCost = _.sortBy(deliveryCost, ['user', 'totalCost']); // Sorting by ascending order of cost

                        deliveryComparison.push(deliveryCost) // Getting the least costed partner by taking the first index

                    } else {

                        deliveryComparison.push([{
                            "deliveryPartner": inputData[i]['deliveryPartner'],
                            "deliveryPossible": deliveryPossible,
                            "totalCost": "",
                            "partnerID": ""
                        }])

                    }


                }


                var totalSize = 0, currentPartnerKey = { "partnerId": capacityData[0]['Partner ID'], index: 0 };
                for (var i in deliveryComparison) {

                    if (deliveryComparison[i].length < 2) {

                        if (deliveryComparison[i][0]['deliveryPossible'] == true) {
                            if (totalSize == 0) {
                                totalSize = deliveryComparison[i][0]['totalSize'];
                            } else {
                                totalSize = totalSize + deliveryComparison[i][0]['totalSize'];
                            }
                        }
                        var outputVal=deliveryComparison[i][0];
                        delete outputVal['totalSize']; 
                        outputData.push(outputVal);
                    } else {

                        for (var j in deliveryComparison[i]) {
                            if (totalSize == 0) {
                                totalSize = parseInt(deliveryComparison[i][j]['totalSize']);
                            } else {
                                totalSize = totalSize + parseInt(deliveryComparison[i][j]['totalSize']);
                            }


                            var partnerAssigned = false;
                            while (partnerAssigned == false) {

                                if (totalSize < capacityGroupedData[currentPartnerKey['partnerId']][0]['Capacity (in GB)']) {
                                    var outputVal=deliveryComparison[i][j];
                                    delete outputVal['totalSize']; 
                                    outputData.push(outputVal);
                                    partnerAssigned = true;
                                } else {
                                    var indexVal = currentPartnerKey['index'] + 1;
                                    if (indexVal + 1 < capacityData.length) {
                                        currentPartnerKey['partnerId'] = capacityData[indexVal]['Partner ID'];
                                        currentPartnerKey['index'] = currentPartnerKey['index'] + 1;
                                        totalSize = totalSize - parseInt(deliveryComparison[i][j]['totalSize']);
                                        partnerAssigned = "nextKey";
                                    } else {
                                        outputData.push({
                                            "deliveryPartner": deliveryComparison[i][j]['deliveryPartner'],
                                            "deliveryPossible": false,
                                            "totalCost": "",
                                            "partnerID": ""
                                        });
                                        partnerAssigned = true;
                                    }
                                }
                            }
                            if (partnerAssigned == true) {
                                break;
                            }
                        }
                    }
                }


                jsonexport(outputData, function (err, csv) {
                    if (err) return console.log(err);
                    fs.writeFileSync('./output/output2.csv', csv);
                    console.log("Output generated .. Please goto output -> output2.csv to check the results")
                });

            })
    });
}catch(e){
    console.log("Sorry unable to generate output")
    console.log(e);
}
function rangeCheck(dataSet, totalSize) { // to check if the data falls inside the range for the partner
    try{
    var range = dataSet['Size Slab (in GB)'];
    range = range.split('-');
    //console.log("range",range);
    if (totalSize > range[0] && totalSize < range[1]) {
        return true;
    } else {
        return false;
    }
}catch(e){
    console.log(e);
}
}