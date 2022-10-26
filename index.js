const {problem1 } = require("./problem1.js");
const {problem2 } = require("./problem2.js");

const Problem = async () => {
  //Problem Solution 1 Call
 // await problem1();
     
    await problem2();
    await problem1();
   
};

Problem();