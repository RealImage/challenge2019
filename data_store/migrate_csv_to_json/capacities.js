let migrate = require('../../utils/migration');

migrate({
    'name': 'capcities',
    'mappingHeaders': {
        'Partner ID': {'key': 'partner_id', 'type': 'String'},
        'Capacity (in GB)': {'key': 'max_capacity', 'type': 'Number'}
    },
    'csv_file_path': './data_store/csv/capacities.csv',
    'json_file_path': './data_store/capacities.json'
});