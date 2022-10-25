const fs = require('fs')
const { parseInputs, parsePartners} = require('./parser');
 
//To get the minimum cost of delivery
solution1();

//solution1() - Inside this function the solution for problem statement 1 is writtened here
    async function solution1() {
       const partners = await parsePartners();
       const input = await parseInputs();
       let result ;
       //clearing output1.csv file
       fs.writeFile("output1.csv", '', (err) => {
        if (err) throw err;
      });
//Iterating the deivery input
       for(let iterateInput = 0; iterateInput<input.length; iterateInput++){
        let minimumCost;
        let usageCost;
        let deliverable = false;
        let partnerId;

//Iterating the partners        
        for(let iteratePartners = 0; iteratePartners<partners.length; iteratePartners++){

          //conditions to check for the slab range and theatre
            if(input[iterateInput].theatre == partners[iteratePartners].theatre 
              && partners[iteratePartners].minSizeSlab <= input[iterateInput].size && partners[iteratePartners].maxSizeSlab >= input[iterateInput].size ){
    
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
                 partnerId = partnerId == undefined ?'""':partnerId;
                 minimumCost = minimumCost == undefined ?'""':minimumCost;
                 result = input[iterateInput].id +","+deliverable+","+ partnerId +","+minimumCost;
                 console.log("Final result",result);
                //appending the  result in output1.csv for every iteration
                 fs.appendFileSync("output1.csv", result + '\n');
                 
    }
} 

module.exports = {
    solution1
}