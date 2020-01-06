let migrate = require('../../../utils/migration');

migrate({
    'name': 'output1',
    'mappingHeaders': {
        'Delivery ID': {'key': 'delivery_id', 'type': 'String'},
        'Delivery Possible': {'key': 'delivery_possible', 'type': 'Boolean'},
        'Partner ID': {'key': 'partner_id', 'type': 'String'},
        'Cost of Delivery': {'key': 'cost_of_delivery', 'type': 'Number'},
    },
    'csv_file_path': './test_store/csv/output/output1.csv',
    'json_file_path': './test_store/output1.json'
});