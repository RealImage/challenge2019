const { parsePartnersCSV, parseInputCSV } = require('./index.js')
var fs = require('fs');

async function minimumCostOfDilivery() {
    const theater = await parsePartnersCSV();
    const input = await parseInputCSV();
    //iterate each delivery
    for (let delevery in input) {
        let deliverySlabSize = input[delevery].slabSize
        let deliveryTheater = input[delevery].theater
        let deliveryStatus = false;
        let deliveryPartner = "";
        let minimumDeliveryCost = "";
        //iterate all partner to get the right delivery value
        for (let partner in theater[deliveryTheater]) {
            //Condition to check if delivery slab size lies between partner slab size range
            if (deliverySlabSize <= theater[deliveryTheater][partner].slabSizeMax &&
                deliverySlabSize >= theater[deliveryTheater][partner].slabSizeMin) {
                let cost = deliverySlabSize * theater[deliveryTheater][partner].costPerGB
                if (cost <= theater[deliveryTheater][partner].minimumCost) {
                    cost = theater[deliveryTheater][partner].minimumCost
                }
                if (minimumDeliveryCost) {
                    if (minimumDeliveryCost > cost) {
                        minimumDeliveryCost = Math.min(cost, minimumDeliveryCost)
                        deliveryPartner = theater[deliveryTheater][partner].partnerID
                    }
                } else {
                    minimumDeliveryCost = cost;
                    deliveryPartner = theater[deliveryTheater][partner].partnerID
                }
                deliveryStatus = true;
            }
        }
        deliveryPartner = deliveryPartner ? deliveryPartner : '""';
        minimumDeliveryCost = minimumDeliveryCost ? minimumDeliveryCost : '""';
        let finalData = delevery + "," + deliveryStatus + "," + deliveryPartner + "," + minimumDeliveryCost;
        fs.appendFileSync("output.csv", finalData + "\n");   //output is stored in output.csv file
    }
    fs.appendFileSync("output.csv", "\n")
}

minimumCostOfDilivery() 