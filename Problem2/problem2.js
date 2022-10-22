/* 
Given a list of content size and Theatre ID, 
Assign deliveries to partners in such a way that all deliveries are possible (Higher Priority) 
and overall cost of delivery is minimum (i.e. First make sure no delivery is impossible and 
then minimise the sum of cost of all the delivery). 
If delivery is not possible to a theatre, mark that delivery impossible. 
Take partner capacity into consideration as well.
*/

const parsedData = require('../parsedFiles/parsedFiles');
const fs = require('fs');

async function minimumCapacity() {

    // Erasing the contents of output2.csv file
    fs.writeFile('../challenge2019/outputs/output2.csv', '', (error) => {
        if (error) throw error;
    });

    const partnersArray = await parsedData.partnersArr();
    const inputArray = await parsedData.inputArr();
    const capacitiesArray = await parsedData.capacitiesArr();

    let partnerId = '';
    let deliveryStatus = false;
    let csvData = [];
    let finalCSV = '';
    let finalCost = '';

    //Sorting the input file based on the deliverySize
    inputArray.sort((a, b) => parseInt(b.deliverySize) - parseInt(a.deliverySize));

    for (let count = 0; count < inputArray.length; count++) {
        let currentTheatre = partnersArray.filter((value) => value.theatreId == inputArray[count].theatreId);
        for (let index = 0; index < currentTheatre.length; index++) {
            partnerCapacity = capacitiesArray.filter((value) => value.partnerId == currentTheatre[index].partnerId)[0].capacityInGb;

            // Condition to check if the delivery size falls between the slabs and does not exceed the partner capacity
            if (inputArray[count].deliverySize <= currentTheatre[index].maxSlab && inputArray[count].deliverySize >= currentTheatre[index].minSlab && partnerCapacity >= inputArray[count].deliverySize) {
                deliveryStatus = true;
                let tempCost = (inputArray[count].deliverySize * currentTheatre[index].costPerGb) > currentTheatre[index].minCost ? (inputArray[count].deliverySize * currentTheatre[index].costPerGb) : currentTheatre[index].minCost;

                if (finalCost != "") {
                    if (finalCost > tempCost) {
                        finalCost = tempCost;
                        partnerId = currentTheatre[index].partnerId
                    }
                }
                else {
                    finalCost = tempCost;
                    partnerId = currentTheatre[index].partnerId
                }
            }
        }

        //Checking if delivery is possible and updating the new capacity by subtracting the currently covered capacity
        if (deliveryStatus) {
            let filteredCapArr = capacitiesArray.filter((value) => value.partnerId == partnerId);
            for (let rows = 0; rows < filteredCapArr.length; rows++) {
                filteredCapArr[rows].capacityInGb = filteredCapArr[rows].capacityInGb - inputArray[count].deliverySize;
            }
        }

        // 
        csvData.push(`${inputArray[count].deliveryId},${deliveryStatus},${partnerId != "" ? partnerId : '""'},${finalCost != "" ? finalCost : '""'}`)
        //Resetting the variables
        finalCost = '';
        partnerId = '';
        deliveryStatus = false;

        //Sorting the data in the array and storing it in another variable and appending to the output2.csv file.
        csvData.sort();

        //Resetting each time to get a new row
        finalCSV = '';
        for (let data = 0; data < csvData.length; data++) {
            if (data == 0) {
                finalCSV += `Challenge 2 Solution \n${csvData[data]} \n`
            }
            else {
                finalCSV += `${csvData[data]}\n`
            }
        }
    }
    // Appending the final array to output2.csv file
    fs.appendFileSync("../challenge2019/outputs/output2.csv", finalCSV + "\n");
}

module.exports = { minimumCapacity }

