let migrate = require('../../../utils/migration');

migrate({
    'name': 'no_exceeds_output',
    'mappingHeaders': {
        'Delivery ID': {'key': 'delivery_id', 'type': 'String'},
        'Delivery Possible': {'key': 'delivery_possible', 'type': 'Boolean'},
        'Partner ID': {'key': 'partner_id', 'type': 'String'},
        'Cost of Delivery': {'key': 'cost_of_delivery', 'type': 'Number'},
    },
    'csv_file_path': './test_store/csv/output/no_exceeds_output.csv',
    'json_file_path': './test_store/no_exceeds_output.json'
});