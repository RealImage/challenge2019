//Import  all the parse functions for data
const { inputParse, partnerParse,capacityParse } = require("./parser.js");
const fs = require("fs");

async function problem2() {
  try {
    // Erasing the contents of output2.csv to add the new add data for clear representation.
    fs.writeFile("output2.csv", "", (error) => {
      if (error) throw error;
    });

    //declare a variable to store the parsed Data of partner file
    const result = await partnerParse("partners.csv");

    //declare a variable to store the parsed Data of input file
    const input = await inputParse("input.csv");
    //declare a variable to store the parsed Data of capacities file
    const capacities = await capacityParse("capacities.csv");
    console.log(result);
    let partnerId = "";
    let Status = false;
    let output = [];
    let finalOutput = "";
    let finalCost = "";

    //Sorting the input file based on the  size used  as gb
    input.sort((a, b) => parseInt(b.sizeUsed) - parseInt(a.sizeUsed));
    //loop the input file data  with input variable which as the data of inputb file
    for (let i = 0; i < input.length; i++) {
      //separate all the input data with filter function
      let seperateTheatre = result.filter(
        (value) => value.Theater == input[i].theater
      );
      // loop the data of seperate data
      for (let j = 0; j < seperateTheatre.length; j++) {
        //seperate the capacity with filter function with partnerId of capacity and seperate theater 
        let capacity = capacities.filter(
          (value) => value.partnerID == seperateTheatre[j].partnerID
        )[0].capacity;
        
        //condition to check  size used is in slab and capacity for that sizeused is greater to assign the partner that as more capacity than size used
        if (
          input[i].sizeUsed <= seperateTheatre[j].endingSlab &&
          input[i].sizeUsed >= seperateTheatre[j].startingSlab &&
          capacity >= input[i].sizeUsed
        ) {
          Status = true;

          // declare a variable to store the cost that is multiply by costPerGb with size used.If cost is above minimum cost it will be set as cost else minimum cost will be assigned as cost
          let cost =
            input[i].sizeUsed * seperateTheatre[j].costPerGB >
            seperateTheatre[j].minimumCost
              ? input[i].sizeUsed * seperateTheatre[j].costPerGB
              : seperateTheatre[j].minimumCost;

          if (finalCost != "") {
            if (finalCost > cost) {
              finalCost = cost;
              partnerId = seperateTheatre[j].partnerID;
            }
          } else {
            finalCost = cost;
            partnerId = seperateTheatre[j].partnerID;
          }
        }
      }

      //Checking if delivery is possible and updating the new capacity by subtracting the currently covered capacity
      if (Status) {
        let filteredCapArr = capacities.filter(
          (value) => value.partnerID == partnerId
        );
        for (let rows = 0; rows < filteredCapArr.length; rows++) {
          filteredCapArr[rows].capacity =
            filteredCapArr[rows].capacity - input[i].sizeUsed;
        }
      }

      //
      output.push(
        `${input[i].delivery},${Status},${partnerId != "" ? partnerId : '""'},${
          finalCost != "" ? finalCost : '""'
        }`
      );
      //Resetting the variables
      finalCost = "";
      partnerId = "";
      Status = false;

      //Sorting the data in the array and storing it in another variable and appending to the output2.csv file.
      output.sort();
      //Resetting each time to get a new row
      finalOutput = "";
      for (let data = 0; data < output.length; data++) {
        if (data == 0) {
          finalOutput += `${output[data]} \n`;
        } else {
          finalOutput += `${output[data]}\n`;
        }
      }
    }
    // Appending the final array to output2.csv file
    fs.appendFileSync("output2.csv", finalOutput + "\n");
  }
   
  catch (err) {
    console.log("error while calling problem2", error);
  }
}

module.exports = { problem2 };
