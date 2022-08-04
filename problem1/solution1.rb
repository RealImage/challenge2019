# frozen_string_literal: true

require 'csv'

# Define snake_case method for String Class
class String
  def snake_case
    strip.downcase.gsub(' ', '_').gsub(/[^0-9A-Za-z_]/, '')
  end
end

#Read partners data from csv
partners_data = CSV.read('./partners.csv',
                         headers: true).map(&:to_h).map do |a|
  a.each_value(&:strip!).transform_keys(&:snake_case)
end
#Group Data by theatre
grouped_data = partners_data.group_by { |row| row['theatre'] }
#Read input data for deliveries from csv
input_data = CSV.read('./input.csv')

def fetch_possible_partners(grouped_data, theatre, size_of_delivery)
  grouped_data[theatre].select do |a|
    b = a['size_slab_in_gb'].split('-').map(&:to_i)
    size_of_delivery.to_i.between?(b[0], b[1])
  end
end

def fetch_cost(data_hash, size_of_delivery)
  calc_cost = data_hash['cost_per_gb'].to_i * size_of_delivery.to_i
  data_hash['minimum_cost'].to_i < calc_cost ? calc_cost : data_hash['minimum_cost'].to_i
end

CSV.open('problem1/output.csv', 'w') do |csv|
  input_data.each do |input_row|
    possible_partners = fetch_possible_partners(grouped_data, input_row[2], input_row[1])
    csv << [input_row[0], false, '', ''] and next if possible_partners.empty?

    data_with_cost = possible_partners.map do |a|
      [a['partner_id'], fetch_cost(a, input_row[1])]
    end
    data_with_min_cost = data_with_cost.min_by { |_a| [1] }
    csv << [input_row[0], true, data_with_min_cost[0], data_with_min_cost[1]]
  end
end
