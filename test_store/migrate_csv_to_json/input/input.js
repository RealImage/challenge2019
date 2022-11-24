let migrate = require('../../../utils/migration');

migrate({
    'name': 'input',
    'mappingHeaders': {
        'Delivery ID': {'key': 'delivery_id', 'type': 'String'},
        'Size of Delivery': {'key': 'size_of_delivery', 'type': 'Number'},
        'Theatre': {'key': 'theatre_id', 'type': 'String'},
    },
    'csv_file_path': './test_store/csv/input/input.csv',
    'json_file_path': './test_store/input.json'
});