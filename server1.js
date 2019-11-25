/*
Solution for problem statement 1
*/
const csv = require('csv-parser');
const fs = require('fs');
const _ = require('lodash');
const jsonexport = require('jsonexport');

const inputData = require('./inputData.json');
var partnerData = [];
var outputData = [];
try{
fs.createReadStream('partners.csv')
    .pipe(csv())
    .on('data', (data) => {
        Object.keys(data).map(k => data[k] = data[k].trim()); //  trim the values
        partnerData.push(data);
    })
    .on('end', () => {

        var groupedTheatreData = _.mapValues(_.groupBy(partnerData, 'Theatre'),
            clist => clist.map(partnerData => _.omit(partnerData, 'Theatre'))); // to group all the partner data sets
        //console.log("groupedTheatreData",groupedTheatreData)

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
                        "partnerID": theatreData[j]['Partner ID']

                    })
                }
            }
            if (deliveryPossible == true) {
                deliveryCost = _.sortBy(deliveryCost, ['user', 'totalCost']); // Sorting by ascending order of cost
            //    console.log("deliverCost", deliveryCost);
                outputData.push(deliveryCost[0]) // Getting the least costed partner by taking the first index

            } else {

                outputData.push({
                    "deliveryPartner": inputData[i]['deliveryPartner'],
                    "deliveryPossible": deliveryPossible,
                    "totalCost": "",
                    "partnerID": ""
                })

            }
        }
        //console.log("outputData************",outputData);
        jsonexport(outputData, function (err, csv) {
            if (err) return console.log(err);
            fs.writeFileSync('./output/output1.csv', csv);
            console.log("Output generated .. Please goto output -> output1.csv to check the results")
        });

    })
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