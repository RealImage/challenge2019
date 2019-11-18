require_relative 'core/distributors'
require_relative 'core/csv_helper'

include CSVHelper

def find_the_minimum(input_file, partner_file, capacity_file)
  partners = parse_partners(partner_file)
  capacities = parse_capacities(capacity_file)

  d = Distributors.new(partners, capacities)

  input = parse_input(input_file)

  output = d.calculate_minimum_delivery(input)

  convert_to_csv(output)
end

def find_the_minimum_with_capacity(input_file, partner_file, capacity_file)
  partners = parse_partners(partner_file)
  capacities = parse_capacities(capacity_file)

  d = Distributors.new(partners, capacities)

  input = parse_input(input_file)

  output = d.calculate_minimum_delivery_with_capacity(input)

  final_output = []
  output.sort.each do |k, v|
    final_output << [k, v[:possible], v[:partner_id], v[:cost_of_delivery]]
  end

  convert_to_csv(final_output)
end

# find_the_minimum('input.csv', 'partners.csv', 'capacities.csv')
# find_the_minimum_with_capacity('input.csv', 'partners.csv', 'capacities.csv')
