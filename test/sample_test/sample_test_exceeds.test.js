const assert = require('assert');

const processProblem2 = require('../../utils/test/processProblem2');

const partners = require('../../data_store/sample_partners.json');
const capacities = require('../../data_store/sample_capacities.json');

const input = require('../../test_store/sample_input_exceeds.json');
const output = require('../../test_store/sample_output_exceeds.json');

it('Sample Test Exceeds', function(done){
    const processedOutput = processProblem2(input, partners, capacities);
    assert(JSON.stringify(output) == JSON.stringify(processedOutput));
    done();
});