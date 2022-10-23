
const csv = require("csv-parser");

const fs = require("fs");

//stores the partner data
let results = [];
//store the input data
let input = [];
//stores the capacity data
let capacities = [];

//This is used to parse the partner csv file into a array of object
/*
[{theater,startingslab,endingslab,minimumCost,costPerGB,partnerId},{...},{...}]
*/
function partnerParse(filename) {
  return new Promise((resolve, reject) => {
    fs.createReadStream(filename)
      .pipe(csv({ headers: false, skipLines: 1 }))
      .on("data", (data) => {
        try {
          let slabSizeArray = data["1"].split("-");
          results.push({
            Theater: data["0"].trim(),
            startingSlab: parseInt(slabSizeArray[0].trim()),
            endingSlab: parseInt(slabSizeArray[1].trim()),
            minimumCost: parseInt(data[2].trim()),
            costPerGB: parseInt(data[3].trim()),
            partnerID: data[4].trim(),
          });
        } catch (error) {
          console.log("error while parsing ", error);
        }
      })
      .on("end", () => {
        resolve(results);
        // console.log("results", results);
      })
      .on("error", (err) => {
        reject(err);
      });
  });
}

//This is used to parse the input csv file into a array of object
/*
[{delivery,sizeUsed,theater},{...},{...}]
*/
function inputParse(filename) {
  return new Promise((resolve, reject) => {
    fs.createReadStream(filename)
      .pipe(csv({ headers: false }))
      .on("data", (data) => {
        try {
          input.push({
            delivery: data["0"].trim(),
            sizeUsed: parseInt(data["1"].trim()),
            theater: data["2"].trim(),
          });
        } catch (error) {
          console.log("error while parsing ", error);
        }
      })
      .on("end", () => {
        resolve(input);
        // console.log("results", results);
      })

      .on("error", (err) => {
        reject(err);
      });
  });
}

//This is used to parse the capacity csv file into a array of object
/*
[{partnerID,capacity},{...},{...}]
*/
function capacityParse(filename) {
  return new Promise((resolve, reject) => {
    fs.createReadStream(filename)
      .pipe(csv({ headers: false, skipLines: 1 }))
      .on("data", (data) => {
        try {
          // console.log("data", data);
          capacities.push({
            partnerID: data["0"].trim(),
            capacity: parseInt(data["1"].trim()),
          });
          console.log("capacities", capacities);
        } catch (error) {
          console.log("error while parsing ", error);
        }
      })
      .on("end", () => {
        resolve(capacities);
        // console.log("results", results);
      })

      .on("error", (err) => {
        reject(err);
      });
  });
}
//export the module that are in need for problem 1 and problem 2
module.exports = {
  partnerParse,
  inputParse,
  capacityParse,
};
