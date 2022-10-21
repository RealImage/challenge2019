const parser = require('../parser/parser.js')
const fs = require('fs')

const minimumCapacity = async () => {

    fs.writeFile('../challenge2019/outputs/output2.csv', '', (err) => {
        if (err) throw err;
        console.log('Erased Output2!');
    });


    const partnersData = await parser.GetPartnersDetails()
    const inputData = await parser.getInputDetails()
    const capacitiesData = await parser.getCapacitiesDetails()

    let minimumCost = "";
    let partnerID = "";
    let deliveryStatus = false;
    let fileContent = '';
    let output = [];

    inputData.sort(function (a, b) {
        return parseInt(b.deliverySize) - parseInt(a.deliverySize);
    });

    inputData.forEach((delivery, index) => {
        let currentTheatre = partnersData.filter((value) => (value.theatreID === delivery.theatreID))
        currentTheatre.forEach((currentPartner, index) => {
            // get the current partner capacity in Gb
            let currentPartnerCapacity = capacitiesData.filter((value) => value.partnerID === currentPartner.partnerID)[0].capacityInGB;
            // check partner capacity is greater than or equal to delivery.
            if (delivery.deliverySize <= currentPartner.slabMaxSize && delivery.deliverySize >= currentPartner.slabMinSize && currentPartnerCapacity >= delivery.deliverySize) {
                deliveryStatus = true;
                //if total cost comes less than minimum cost, minimum cost will be charged.
                let totalCost = (delivery.deliverySize * currentPartner.costPerGB) > currentPartner.minimumCost ? (delivery.deliverySize * currentPartner.costPerGB) : currentPartner.minimumCost;
                if (minimumCost != "") {
                    if (minimumCost > totalCost) {
                        minimumCost = totalCost
                        partnerID = currentPartner.partnerID
                    }
                } else {
                    minimumCost = totalCost;
                    partnerID = currentPartner.partnerID
                }
            }

        })
        if (deliveryStatus) {
            capacitiesData.filter((value) => value.partnerID === partnerID).forEach((Val) => Val.capacityInGB = Val.capacityInGB - delivery.deliverySize)
        }

        output.push(`${delivery.deliveryID},${deliveryStatus} ,${partnerID != "" ? partnerID : '""'},${minimumCost != "" ? minimumCost : '""'}`);

        minimumCost = "";
        partnerID = "";
        deliveryStatus = false;
    })

    output.sort();
    fileContent = '';
    output.forEach((data, index) => {
        if (index == 0) {
            fileContent += `Problem Statement 2 Output \n${data} \n`
        }
        else {
            fileContent += `${data}\n`
        }
    })


    fs.appendFileSync("./outputs/output2.csv", fileContent + "\n");

}

module.exports = {
    minimumCapacity
}