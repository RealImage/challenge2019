let migrate = require('../../../utils/migration');

migrate({
    'name': 'no_exceeds_input',
    'mappingHeaders': {
        'Delivery ID': {'key': 'delivery_id', 'type': 'String'},
        'Size of Delivery': {'key': 'size_of_delivery', 'type': 'Number'},
        'Theatre': {'key': 'theatre_id', 'type': 'String'},
    },
    'csv_file_path': './test_store/csv/input/no_exceeds_input.csv',
    'json_file_path': './test_store/no_exceeds_input.json'
});