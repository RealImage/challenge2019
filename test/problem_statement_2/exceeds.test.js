const assert = require('assert');

const processProblem2 = require('../../utils/test/processProblem2');

const partners = require('../../data_store/partners.json');
const capacities = require('../../data_store/capacities.json');

const input = require('../../test_store/input.json');
const output = require('../../test_store/output2.json');

it('Problem Statement 2 - exceeds maximum capacity - suggested output', function(done){
    const processedOutput = processProblem2(input, partners, capacities);
    assert(JSON.stringify(output) == JSON.stringify(processedOutput));
    done();
});