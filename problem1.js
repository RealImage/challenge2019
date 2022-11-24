const Fs = require('fs');
const CsvReadableStream = require('csv-reader');

let partnerStream = Fs.createReadStream('partners.csv', 'utf8');
let partners = [];
partnerStream
    .pipe(new CsvReadableStream({ skipHeader: true, parseNumbers: true, trim: true }))
    .on('data', function (row) {
        const sizes = row[1].split('-');
        partners.push({
            'theatre': row[0],
            'sizeSlab': row[1],
            'minCost': row[2],
            'costPerGB': row[3],
            'partnerId': row[4],
            'minSize': +sizes[0],
            'maxSize': +sizes[1]
        })
    })
    .on('end', function () {
        let inputStream = Fs.createReadStream('input.csv', 'utf8');
        let input = [];
        inputStream
            .pipe(new CsvReadableStream({ parseNumbers: true, trim: true }))
            .on('data', function (row) {
                input.push({
                    'id': row[0],
                    'size': row[1],
                    'theatre': row[2],
                })
            })
            .on('end', function () {
                input.forEach(row => {
                    let filteredPartners = partners.filter(partner => partner.theatre == row.theatre && partner.minSize <= row.size && row.size <= partner.maxSize)
                    let result;
                    if (filteredPartners.length > 0) {
                        let calculatedCost = filteredPartners.map(({ minCost, costPerGB, partnerId }) => {
                            return { transferCost: Math.max(minCost, costPerGB * row.size), partnerId }
                        })
                        calculatedCost.sort((a, b) => a.transferCost - b.transferCost);
                        result = [row.id, true, calculatedCost[0].partnerId, calculatedCost[0].transferCost];
                    }
                    else result = [row.id, false, "\"\"", "\"\""];
                    console.log(result.join(','));
                })
            });
    });