let migrate = require('../../utils/migration');

migrate({
    'name': 'sample_capcities',
    'mappingHeaders': {
        'Partner ID': {'key': 'partner_id', 'type': 'String'},
        'Capacity (in GB)': {'key': 'max_capacity', 'type': 'Number'}
    },
    'csv_file_path': './data_store/csv/sample_capacities.csv',
    'json_file_path': './data_store/sample_capacities.json'
});