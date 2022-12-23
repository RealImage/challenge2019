require 'csv'
require 'json'

# partner_data = CSV.parse(File.read('partners.csv'), :headers=>true, header_converters: lambda {|f| f.strip}, :converters => :all )

def partner_data
	CSV.parse(File.read('partners.csv'), :headers=>true, header_converters: lambda {|f| f.strip}, converters: lambda {|f| f ? f.strip : nil})
end

def input_data
	input_data = CSV.parse(File.read('input.csv')).sort_by {|row| row[1].to_i}.reverse
end

def check_content_size(row, delivery_content_size)
	row["Size Slab (in GB)"].split('-').first.to_i <= delivery_content_size && delivery_content_size <= row["Size Slab (in GB)"].split('-').last.to_i
end

def calculate_delivery_cost(partner_row, delivery_content_size)
	cost_per_rate = delivery_content_size * partner_row['Cost Per GB'].to_i
	partner_minimum_cost = partner_row['Minimum cost'].to_i
	cost = cost_per_rate < partner_minimum_cost ? partner_minimum_cost : cost_per_rate
end

def find_eligible_delivery_partners(delivery_theatre, delivery_content_size)
	eligible_delivery_partners = partner_data.select{|row| row['Theatre'] == delivery_theatre && check_content_size(row, delivery_content_size)}
end

def capacities_data
	CSV.parse(File.read("capacities.csv"), headers: true, header_converters: lambda {|f| f.strip}, converters: lambda {|f| f ? f.strip : nil})
end

def checking_partner_capacity?(partner_id, delivery_content_size)
	partner_total_capacity = capacities_data.find { |row| row['Partner ID'] == partner_id }['Capacity (in GB)'].to_i 
	delivery_content_size <= (partner_total_capacity - $partner_allocated_capacities[partner_id])
end

def solution
	partner_data
	input_data
	output1 = []
	output2 = []
	$partner_allocated_capacities = Hash.new(0)
	input_data.each do |input_row|
		delivery_id = input_row[0]
		delivery_content_size = input_row[1].to_i
		delivery_theatre = input_row[2]
		eligible_delivery_partners = find_eligible_delivery_partners(delivery_theatre, delivery_content_size)
		costs = []
		if eligible_delivery_partners.length != 0
			eligible_delivery_partners.each do |partner_row|
				costs << calculate_delivery_cost(partner_row, delivery_content_size)
			end
			minimum_cost = costs.min
			minimum_cost_index = costs.index(minimum_cost)
			cheapest_partner = eligible_delivery_partners[minimum_cost_index]
			output1 << [delivery_id, true, cheapest_partner['Partner ID'], minimum_cost]
			until checking_partner_capacity?(cheapest_partner['Partner ID'], delivery_content_size)
				costs.delete_at(minimum_cost_index)
				eligible_delivery_partners.delete_at(minimum_cost_index)
				break if eligible_delivery_partners.empty?
				minimum_cost = costs.min
				minimum_cost_index = costs.index(minimum_cost)
				cheapest_partner = eligible_delivery_partners[minimum_cost_index]
			end
			if eligible_delivery_partners.empty?
				output2 << [delivery_id, false, "", ""]
			else
				output2 << [delivery_id, true, cheapest_partner['Partner ID'], minimum_cost]
      			$partner_allocated_capacities[cheapest_partner['Partner ID']] += minimum_cost
			end
		else
			output1 << [delivery_id, false, "", ""]
			output2 << [delivery_id, false, "", ""]
		end
	end
	CSV.open("my_output1.csv", "w") do |csv|
		output1.sort_by {|row| row[0]}.each do |row|
			csv << row
		end
	end
	puts "Problem statement 1 output is created and the name is 'my_output1.csv'"

	CSV.open("my_output2.csv", "w") do |csv|
		output2.sort_by {|row| row[0]}.each do |row|
			csv << row
		end
	end
	puts "Problem statement 2 output is created and the name is 'my_output2.csv'"

end

solution