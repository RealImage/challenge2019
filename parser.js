const csv = require("csv-parser");

const fs = require("fs");
let results = [];
let input = [];
let capacities = [];
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

module.exports = {
  partnerParse,
  inputParse,
  capacityParse,
};
