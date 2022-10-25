const fs = require('fs')
const { parseInputs, parsePartners, parseCapacity} = require('./parser');

//To get the minimum cost of total delivery
solution2();

async function solution2() {
    const partners = await parsePartners();
    const input = await parseInputs();
    let result ;
    let sortedResult=[];
    //clearing output2.csv file
    fs.writeFile("output2.csv", '', (err) => {
     if (err) throw err;
   });
  
   //sorting the input based on content size in Descending order
  input.sort(function(a, b){return b.size - a.size});
  const capacity = await parseCapacity();
  //Iterating the deivery input
  for(let iterateInput = 0; iterateInput<input.length; iterateInput++){
  
     let minimumCost;
     let usageCost;
     let deliverable = false;
     let partnerId;
     
     for(let iteratePartners = 0; iteratePartners<partners.length; iteratePartners++){
     //conditions to check for the slab range, theatre and partner cost usage
         if(input[iterateInput].theatre == partners[iteratePartners].theatre 
           && partners[iteratePartners].minSizeSlab <= input[iterateInput].size 
           && partners[iteratePartners].maxSizeSlab >= input[iterateInput].size 
           && capacity[partners[iteratePartners].partnerId]>=input[iterateInput].size){
  
              usageCost = input[iterateInput].size * partners[iteratePartners].costPerGB;
              //appending the minimum usage cost
              usageCost = usageCost<partners[iteratePartners].minCost ? partners[iteratePartners].minCost : usageCost
              deliverable = true;
              console.log("usageCost", minimumCost);
  
              if(minimumCost){
                 partnerId =minimumCost > usageCost ? partners[iteratePartners].partnerId : partnerId;
                 minimumCost = minimumCost > usageCost ? usageCost : minimumCost;
             }
             else{ 
                 minimumCost = usageCost;
                 partnerId = partners[iteratePartners].partnerId;
             }}
  }
              //minusing partner capacity based on the delivery assigned 
              capacity[partnerId]=capacity[partnerId]-input[iterateInput].size;
              console.log("capacityDecreased",capacity);
              partnerId = partnerId == undefined ?'""':partnerId;
              minimumCost = minimumCost == undefined ?'""':minimumCost;
              result = input[iterateInput].id +","+deliverable+","+ partnerId +","+minimumCost;
              console.log("result",result);
              sortedResult.push(result);
              
  }
  console.log("sortedResult",sortedResult.sort());
  for (let writeFile of sortedResult){
              //appending the  result in output2.csv for every iteration
              fs.appendFileSync("output2.csv", writeFile + '\n');
  }
  
  }

module.exports = {
    solution2
}