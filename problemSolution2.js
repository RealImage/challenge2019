const { parsePartnersCSV, parseInputCSV, parseCapacityCSV } = require('./parseCSVFiles')
var fs = require('fs');

async function maximumCapacity() {
    let final_result = []
    const theater = await parsePartnersCSV();
    const input = await parseInputCSV();
    let sortSlabSize = []
    //Descending sort to delivery based on slab size required
    for (let key in input) {
        sortSlabSize.push([key, input[key].slabSize, input[key].theater])
    }
    sortSlabSize.reverse()
    const capacity = await parseCapacityCSV();
    //Iterate each delivery
    for (let delevery of sortSlabSize) {
        let deliverySlabSize = delevery[1]
        let deliveryTheater = delevery[2]
        let deliveryStatus = false;
        let deliveryPartner = "";
        let minimumDeliveryCost = "";
        //Iterate all partner to get the right delivery value
        for (let partner in theater[deliveryTheater]) {
            //Condition to check if delivery slab size lies between partner slab size range and within partner capacity
            if (deliverySlabSize <= theater[deliveryTheater][partner].slabSizeMax &&
                deliverySlabSize >= theater[deliveryTheater][partner].slabSizeMin &&
                deliverySlabSize <= capacity[theater[deliveryTheater][partner].partnerID]) {
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
        capacity[deliveryPartner] = capacity[deliveryPartner] - deliverySlabSize;
        deliveryPartner = deliveryPartner ? deliveryPartner : '""';
        minimumDeliveryCost = minimumDeliveryCost ? minimumDeliveryCost : '""';
        final_result.push(delevery[0] + "," + deliveryStatus + "," + deliveryPartner + "," + minimumDeliveryCost);

    }
    //Ascending sort to get the right ouput format 
    for (let data of final_result.sort()) {
        fs.appendFileSync("output.csv", data + "\n");   //output is stored in output.csv file
    }
    fs.appendFileSync("output.csv", "\n");
}

maximumCapacity()