const assert = require('assert');

const processProblem2 = require('../../utils/test/processProblem2');

const partners = require('../../data_store/partners.json');
const capacities = require('../../data_store/capacities.json');

const input = require('../../test_store/no_exceeds_input.json');
const output = require('../../test_store/no_exceeds_output.json');

it('Problem Statement 2 - does not exceed maximum capacity', function(done){
    const processedOutput = processProblem2(input, partners, capacities);
    assert(JSON.stringify(output) == JSON.stringify(processedOutput));
    done();
});