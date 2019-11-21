const csvtojson = require('csvtojson');
const fs = require('fs');

const INPUT = fs.readFileSync('./input.csv')
    .toString()
    .split('\n')
    .map(input => input.split(','));

var PARTNERS_FILE = './partners.csv';
var CAPACITIES_FILE = './capacities.csv';

var make_ranklist = function(input, partners) {
  var delivery = input[0];
  var size = parseInt(input[1]);
  var partner_list = [];

  for(let partner of partners) {
    let slab = partner['Size Slab (in GB)'].split('-');
    let partner_slab_min = parseInt(slab[0]);
    let partner_slab_max = parseInt(slab[1]);

    if((partner_slab_min > size) || (partner_slab_max < size)) continue;

    let partner_min_cost = parseInt(partner['Minimum cost']);
    let partner_per_gb_rate = parseInt(partner['Cost Per GB']);
    let partner_id = partner['Partner ID'];


    let cost = partner_per_gb_rate * size;
    if(cost < partner_min_cost) cost = partner_min_cost;

    partner_list.push([ partner_id, cost, size ]);
  }

  partner_list.sort((a,b) => {
    if(a[1] < b[1]) return -1;
    return 1;
  });

  return partner_list;
};

var generate_rank_list = async function(input) {
  var partners = await csvtojson().fromFile(PARTNERS_FILE);
  var rank_data = {};

  for(let i of input) {
    var possible_partners = filter_partners_by_theatre(partners, i[2]);

    rank_data[i[0]] = make_ranklist(i, possible_partners);
  }

  return rank_data;
};

var filter_partners_by_theatre = function(partners, theatre) {
  return partners.filter(p => p.Theatre == theatre);
};


var print_ps1 = async function(input) {
  var ranked_data = await generate_rank_list(input);

  let result = [];
  for(let i of input) {
    let data = [ i[0] ];

    if(!ranked_data[i[0]].length) {
      data.push([false, "", "" ]);
    }
    else {
      data.push(true);
      let partner_id = ranked_data[i[0]][0][0];
      let cost = ranked_data[i[0]][0][1];

      data.push(cost);
      data.push(partner_id);
    }

    result.push(data.toString());
  }
  await fs.writeFileSync('./out2.csv', result);
  console.log(result);
};

print_ps1(INPUT);
