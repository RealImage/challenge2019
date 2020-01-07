let processedOutputArr;
let unprocessed_delivery_items;
let available_capacities;
let temp_available_capacities;

let optimize_processed_items;
let temp_total_cost;

module.exports = function(deliveryArr, partners, capacities){
    try{
        processedOutputArr = [];
        unprocessed_delivery_items = [];
        available_capacities = {};
        temp_available_capacities = {};

        optimize_processed_items = [];
        temp_total_cost = Number.MAX_VALUE;

        capacities
            .forEach(dataObj => {
                available_capacities[dataObj.partner_id] = dataObj.max_capacity;
            });

        deliveryArr
            .forEach((delivery) => {
                let delivery_id = delivery.delivery_id;
                let outputObj = {
                    "delivery_id": delivery_id,
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
                    let mappedFilteredPartners = filteredPartners
                        .map((partner) => {
                            let calculatedCost = partner.cost_per_gb * delivery.size_of_delivery;
                            let estimated_cost = partner.minimum_cost > calculatedCost ? partner.minimum_cost : calculatedCost;
                            return {
                                partner_id: partner.partner_id,
                                total_cost: estimated_cost
                            }
                        });
                    if(mappedFilteredPartners.length == 1){
                        let optimizedPartner = mappedFilteredPartners[0];

                        outputObj.partner_id = optimizedPartner.partner_id;
                        outputObj.cost_of_delivery = optimizedPartner.total_cost;

                        processedOutputArr.push(outputObj);

                        available_capacities[outputObj.partner_id] = available_capacities[outputObj.partner_id] - delivery.size_of_delivery;
                    }else{
                        unprocessed_delivery_items.push({
                            ...delivery,
                            mappedFilteredPartners
                        });
                    }
                }else{
                    processedOutputArr.push(outputObj);
                }
            });

        if(unprocessed_delivery_items.length > 0){
            let outputArr = [];
            temp_available_capacities = { ...available_capacities };
            let not_exceeds_capacity = unprocessed_delivery_items
                .every((delivery) => {
                    let delivery_id = delivery.delivery_id;
                    let outputObj = {
                        "delivery_id": delivery_id,
                        "delivery_possible": true
                    };

                    let optimizedPartner = delivery.mappedFilteredPartners.reduce((res, obj) => (
                        (obj.total_cost < res.total_cost) ? obj : res
                    ));
                    outputObj.partner_id = optimizedPartner.partner_id;
                    outputObj.cost_of_delivery = optimizedPartner.total_cost;

                    outputArr.push(outputObj);

                    temp_available_capacities[outputObj.partner_id] = temp_available_capacities[outputObj.partner_id] - delivery.size_of_delivery;
                    return temp_available_capacities[outputObj.partner_id] >= 0 ? true : false;
                });

            if(not_exceeds_capacity){
                processedOutputArr = processedOutputArr.concat(outputArr);
                available_capacities = temp_available_capacities;
                unprocessed_delivery_items = [];
            }else{
                optimizeExceedItems(0, []);

                processedOutputArr = processedOutputArr.concat(optimize_processed_items);
                available_capacities = temp_available_capacities;
                unprocessed_delivery_items = [];
            }
        }
        processedOutputArr.sort((a, b) => ( (a.delivery_id).localeCompare(b.delivery_id) ));
        return processedOutputArr;
    }catch(err){
        console.log(`ERROR LOG --- processProblem2 --- ${err.stack}`);
    }
}

function optimizeExceedItems(level, combArr){
    try{
        if(level >= unprocessed_delivery_items.length){
            let temp_capacities = {...available_capacities};
            let aggregated_cost = 0;
            let not_exceeds_capacity = combArr.every(obj => {
                aggregated_cost = aggregated_cost + obj.total_cost;
                temp_capacities[obj.partner_id] = temp_capacities[obj.partner_id] - obj.size_of_delivery;
                return temp_capacities[obj.partner_id] >= 0 ? true : false;
            });
            
            if(not_exceeds_capacity && (aggregated_cost < temp_total_cost)){
                optimize_processed_items = combArr.map(data => ({
                    "delivery_id": data.delivery_id,
                    "delivery_possible": true,
                    "partner_id": data.partner_id,
                    "cost_of_delivery": data.total_cost
                }));
                temp_total_cost = aggregated_cost;
                temp_available_capacities = temp_capacities;
            }

            return;
        }
        let unprocessed_delivery_item = unprocessed_delivery_items[level];
        let mappedFilteredPartners = unprocessed_delivery_item.mappedFilteredPartners;
        for(let i = 0; i < mappedFilteredPartners.length; i++){
            let tempCombArr = [...combArr];
            tempCombArr.push({
                delivery_id: unprocessed_delivery_item.delivery_id,
                size_of_delivery: unprocessed_delivery_item.size_of_delivery,
                theatre_id: unprocessed_delivery_item.theatre_id,
                ...mappedFilteredPartners[i]
            })
            optimizeExceedItems(level + 1, tempCombArr);
        }
    }catch(err){
        console.log(`ERROR LOG --- optimizeExceedItems --- ${err.stack}`);
    }
}