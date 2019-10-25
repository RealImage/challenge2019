var fs = require('fs');

function csvparser(fileName) {
    var array = fs.readFileSync(fileName).toString().split("\n");
    let list = [];
    let propname = array[0].replace(/"/g, "").split(',');

    for (var i = 1; i < array.length; i++) {
        let obj = {};
        let propvalue = array[i].split(',');
        if (propvalue[0].trim() !== "") {
            for (var j in propvalue) {
                obj[propname[j].trim()] = propvalue[j].trim();
            }
            list.push(obj);
        }
    }
    return list;
}

function inputparser(fileName) {
    var array = fs.readFileSync(fileName).toString().split("\n");
    let list = [];
    let propname = array[0].replace(/"/g, "").split(',');

    for (var i = 0; i < array.length; i++) {
        let obj = [];
        let propvalue = array[i].split(',');
        if (propvalue[0].trim() !== "") {
            for (var j in propvalue) {
                obj.push(propvalue[j].trim());
            }
            list.push(obj);
        }
    }
    return list;
}


const table1 = csvparser("./partners.csv");

const inputs = inputparser("./input.csv");

inputs.forEach((e) => {
    let deliveryID = e[0];
    let sizeofdelivery = parseInt(e[1]);
    let theatreID = e[2];
    // console.log(deliveryID, sizeofdelivery, theatreID);
    let max = Number.MAX_VALUE;
    let parterID = "";
    // console.log(max);
    for (var i = 0; i < table1.length; i++) {
        if (table1[i].Theatre === theatreID) {
            let sizeslabfrom = parseInt(table1[i]["Size Slab (in GB)"].split("-")[0]);
            let sizeslabto = parseInt(table1[i]["Size Slab (in GB)"].split("-")[1]);
            // console.log(sizeslabfrom, sizeslabto, sizeofdelivery);
            if (sizeslabto >= sizeofdelivery && sizeslabfrom <= sizeofdelivery) {
                // console.log(sizeslabfrom, sizeslabto);
                let cost = sizeofdelivery * parseInt(table1[i]["Cost Per GB"]);
                if (max > cost) {
                    max = cost;
                    max = parseInt(table1[i]["Minimum cost"]) > max ? parseInt(table1[i]["Minimum cost"]) : max;
                    parterID = table1[i]["Partner ID"];
                }
            }
        }
    }
    let transferable = max === Number.MAX_VALUE ? false : true;
    let value = max === Number.MAX_VALUE ? "" : max;
    console.log(`${deliveryID},${transferable},${parterID},${value}`);
})