/* Problem statement:
Given a list of content size and Theatre ID, 
Find the partner for each delivery where cost of delivery is minimum. 
If delivery is not possible, mark that delivery impossible.
*/
const parsedData = require('../parsedFiles/parsedFiles');
const fs = require('fs');

async function minimumCost() {

    // Erasing the contents of output1.csv file
    fs.writeFile('../challenge2019/outputs/output1.csv', '', (error) => {
        if (error) throw error;
    });

    const partnersArray = await parsedData.partnersArr();
    const inputArray = await parsedData.inputArr();

    let partnerId = '';
    let deliveryStatus = false;
    let finalCSV = [];
    let finalCost = '';

    // Iterating the input theatres to find the minimun delivery cost
    for (let count = 0; count < inputArray.length; count++) {
        let currentTheatre = partnersArray.filter((value) => value.theatreId == inputArray[count].theatreId);

        for (let index = 0; index < currentTheatre.length; index++) {
            // Checking whether the delievery size falls between the mentioned size slabs
            if ((inputArray[count].deliverySize <= currentTheatre[index].maxSlab && inputArray[count].deliverySize >= currentTheatre[index].minSlab)) {
                deliveryStatus = true;
                let tempCost = (inputArray[count].deliverySize * currentTheatre[index].costPerGb) > (currentTheatre[index].minCost) ? (inputArray[count].deliverySize * currentTheatre[index].costPerGb) : (currentTheatre[index].minCost);
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
        // Pushing each object into an array
        if (count == 0) {
            finalCSV.push(`Challenge 1 Solution \n${inputArray[count].deliveryId},${deliveryStatus},${partnerId},${finalCost}`)
        }
        else {
            finalCSV.push(`\n${inputArray[count].deliveryId},${deliveryStatus},${partnerId != "" ? partnerId : '""'},${finalCost != "" ? finalCost : '""'}`)
        }

        // Resetting the variables.
        finalCost = "";
        partnerId = "";
        deliveryStatus = false;
    }
    // Appending the final array to output1.csv file
    fs.appendFileSync("../challenge2019/outputs/output1.csv", finalCSV + "\n");
}

module.exports = { minimumCost }