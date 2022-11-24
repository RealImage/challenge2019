let fs = require('fs');
let fastcsv = require('fast-csv');

module.exports = function(config){
    try{
        console.log(`--- migration - ${config.name} - started ---`);
        let readableStreamInput = fs.createReadStream(config.csv_file_path);
        let jsonData = [];

        let mappingHeaders = config.mappingHeaders;
        
        fastcsv
            .fromStream(readableStreamInput, {headers: true})
            .on('data', (data) => {
                let rowData = {};
            
                Object.keys(data).forEach(current_key => {
                    let mappedHeader = mappingHeaders[current_key];

                    let value = (data[current_key]).trim();
                    switch(mappedHeader.type){
                        case 'Number':
                            value = parseInt(value);
                            break;
                        case 'Boolean':
                            value = (value == 'true');
                            break;
                    }

                    switch(mappedHeader.key){
                        case 'size_slab':
                            // to construct atomic value
                            let size = value.split('-');
                            rowData['min_size'] = parseInt(size[0]);
                            rowData['max_size'] = parseInt(size[1]);
                            break;
                        default:
                            rowData[mappedHeader.key] = value;
                    }
                });
            
                jsonData.push(rowData);
            }).on('end', () => {
                fs.writeFile(config.json_file_path, JSON.stringify(jsonData, null, '\t'), function(err) {
                    if(err) {
                        return console.log(err);
                    }
                    console.log(`--- migration - ${config.name} - terminated ---`);
                });
            });
    }catch(err){
        console.log(`ERROR LOG --- migration --- ${err.stack}`);
    }
}