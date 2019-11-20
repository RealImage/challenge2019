var masterTable = [
    {
        theatreId: "T1",
        sizeSlab: {
            min: 0,
            max: 200
        },
        minCost: 2000,
        costPerGB: 20,
        partnerId: "P1"
    },
    {
        theatreId: "T1",
        sizeSlab: {
            min: 200,
            max: 400
        },
        minCost: 3000,
        costPerGB: 15,
        partnerId: "P1"
    },
    {
        theatreId: "T3",
        sizeSlab: {
            min: 100,
            max: 200
        },
        minCost: 4000,
        costPerGB: 30,
        partnerId: "P1"
    },
    {
        theatreId: "T3",
        sizeSlab: {
            min: 200,
            max: 400
        },
        minCost: 5000,
        costPerGB: 25,
        partnerId: "P1"
    },
    {
        theatreId: "T5",
        sizeSlab: {
            min: 100,
            max: 200
        },
        minCost: 2000,
        costPerGB: 30,
        partnerId: "P1"
    },
    {
        theatreId: "T1",
        sizeSlab: {
            min: 0,
            max: 400
        },
        minCost: 1500,
        costPerGB: 25,
        partnerId: "P2"
    }
]

var capacityList = [
    {
        partnerId: "P1",
        capacity: 500
    },
    {
        partnerId: "P2",
        capacity: 300
    }
]

function Statement_1(problemInput) {
    var output = [];
    for (let i = 0; i < problemInput.length; i++) {    
        var deliveryId = problemInput[i]['deliveryId'];
        var theatreId = problemInput[i]['theatreId'];
        var sizeOfDelivery = problemInput[i]['sizeOfDelivery'];
        var maxNumber = Number.MAX_VALUE;
        var outputObj = {};
        outputObj['deliveryId'] = deliveryId;
        outputObj['partnerId'] = "";
        for (let j = 0; j < masterTable.length; j++) {        
            if(masterTable[j].theatreId == theatreId && (sizeOfDelivery > masterTable[j].sizeSlab.min && sizeOfDelivery < masterTable[j].sizeSlab.max)) {
                var deliveryCost = sizeOfDelivery*masterTable[j].costPerGB;           
                if(maxNumber > deliveryCost) {
                    maxNumber = deliveryCost;
                    maxNumber = masterTable[j].minCost > maxNumber ? masterTable[j].minCost : maxNumber;
                    outputObj['partnerId'] = masterTable[j].partnerId;
                }
            }        
        }
        outputObj['deliverable'] = maxNumber === Number.MAX_VALUE ? false : true;
        outputObj['costOfDelivery'] = maxNumber === Number.MAX_VALUE ? "" : maxNumber;
        output.push(outputObj);
    }
    return output;
}

function Statement_2(problemInput) {
    var output = [];
    for (let i = 0; i < problemInput.length; i++) {    
        var deliveryId = problemInput[i]['deliveryId'];
        var theatreId = problemInput[i]['theatreId'];
        var sizeOfDelivery = problemInput[i]['sizeOfDelivery'];
        var maxNumber = Number.MAX_VALUE;
        var outputObj = {};
        outputObj['deliveryId'] = deliveryId;
        outputObj['partnerId'] = "";
        for (let j = 0; j < capacityList.length; j++) {
            for (let k = 0; k < masterTable.length; k++) {        
                if(masterTable[k].theatreId == theatreId 
                    && (sizeOfDelivery > masterTable[k].sizeSlab.min && sizeOfDelivery < masterTable[k].sizeSlab.max)
                    && capacityList[j].capacity >= sizeOfDelivery
                    && capacityList[j].partnerId == masterTable[k].partnerId) {
                    var deliveryCost = sizeOfDelivery*masterTable[k].costPerGB;           
                    if(maxNumber > deliveryCost) {
                        maxNumber = deliveryCost;
                        maxNumber = masterTable[k].minCost > maxNumber ? masterTable[k].minCost : maxNumber;
                        outputObj['partnerId'] = masterTable[k].partnerId;
                    }
                }        
            }        
        } 
        capacityList.map((c) => {
            if(c.partnerId === outputObj.partnerId) {
                c.capacity = c.capacity - sizeOfDelivery
            }
            return c;
        })   
        outputObj['deliverable'] = maxNumber === Number.MAX_VALUE ? false : true;
        outputObj['costOfDelivery'] = maxNumber === Number.MAX_VALUE ? "" : maxNumber;
        output.push(outputObj);
    }
    return output;
}

var deliveryInput = [
    {
        deliveryId: "D1",
        sizeOfDelivery: 100,
        theatreId: "T1"
    },
    {
        deliveryId: "D2",
        sizeOfDelivery: 300,
        theatreId: "T1"
    },
    {
        deliveryId: "D3",
        sizeOfDelivery: 350,
        theatreId: "T1"
    }
]

var statementOneResult = Statement_1(deliveryInput);
console.log('=======Statement One Result========');
console.log(statementOneResult);
console.log('====================================');

var capacityInput = [
    {
        deliveryId: "D1",
        sizeOfDelivery: 100,
        theatreId: "T1"
    },
    {
        deliveryId: "D2",
        sizeOfDelivery: 240,
        theatreId: "T1"
    },
    {
        deliveryId: "D3",
        sizeOfDelivery: 260,
        theatreId: "T1"
    }
]

capacityInput = capacityInput.sort(function(a, b) {
    return b['sizeOfDelivery'] - a['sizeOfDelivery'];
})

var statementTwoResult = Statement_2(capacityInput);
console.log('=======Statement Two Result========');
console.log(statementTwoResult);
console.log('====================================');