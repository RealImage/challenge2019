module.exports = function(deliveryArr, partners){
    try{
        let outputArr = [];
        deliveryArr
            .forEach((delivery) => {
                let outputObj = {
                    "delivery_id": delivery.delivery_id,
                    "delivery_possible": false,
                    "partner_id": "",
                    "cost_of_delivery": null
                };
                let filteredPartners = partners
                    .filter(data => (
                        (data.theatre_id == delivery.theatre_id) &&
                        (
                            (data.min_size <= delivery.size_of_delivery) && 
                            (delivery.size_of_delivery <= data.max_size)
                        )
                    ));

                if(filteredPartners.length > 0){
                    outputObj.delivery_possible = true;
                    let optimizedPartner = filteredPartners
                        .map((partner) => {
                            let calculatedCost = partner.cost_per_gb * delivery.size_of_delivery;
                            let estimated_cost = partner.minimum_cost > calculatedCost ? partner.minimum_cost : calculatedCost;
                            return {
                                partner_id: partner.partner_id,
                                total_cost: estimated_cost
                            }
                        })
                        .reduce((res, obj) => (
                            (obj.total_cost < res.total_cost) ? obj : res
                        ));
                    outputObj.partner_id = optimizedPartner.partner_id;
                    outputObj.cost_of_delivery = optimizedPartner.total_cost;
                }
                outputArr.push(outputObj);
            });
        return outputArr;
    }catch(err){
        console.log(`ERROR LOG --- processProblem1 --- ${err.stack}`);
    }
}