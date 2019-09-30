const { findDeliveryPartner,allocatePartnerForAllDelivery } = require("./deliveryService");
const fastcsv = require("fast-csv");
const fs = require("fs");

function createOutputFile(data,fileName){
    fastcsv.writeToPath('./assets/'+fileName,data)
      .on('error',(err)=>console.log(err))
      .on('finish',()=>console.log('Sucess'))
  }

  function getPartnersDetails() {
    var partnerDetails = [];
    return new Promise(function(resolve, reject) {
      fs.createReadStream("./assets/partners.csv")
        .pipe(
          fastcsv.parse({
            headers: [
              "theatreId",
              "sizeSlab",
              "minimumCost",
              "costPerGB",
              "partnerId"
            ],
            renameHeaders: true,
            trim: true
          })
        )
        .on("data", function(row, err) {
          if (err) console.log("partners file read", err);
          partnerDetails.push(row);
        })
        .on("end", ()=>resolve(partnerDetails));
    });
  }
  
 
  
  
  function getCapacities(){
    var capacities = [];
    return new Promise(function(resolve, reject) {
      fs.createReadStream("./assets/capacities.csv")
        .pipe(
          fastcsv.parse({
            headers: [
              "partnerId",
              "capacity"
            ],
            renameHeaders: true,
            trim:true
          })
        )
        .on("data", function(row, err) {
          if (err) console.log("capacity file read", err);
          capacities.push(row);
        })
        .on("end", ()=>resolve(capacities));
        
    });
  }
  
  function problemStateMent_1(partnerDetails) {
    var resp = [];
    return new Promise(function(resolve, reject) {
      fs.createReadStream("./assets/input.csv")
        .pipe(
          fastcsv.parse({ headers: ["deliveryId", "contentSize", "theatreId"],trim:true })
        )
        .on("data", function(row, err) {
          if (err) console.log("in input file read", err);
           const result = findDeliveryPartner(row, partnerDetails);
          resp.push(result);
        })
        .on("end", ()=>resolve(resp));
    });
  }
  
function problemStateMent_2(partners,capacities){
    var inputFile = [];
    var response = {};

    return new Promise(function(resolve, reject) {
      fs.createReadStream("./assets/input.csv")
        .pipe(
          fastcsv.parse({ headers: ["deliveryId", "contentSize", "theatreId"],trim:true })
        )
        .on("data", function(row, err) {
          if (err) console.log("in input file read", err);
          inputFile.push(row);
        })
        .on('finish',() => {
          response =  allocatePartnerForAllDelivery(inputFile,partners,capacities);
        })
        .on('end',()=>resolve(response));
    });
  }

  module.exports = {
    problemStateMent_2,
    getCapacities,
    problemStateMent_1,
    getPartnersDetails,
    createOutputFile
  };