//import the fs
const fs = require("fs");
// import partnerParse function from problem1.js
const { partnerParse, inputParse, capacityParse } = require("./problem1.js");
//import inputParse function from problem1.js

// create a async function which calls the above functions and returns the result to index.js
async function problem2() {
  try {
    let result = await partnerParse("partners.csv");
    let input = await inputParse("input.csv");
    let capacitydata = await capacityParse("capacities.csv");
    let status = false;
    let output = [];
    // console.log("capacitydata", capacitydata);
    //sort the input array based on sizeUsed
    input.sort((a, b) => b.sizeUsed - a.sizeUsed);
    // console.log("input", input);

    //Declare a for loop to iterate through the input array and result array
    for (let i = 0; i < input.length; i++) {
      for (let j = 0; j < result.length; j++) {
        //create a let variable to store the capacity of the partner and minus the sizeUsed from it

        //

        if (
          input[i].theater === result[j].Theater &&
          input[i].sizeUsed >= result[j].startingSlab &&
          input[i].sizeUsed <= result[j].endingSlab
          //  capacitydata[j].capacity >= input[i].sizeUsed
        ) {
          status = true;
          let cost = input[i].sizeUsed * result[j].costPerGB;
          if (cost < result[j].minimumCost) {
            cost = result[j].minimumCost;
          }

          output.push({
            delivery: input[i].delivery,
            status: status,
            sizeUsed: input[i].sizeUsed,
            theater: input[i].theater,
            partnerID: result[j].partnerID,
            cost: cost,
          });
        }
      }
    }

    // capacity.sort((a, b) => b.capacities - a.capacities);
    // console.log("capacity", capacity);

    // // declare a for loop to iterate through the output array and capacity array and check the capacity of the partner
    // for (let i = 0; i < output.length; i++) {
    //     for (let j = 0; j < capacity.length; j++) {

    //     }
    // }

    //    console.log("output", output);
    //    console.log("capacitydata11", capacitydata);
    // create a for loop to iterate output and capacity array and check the capacity of the partner and minus the sizeUsed from it and push the result to output array

    //console.log("output111", output);

    //sort the output array based on cost
    output.sort((a, b) => a.cost - b.cost);

    //  remove the duplicate values from the output array
    output = output.filter(
      (object, index, self) =>
        index === self.findIndex((t) => t.delivery === object.delivery)
    );
    //write a condition to push the data which is not present in slab range

    output.sort((a, b) => a.delivery[1] - b.delivery[1]);
    console.log("output", output);

    //write the finalData to output1.csv file using fs module and csv-writer module after erasing the previous data

    let finalData = output.map((item) => {
      return Object.values(item);
    });

    // console.log("finalData", finalData);
    fs.writeFile("output2.csv", "", function (err) {
      if (err) throw err;
      console.log("File is created successfully.");
    });
    // fs.appendFileSync("output1.csv", finalData + "\n\n");
    //append the finalData to output1.csv file by leaving a line after each row
    for (let i = 0; i < finalData.length; i++) {
      fs.appendFileSync("output2.csv", finalData[i] + "\n");
    }

    //console.log("capacities", capacity);
  } catch (error) {
    console.log("error while calling problem 2", error);
  }
}

module.exports = { problem2 };
