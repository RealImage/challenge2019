const express = require("express");
const path = require("path");
const app = express();
const _ = require("lodash");
const cors = require("cors");
const bodyParser = require("body-parser");
const { createOutputFile,getCapacities,getPartnersDetails,problemStateMent_1,problemStateMent_2} = require('./service/contentManagerService');

app.use("/assets", express.static("assets", { redirect: false }));

app.use(cors());

app.use(bodyParser.json());
app.use(
  bodyParser.urlencoded({
    extended: true
  })
);
app.use(
  bodyParser.urlencoded({
    extended: true
  })
);

app.get('/',(req,res)=>{
  res.send("\"/problem1\" -> To run problem statement 1 <br> </br> \"/problem2\" -> To run problem statement 2");
})

app.get("/problem1", (req, res) => {
  getPartnersDetails()
    .then((partner) =>{
      problemStateMent_1(partner)
        .then(data => { createOutputFile(data,'output1.csv'); res.send(data)})
        .catch(err => console.log("ERROR in problem statement 1", err))
    })
    .catch((err) => console.log("ERROR in getting partner details", err));
});


app.get("/problem2", (req, res) => {
  getPartnersDetails()
    .then((partner) =>{
      getCapacities()
        .then((capcities)=>{
          problemStateMent_2(partner,capcities)
          .then(data => {createOutputFile(data,'output2.csv'); res.send(data) })
          .catch(err => console.log("ERROR in problem statement 2", err))
        })
        .catch((err)=> console.log("ERROR in getting capacity details", err))
     
    })
    .catch((err) => console.log("ERROR in getting partner details", err));
});


app.listen(8090, error => {
  if (error) {
    console.error(error);
  } else {
    console.info(
      "==> 🌎  Listening on port %s. Open up %s in your browser.",
      8090,
      process.env.PWD
    );
  }
});
