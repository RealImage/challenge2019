const fs = require("fs");
const { inputParse, partnerParse } = require("./parser.js");

//write a async function which calls the above function and returns the result to index.js

async function problem1() {
  try {
    let result = await partnerParse("partners.csv");
    let input = await inputParse("input.csv");
    let status = false;
    let output = [];
    console.log("helllllll00", input, result);

    //Declare a for loop to iterate through the input array and result array
    for (let i = 0; i < input.length; i++) {
      for (let j = 0; j < result.length; j++) {
        if (
          input[i].theater === result[j].Theater &&
          input[i].sizeUsed >= result[j].startingSlab &&
          input[i].sizeUsed <= result[j].endingSlab
        ) {
          status = true;
          let cost = input[i].sizeUsed * result[j].costPerGB;
          if (cost < result[j].minimumCost) {
            cost = result[j].minimumCost;
          }

          output.push({
            delivery: input[i].delivery,
            status: status,
            // sizeUsed: input[i].sizeUsed,
            //theater: input[i].theater,
            partnerID: result[j].partnerID,
            cost: cost,
          });
          console.log("output1112", output);
        }

        console.log("output22221", output);
      }
    }

    for (let i = 0; i < input.length; i++) {
      for (let j = 0; j < output.length; j++) {
        if (
          input[i].theater !== result[j].Theater &&
          input[i].sizeUsed <= result[j].startingSlab &&
          input[i].sizeUsed >= result[j].endingSlab
        ) {
          //output[j].theater = input[i].theater;
          status = false;
        }
      }
      if (status === false) {
        output.push({
          delivery: input[i].delivery,
          status: status,
          //sizeUsed: input[i].sizeUsed,
          //theater: input[i].theater,
          partnerID: '""',
          cost: '""',
        });
      }
      status = false;
    }

    //sort the output array based on cost
    output.sort((a, b) => a.cost - b.cost);

    //  remove the duplicate values from the output array
    output = output.filter(
      (object, index, self) =>
        index === self.findIndex((t) => t.delivery === object.delivery)
    );
    //write a condition to push the data which is not present in slab range

    output.sort((a, b) => a.delivery[1] - b.delivery[1]);

    //write the finalData to output1.csv file using fs module and csv-writer module after erasing the previous data

    let finalData = output.map((item) => {
      return Object.values(item);
    });

    console.log("finalData", finalData);
    fs.writeFile("output1.csv", "", function (err) {
      if (err) throw err;
      console.log("File is created successfully.");
    });
    // fs.appendFileSync("output1.csv", finalData + "\n\n");
    //append the finalData to output1.csv file by leaving a line after each row
    for (let i = 0; i < finalData.length; i++) {
      fs.appendFileSync("output1.csv", finalData[i] + "\n");
    }

    // console.log("output1", output);
  } catch (error) {
    console.log("error while calling partnerParse", error);
  }
}

module.exports = {
  problem1,
  partnerParse,
  inputParse,
};
