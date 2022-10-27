const fs = require('fs')

async function problem1Solution(partnerData, inputData) {
    console.log('inside problem1Solution')
    // ref.p1.1 declare a variable to store output
    // ref.p1.2 declare a temporary object to store output data for each input
    // ref.p1.3 loop the input data
    // ref.p1.4 loop the partner data
    // ref.p1.5 assign delivery id to the output object
    // ref.p1.6 check to get the partner data for the current theater id
    // ref.p1.7 check to get the input gb size lies between the partner min and max slab size and set delivery status as true
    // ref.p1.8 find the product of input gb size and cost per gb in current partner data
    // ref.p1.9 if the product is less than the current partner's minimum cost assign the input cost to minimum cost else assign the product
    // ref.p1.10 find then minimum cost found for the input and set it to a temporary object
    // ref.p1.11 push the temporary object to output1 array
    // ref.p1.12 reset the temporary object
    // ref.p1.13 create a new csv output1.csv
    // ref.p1.14 push the output1array to output1.csv


    // ref.p1.1 
    let output1 = [];
    // ref.p1.2
    let tempObject = {
        deliveryId: '',
        deliveryStatus: '',
        finalCost: -1,
        partnerId: '',
    }


    //ref.p1.3
    for (let inputIndex = 0; inputIndex < inputData.length; inputIndex++) {
        //ref.p1.4
        for (let parentIndex = 0; parentIndex < partnerData.length; parentIndex++) {
            // ref.p1.5
            tempObject.deliveryId = inputData[inputIndex].deliveryId
            // ref.p1.6,ref.p1.7
            if (inputData[inputIndex].theaterId == partnerData[parentIndex].theaterId &&
                (inputData[inputIndex].gbSize >= partnerData[parentIndex].minSlabSize && inputData[inputIndex].gbSize <= partnerData[parentIndex].maxSlabSize
                )) {
                tempObject.deliveryStatus = true
                // ref.p1.8, ref.p1.9
                totalCost = inputData[inputIndex].gbSize * partnerData[parentIndex].costPerGb > partnerData[parentIndex].minimumCost
                    ? inputData[inputIndex].gbSize * partnerData[parentIndex].costPerGb
                    : partnerData[parentIndex].minimumCost;
                // ref.p1.10
                if (tempObject.finalCost == -1) {
                    tempObject.finalCost = totalCost,
                    tempObject.partnerId = partnerData[parentIndex].partnerId
                }
                else if (tempObject.finalCost > totalCost) {
                    tempObject.finalCost = totalCost,
                    tempObject.partnerId = partnerData[parentIndex].partnerId
                }
                // console.log(finalCost)
            }
        }
        console.log(tempObject)
        // ref.p1.11
        if (inputIndex == 0) {
            output1.push(tempObject.deliveryId, tempObject.deliveryStatus, tempObject.partnerId, tempObject.finalCost == -1 ? '' : tempObject.finalCost)
        }
        else {
            output1.push('\n' + tempObject.deliveryId, tempObject.deliveryStatus, tempObject.partnerId, tempObject.finalCost == -1 ? '' : tempObject.finalCost)
        }
        // ref.p1.12
        tempObject.deliveryStatus = false
        tempObject.finalCost = -1
        tempObject.partnerId = ''
    }
    console.log(output1)
    // ref.p1.13
    fs.writeFile("output1.csv", "", (error) => {
        if (error) throw error;
    });
    // ref.p1.14
    fs.appendFileSync("output1.csv", output1 + "\n");
}






// export problem1Solution function
module.exports = {
    problem1Solution
}