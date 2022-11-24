/* Execute the following steps */

---to install dependencies---
$npm install

---to migrate CSV file to JSON file (Data Store)---
$npm run MIGRATE_DATA_STORE

---to migrate CSV file to JSON file (Test Store)---
$npm run MIGRATE_TEST_STORE

---to run automated test (using mocha)---
$npm run test


/***********

- All tests will run successfully except the test `Problem Statement 2 - exceeds maximum capacity - suggested output` because it has different output other than suggested output (I hope the suggested ouput is wrong).

- You can see the test `Problem Statement 2 - exceeds maximum capacity - other possibility` where the output based on my logic is correct.

************/