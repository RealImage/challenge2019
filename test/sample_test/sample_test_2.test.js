const assert = require('assert');

const processProblem1 = require('../../utils/test/processProblem1');

const partners = require('../../data_store/sample_partners.json');

const input = require('../../test_store/sample_input2.json');
const output = require('../../test_store/sample_output2.json');

it('Sample Test 2', function(done){
    const processedOutput = processProblem1(input, partners);
    assert(JSON.stringify(output) == JSON.stringify(processedOutput));
    done();
});