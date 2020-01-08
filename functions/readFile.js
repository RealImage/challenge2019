const fs = require("fs");
const path = require("path");
const csv = require("csv-parser");

const Capacity = require("../model/capacity");
const Partner = require("../model/partner");
const constant = require("../constants");

module.exports.fetchCapacities = () => {
  try {
    const readStream = fs.createReadStream(
      path.resolve(__dirname, "../capacities.csv"),
      "utf8"
    );

    let capacityList = [];

    readStream
      .pipe(csv())
      .on("data", data => {
        try {
          let capacityObj = new Capacity();
          Object.keys(data).forEach(key => {
            let value = data[key].trim();

            switch (key) {
              case constant.partner_id:
                capacityObj.partnerId = value;
                break;

              case constant.capacity:
                capacityObj.capacity = parseInt(value);
                break;
            }
          });
          capacityList.push(capacityObj);
        } catch (e) {
          throw e;
        }
      })
      .on("end", () => {
        fs.writeFile(
          path.resolve(__dirname, "../data/capacities.json"),
          JSON.stringify(capacityList),
          err => {
            if (err) throw err;
          }
        );
      });
  } catch (e) {
    throw e;
  }
};

module.exports.fetchPartners = () => {
  try {
    const readStream = fs.createReadStream(
      path.resolve(__dirname, "../partners.csv"),
      "utf8"
    );

    let partnersList = [];

    readStream
      .pipe(csv())
      .on("data", data => {
        try {
          let partnerObj = new Partner();
          Object.keys(data).forEach(key => {
            let value = data[key].trim();

            switch (key) {
              case constant.theatre:
                partnerObj.theatre = value;
                break;

              case constant.size_slab:
                let size = value.split("-");
                size[0] = parseInt(size[0]);
                size[1] = parseInt(size[1]);
                partnerObj.sizeslab = size;
                break;

              case constant.min_cost:
                partnerObj.minCost = parseInt(value);
                break;

              case constant.cost_per_gb:
                partnerObj.costGB = parseInt(value);

              case constant.partner_id:
                partnerObj.partnerId = value;
            }
          });
          partnersList.push(partnerObj);
        } catch (e) {
          throw e;
        }
      })
      .on("end", () => {
        fs.writeFile(
          path.resolve(__dirname, "../data/partners.json"),
          JSON.stringify(partnersList),
          err => {
            if (err) throw err;
          }
        );
      });
  } catch (e) {
    throw e;
  }
};
