require 'csv'

def qube_partner_cinema
	CSV.parse(File.read('partners.csv'), :headers=>true, header_converters: lambda {|f| f.strip}, converters: lambda {|f| f ? f.strip : ''})
end

def qube_cinema_input_data
	input_data = CSV.parse(File.read('input.csv')).sort_by {|row| row[1].to_i}.reverse
end

def qube_cinema_content_size(row, delivery_content_size)
	row["Size Slab (in GB)"].split('-').first.to_i <= delivery_content_size && delivery_content_size <= row["Size Slab (in GB)"].split('-').last.to_i
end

def qube_cinema_calculate_delivery_cost(partner_row, qube_cinema_delivery_content_size)
	cost_per_rate = qube_cinema_delivery_content_size * partner_row['Cost Per GB'].to_i
	partner_minimum_cost = partner_row['Minimum cost'].to_i
	cost = cost_per_rate < partner_minimum_cost ? partner_minimum_cost : cost_per_rate
end

def qube_cinema_find_eligible_delivery_partners(delivery_theatre, delivery_content_size)
	eligible_delivery_partners = qube_partner_cinema.select{|row| row['Theatre'] == delivery_theatre && qube_cinema_content_size(row, delivery_content_size)}
end

def qube_cinema_capacities_data
	CSV.parse(File.read("capacities.csv"), headers: true, header_converters: lambda {|f| f.strip}, converters: lambda {|f| f ? f.strip : nil})
end

def qube_cinema_partner_capacity?(partner_id, delivery_content_size)
	partner_total_capacity = qube_cinema_capacities_data.find { |row| row['Partner ID'] == partner_id }['Capacity (in GB)'].to_i 
	delivery_content_size <= (partner_total_capacity - $partner_allocated_capacities[partner_id])
end

def generationg_qube_output_for_cinema(output_qube)
	output_qube.each.with_index(1) do | element, index |
		CSV.open("output#{index}.csv", "w") do |csv|
			element.sort_by {|row| row[0]}.each do |row|
				csv << row
			end
		end
		puts "Qube Cinema content 'output#{index}.csv'"
	end
end