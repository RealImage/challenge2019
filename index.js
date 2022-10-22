const solution1 = require('./Problem1/problem1')
const solution2 = require('./Problem2/problem2')

async function solution() {
    await solution1.minimumCost()
    await solution2.minimumCapacity()
};

solution();