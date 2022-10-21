
const parser = require('../parser/parser.js')
const fs = require('fs')

const minimumDeliveryCost = async () => {

    fs.writeFile('../challenge2019/outputs/output1.csv', '', (err) => {
        if (err) throw err;
        console.log('Erased Output1!');
    });

    const partnersData = await parser.GetPartnersDetails()
    const inputData = await parser.getInputDetails()

    let minimumCost = "";
    let partnerID = "";
    let deliveryStatus = false;
    let fileContent = [];

    inputData.forEach((delivery, index) => {
        let currentTheatre = partnersData.filter((value) => (value.theatreID === delivery.theatreID))

        currentTheatre.forEach((currentPartner) => {
            if (delivery.deliverySize <= currentPartner.slabMaxSize && delivery.deliverySize >= currentPartner.slabMinSize) {
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


        if (index == 0) {
            fileContent.push(`Problem Statement 1 Output \n${delivery.deliveryID},${deliveryStatus} ,${partnerID},${minimumCost}`)
        }
        else {
            fileContent.push(`\n${delivery.deliveryID},${deliveryStatus} ,${partnerID != "" ? partnerID : '""'},${minimumCost != "" ? minimumCost : '""'}`)
        }

        minimumCost = "";
        partnerID = "";
        deliveryStatus = false;
    })

    fs.appendFileSync("./outputs/output1.csv", fileContent + "\n");

}



module.exports = {
    minimumDeliveryCost
}

