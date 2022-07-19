class ApplicationController < ActionController::Base
require 'csv'

# data from input.csv
def get_input_data  
    CSV.open('./input.csv').each.to_a.compact.map do |input|
       input.map(&:strip)
    end
end

# strip is used to remove all whitespaces and split is used to break sentence in work with given delimiter and then picking min and max values
def get_partners_data
    partner_data = []
    CSV.foreach("./partners.csv", headers: true) do |row|
        partner_data << {
            theatre: row['Theatre'].strip,
            size_lab_min: row['Size Slab (in GB)'].strip.split('-')[0].to_i,  
            size_lab_max: row['Size Slab (in GB)'].strip.split('-')[1].to_i,
            minimun_cost: row['Minimum cost'].to_i,
            cost_per_gb: row['Cost Per GB'].to_i,
            partner_id: row['Partner ID'].strip
        }
    end
    partner_data
end

def get_capacities_data
    capacity_data = []
    CSV.foreach("./capacities.csv", headers: true) do |row|
        capacity_data << {
            # removing trailing spaces with strip and converting to integers
            partner_id: row['Partner ID'].strip,
            capacity: row['Capacity (in GB)'].to_i
        }
    end
    capacity_data
end
end
