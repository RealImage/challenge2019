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
        let obj = [i];
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

let capacities = csvparser("./capacities.csv");

capacities = capacities.sort(function (a, b) { return b["Capacity (in GB)"] - a["Capacity (in GB)"] });

// console.log(capacities);

let inputs = inputparser("./input.csv");
inputs = inputs.sort(function (a, b) { return b[2] - a[2] });
let outputs = [];

inputs.forEach((e) => {
    let id = e[0];
    let deliveryID = e[1];
    let sizeofdelivery = parseInt(e[2]);
    let theatreID = e[3];
    // console.log(deliveryID, sizeofdelivery, theatreID);
    let max = Number.MAX_VALUE;
    let parterID = "";
    // console.log(max);
    // console.log('cpa', capacities);

    for (var capacity = 0; capacity < capacities.length; capacity++) {
        // console.log(capacities[capacity]["Capacity (in GB)"])
        for (var i = 0; i < table1.length; i++) {
            if (table1[i].Theatre === theatreID) {
                let sizeslabfrom = parseInt(table1[i]["Size Slab (in GB)"].split("-")[0]);
                let sizeslabto = parseInt(table1[i]["Size Slab (in GB)"].split("-")[1]);
                // console.log(sizeslabfrom, sizeslabto, sizeofdelivery);
                if (sizeslabto >= sizeofdelivery && sizeslabfrom <= sizeofdelivery
                    && parseInt(capacities[capacity]["Capacity (in GB)"]) >= sizeofdelivery
                    && capacities[capacity]["Partner ID"] === table1[i]["Partner ID"]) {
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
    }
    capacities.map((x) => {
        if (x["Partner ID"] === parterID)
            x["Capacity (in GB)"] = x["Capacity (in GB)"] - sizeofdelivery;
        return x;
    });
    // console.log('cpa', capacities);

    let transferable = max === Number.MAX_VALUE ? false : true;
    let value = max === Number.MAX_VALUE ? "" : max;
    outputs.push([id, deliveryID, transferable, parterID, value])
    // console.log(`deliveryID},${transferable},${parterID},${value}`);
});

outputs = outputs.sort(function (a, b) { return a[0] - b[0] });

outputs.forEach(x=> console.log(`${x[1]}, ${x[2]}, ${x[3]}, ${x[4]}`));