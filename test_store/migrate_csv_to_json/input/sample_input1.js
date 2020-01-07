let migrate = require('../../../utils/migration');

migrate({
    'name': 'sample_input1',
    'mappingHeaders': {
        'Delivery ID': {'key': 'delivery_id', 'type': 'String'},
        'Size of Delivery': {'key': 'size_of_delivery', 'type': 'Number'},
        'Theatre': {'key': 'theatre_id', 'type': 'String'},
    },
    'csv_file_path': './test_store/csv/input/sample_input1.csv',
    'json_file_path': './test_store/sample_input1.json'
});