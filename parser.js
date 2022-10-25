const fs = require('fs')
const CsvReadableStream = require('csv-reader')
let input = [];

// parsing  partner csv and destructuring it
let partners = [];
function parsePartners() {
  return new Promise((resolve, reject) => {
    let partnerStream = fs.createReadStream('partners.csv', 'utf8');
    partnerStream
          .pipe(new CsvReadableStream({ skipHeader: true, parseNumbers: true, trim: true }))
          .on('data', function (row) {
              try {
                  const sizes = row[1].split('-');
             partners.push({
            'theatre': row[0],
            'minCost': row[2],
            'costPerGB': row[3],
            'partnerId': row[4],
            'minSizeSlab': +sizes[0],
            'maxSizeSlab': +sizes[1]
              })
              console.log("partners", partners);
            }
              catch (err) {
                  console.error("Error while parsing : " + err)
                  reject(err)
              }
          })
          .on('end', function () {
              resolve(partners)
          });
  })
}

// parsing input csv and destructuring it
function parseInputs() {
  return new Promise((resolve, reject) => {
    let inputStream = fs.createReadStream('input.csv', 'utf8');
    inputStream
          .pipe(new CsvReadableStream({ parseNumbers: true, trim: true }))
          .on('data', function (row) {
              try {
                input.push({
                  'id': row[0],
                  'size': row[1],
                  'theatre': row[2],
              })
            }
              catch (err) {
                  console.error("Error while parsing : " + err)
                  reject(err)
              }
          })
          .on('end', function () {
              resolve(input);
          });
  })
} 

// parsing capacity csv and destructuring it
let capacity = {};
function parseCapacity() {
    return new Promise((resolve, reject) => {
      let inputStream = fs.createReadStream('capacities.csv', 'utf8');
      inputStream
            .pipe(new CsvReadableStream({ skipHeader: true, parseNumbers: true, trim: true, }))
            .on('data', function (row) {
                try {
                    capacity[row[0]]= row[1];
              }
                catch (err) {
                    console.error("Error while parsing : " + err)
                    reject(err)
                }
            })
            .on('end', function () {
                resolve(capacity)
            });
    })
  } 

module.exports = {
    parsePartners,
    parseInputs,
    parseCapacity
}