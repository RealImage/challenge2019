const fs = require("fs");

const { inputParse, partnerParse } = require("./parser.js");

//write a async function which calls parnterPatrse,inputParse in problem 1 and store in separate variables

async function problem1() {
  try {
    let result = await partnerParse("partners.csv");
    let input = await inputParse("input.csv");
    //status variable to set the status
    let status = false;
    //output variable to push the output of problem 1
    let output = [];

    //Declare a for loop to iterate through the input array and result array
    for (let i = 0; i < input.length; i++) {
      for (let j = 0; j < result.length; j++) {
        //condition to check the theater is present or not for the input with partner
        /*
        check the sizeUsed is between the slab
        */
        if (
          input[i].theater === result[j].Theater &&
          input[i].sizeUsed >= result[j].startingSlab &&
          input[i].sizeUsed <= result[j].endingSlab
        ) {
          status = true;
          //calculation for cost by multiplying the sizeused with cost per gb and store in variable.
          let cost = input[i].sizeUsed * result[j].costPerGB;

          //if cost is less than minimum cost then assign the minimumCost to cost
          if (cost < result[j].minimumCost) {
            cost = result[j].minimumCost;
          }
          //Push the data to output with push function and assign the partner for every possible partner whose able to deliver
          output.push({
            delivery: input[i].delivery,
            status: status,
            // sizeUsed: input[i].sizeUsed,
            //theater: input[i].theater,
            partnerID: result[j].partnerID,
            cost: cost,
          });
        }
      }
    }
    //declared a for loop to seperate the input that is not in slab and given theater is not in input

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

    //sort the output array based on cost from low to high so we can filter the low cost for the delivery
    output.sort((a, b) => a.cost - b.cost);

    //  remove the duplicate values from the output array so  there will be only one low cost delivery
    output = output.filter(
      (object, index, self) =>
        index === self.findIndex((t) => t.delivery === object.delivery)
    );

    //sort the delivery from d1 - d2 for better view
    output.sort((a, b) => a.delivery[1] - b.delivery[1]);

    //write the finalData to output1.csv file using fs module and csv-writer module after erasing the previous data

    let finalData = output.map((item) => {
      return Object.values(item);
    });

    //emptyb the output1 file
    fs.writeFile("output1.csv", "", function (err) {
      if (err) throw err;
      console.log("File is created successfully.");
    });

    //append the finalData to output1.csv file by leaving a line after each row
    for (let i = 0; i < finalData.length; i++) {
      fs.appendFileSync("output1.csv", finalData[i] + "\n");
    }
  } catch (error) {
    console.log("error while calling partnerParse", error);
  }
}

module.exports = {
  problem1
};
