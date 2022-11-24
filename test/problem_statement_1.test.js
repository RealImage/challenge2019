const assert = require('assert');

const processProblem1 = require('../utils/test/processProblem1');

const partners = require('../data_store/partners.json');

const input = require('../test_store/input.json');
const output = require('../test_store/output1.json');

it('Problem Statement 1', function(done){
    const processedOutput = processProblem1(input, partners);
    assert(JSON.stringify(output) == JSON.stringify(processedOutput));
    done();
});