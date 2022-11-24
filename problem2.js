const Fs = require('fs');
const CsvReadableStream = require('csv-reader');

let capacityStream = Fs.createReadStream('capacities.csv', 'utf8');
let capacities = {};
capacityStream
    .pipe(new CsvReadableStream({ skipHeader: true, parseNumbers: true, trim: true }))
    .on('data', function (row) {
        capacities[row[0]] = { max: row[1] };
    })
    .on('end', function () {
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
                        let inputData = [];
                        input.forEach((row, index) => {
                            let filteredPartners = partners.filter(partner => (partner.theatre == row.theatre) && (partner.minSize <= row.size) && (row.size <= partner.maxSize))
                            if (filteredPartners.length > 0) {
                                let calculatedCost = filteredPartners.map(({ minCost, costPerGB, partnerId }) => {
                                    return { size: row.size, transferCost: Math.max(minCost, costPerGB * row.size), partnerId }
                                })
                                inputData[index] = calculatedCost;
                            }
                            else inputData[index] = [];
                        })
                        const output = calculate(inputData);
                        let result;
                        input.forEach((row, index) => {
                            if (output.selection[index]) {

                                result = [row.id, true, inputData[output.selection[index][0]][output.selection[index][1]].partnerId, inputData[output.selection[index][0]][output.selection[index][1]].transferCost];
                            } else {
                                result = [row.id, false, "\"\"", "\"\""];
                            }
                            console.log(result.join(','));
                        })
                    });
            })
    });

/**
 * Iterates all possible cases and then find the best solution from the input
 */
function calculate(input) {
    let totalCases = input.reduce((count, item) => {
        return [...count, Math.max(1, item.length) * (count.length > 0 ? count[count.length - 1] : 1)];
    }, []);
    let range = totalCases.map(num => totalCases[totalCases.length - 1] / num);
    let score = [];
    for (let c = 0; c < totalCases[totalCases.length - 1]; ++c) {
        Object.keys(capacities).forEach(key => capacities[key]['remaining'] = capacities[key]['max']);
        let selection = {};
        let cost = 0;
        for (let i = 0; i < input.length; ++i) {
            if (input[i].length > 0) {
                let j = Math.floor(c / range[i]) % input[i].length;
                if (capacities[input[i][j]['partnerId']]['remaining'] >= input[i][j]['size']) {
                    selection[i] = [i, j];
                    cost += input[i][j]['transferCost'];
                    capacities[input[i][j]['partnerId']]['remaining'] -= input[i][j]['size'];
                }
            }
        }
        score.push({ cost, selection });
    }
    score.sort((a, b) => Object.keys(b.selection).length - Object.keys(a.selection).length || a.cost - b.cost);
    return score[0];
}