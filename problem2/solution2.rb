# frozen_string_literal: true

require 'csv'

# Define snake_case method for String Class
class String
  def snake_case
    strip.downcase.gsub(' ', '_').gsub(/[^0-9A-Za-z_]/, '')
  end
end

# Read partners data from csv
partners_data = CSV.read('./partners.csv',
                         headers: true).map(&:to_h).map do |a|
  a.each_value(&:strip!).transform_keys(&:snake_case)
end

capacities_data = CSV.read('./capacities.csv',
                           headers: true).map(&:to_h).map do |a|
  a.each_value(&:strip!).transform_keys(&:snake_case)
end

# Group Data by theatre
grouped_data = partners_data.group_by { |row| row['theatre'] }
# Read input data for deliveries from csv
input_data = CSV.read('./input.csv')

# calculate cost per delivery
def fetch_cost(data_hash, size_of_delivery)
  calc_cost = data_hash['cost_per_gb'].to_i * size_of_delivery.to_i
  data_hash['minimum_cost'].to_i < calc_cost ? calc_cost : data_hash['minimum_cost'].to_i
end

# fetch possible partners
def fetch_possible_partners(grouped_data, theatre, size_of_delivery)
  grouped_data[theatre].select do |a|
    b = a['size_slab_in_gb'].split('-').map(&:to_i)
    size_of_delivery.to_i.between?(b[0], b[1])
  end
end

def get_all_combinations(*arr)
  arr.shift.product(*arr)
end

possible_partners_info = {}
possible_partners = {}
input_data.each do |input_row|
  partners = fetch_possible_partners(grouped_data, input_row[2], input_row[1])
  possible_partners_info[input_row[0]] = partners
  partner_names = partners.map { |a| a['partner_id'] }
  possible_partners[input_row[0]] = partner_names.empty? ? ['-'] : partner_names
end

possible_combinations = get_all_combinations(*possible_partners.values)

outputs = {}
possible_combinations.each do |c|
  c1 = {}
  capacities_data.each { |a| c1[a['partner_id']] = a['capacity_in_gb'].to_i }
  total_cost = 0
  c.each_with_index do |partner, i|
    if partner == '-'
      0
    else
      partner_info = possible_partners_info.values[i].find { |a| a['partner_id'] == partner and a['theatre'] == input_data[i][2] }
      delivery_gb = input_data.find { |j| j[0] == possible_partners_info.keys[i] }[1]
      cost = if possible_partners_info[possible_partners_info.keys[i]].empty?
               0
             else
               fetch_cost(partner_info, delivery_gb)
             end
      total_cost += cost
      c1[partner] = c1[partner] - delivery_gb.to_i
    end
  end
  outputs[c] = total_cost unless c1.values.any?(&:negative?)
end

solution_combination = outputs.select { |_k, v| v == outputs.values.min }.first.to_a

CSV.open('problem2/output.csv', 'w') do |csv|
  input_data.each do |input_row|
    delivery = input_row[0]
    possible = (possible_partners_info[delivery].empty? ? false : true)
    partner =  solution_combination[0][possible_partners_info.keys.index(delivery)].gsub('-', '')
    if partner == ''
      cost = ''
    else
      partner_info = possible_partners_info[delivery].find { |a| a['partner_id'] == partner }
      delivery_gb = input_row[1]
      cost = fetch_cost(partner_info, delivery_gb)
    end
    csv << [delivery, possible, partner, cost]
  end
end
