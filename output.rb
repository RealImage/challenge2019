require 'csv'

#Fetches 'partners' file and removes all the unnecessary whitespaces.
PARTNERS_TABLE = CSV.parse(File.read("partners.csv"), headers: true, header_converters: lambda {|f| f.strip}, converters: lambda {|f| f ? f.strip : nil})

CAPACITIES_TABLE = CSV.parse(File.read("capacities.csv"), headers: true, header_converters: lambda {|f| f.strip}, converters: lambda {|f| f ? f.strip : nil})

#Fetches 'input' file and sorts it in descending order of content size (required for Problem 2 to ensure least overall cost).
input_table = CSV.parse(File.read("input.csv"), converters: lambda {|f| f ? f.strip : nil})
SORTED_INPUT_TABLE = input_table.sort_by {|row| row[1].to_i}.reverse

def content_size_within_range?(row, delivery_content_size)
  row["Size Slab (in GB)"].split('-').first.to_i <= delivery_content_size && delivery_content_size <= row["Size Slab (in GB)"].split('-').last.to_i 
end

def calculate_partner_delivery_cost(partner_row, delivery_content_size)
  cost_as_per_rate = delivery_content_size*partner_row['Cost Per GB'].to_i
  minimum_cost = partner_row['Minimum cost'].to_i
  cost = cost_as_per_rate < minimum_cost ? minimum_cost : cost_as_per_rate
end

def partner_has_enough_spare_capacity?(partner_id, delivery_content_size)
  partner_total_capacity = CAPACITIES_TABLE.find { |row| row['Partner ID'] == partner_id }['Capacity (in GB)'].to_i  
  delivery_content_size <= (partner_total_capacity - $partner_allocated_capacities[partner_id])
end

output_1_csv_rows = []
output_2_csv_rows = []
$partner_allocated_capacities = Hash.new(0)
SORTED_INPUT_TABLE.each do |row|
  delivery_id = row[0]
  delivery_content_size = row[1].to_i
  delivery_theatre = row[2]
  eligible_delivery_partners = PARTNERS_TABLE.select { |row| row['Theatre'] == delivery_theatre && content_size_within_range?(row, delivery_content_size) }
  costs_array = []

  if eligible_delivery_partners.any?
    eligible_delivery_partners.each do |partner_row|
      costs_array << calculate_partner_delivery_cost(partner_row, delivery_content_size)
    end
    min_cost = costs_array.min
    min_index = costs_array.index(min_cost)
    cheapest_partner_row = eligible_delivery_partners[min_index]

    output_1_csv_rows << [delivery_id, true, cheapest_partner_row['Partner ID'], min_cost]

    until partner_has_enough_spare_capacity?(cheapest_partner_row['Partner ID'], delivery_content_size)
      #Removes the partner that cannot accomodate the delivery.
      costs_array.delete_at(min_index)
      eligible_delivery_partners.delete_at(min_index)

      break if eligible_delivery_partners.empty?

      #Finds the next cheapest delivery partner.
      min_cost = costs_array.min
      min_index = costs_array.index(min_cost)
      cheapest_partner_row = eligible_delivery_partners[min_index]
    end  

    if eligible_delivery_partners.empty?
      output_2_csv_rows << [delivery_id, false]
    else
      output_2_csv_rows << [delivery_id, true, cheapest_partner_row['Partner ID'], min_cost]
      $partner_allocated_capacities[cheapest_partner_row['Partner ID']] += min_cost
    end
  else
    output_1_csv_rows << [delivery_id, false]
    output_2_csv_rows << [delivery_id, false]
  end
end

#Output of problem statement 1.
CSV.open("myoutput1.csv", "w") do |csv|
  #Sorts the rows by delivery ID and adds them into the output CSV.
  output_1_csv_rows.sort_by {|row| row[0]}.each do |row|
    csv << row
  end
end
puts "Created CSV file for problem statement 1 titled 'myoutput1.csv'"

#Output of problem statement 2.
CSV.open("myoutput2.csv", "w") do |csv|
  #Sorts the rows by delivery ID and adds them into the output CSV.
  output_2_csv_rows.sort_by {|row| row[0]}.each do |row|
    csv << row
  end
end
puts "Created CSV file for problem statement 2 titled 'myoutput2.csv'"