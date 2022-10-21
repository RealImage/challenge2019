const problemStatement1 = require('./problem/problem1.js')
const problemStatement2 = require('./problem/problem2.js')
// const { maximumCapacity } = require('./problem2.js')



const Problem = async()=>{

    //Problem Solution 1 Call
    await problemStatement1.minimumDeliveryCost() 
    //Problem Solution 1 Call
    await problemStatement2.minimumCapacity() 
}

Problem()