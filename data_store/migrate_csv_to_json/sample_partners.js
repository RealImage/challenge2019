let migrate = require('../../utils/migration');

migrate({
    'name': 'sample_partners',
    'mappingHeaders': {
        'Theatre': {'key': 'theatre_id', 'type': 'String'},
        'Size Slab (in GB)': {'key': 'size_slab', 'type': 'String'},
        'Minimum cost': {'key': 'minimum_cost', 'type': 'Number'},
        'Cost Per GB': {'key': 'cost_per_gb', 'type': 'Number'},
        'Partner ID': {'key': 'partner_id', 'type': 'String'}
    },
    'csv_file_path': './data_store/csv/sample_partners.csv',
    'json_file_path': './data_store/sample_partners.json'
});