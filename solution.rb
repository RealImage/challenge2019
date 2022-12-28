load 'lib/methods.rb'
def qube_cinema_solution
	qube_partner_cinema
	qube_cinema_input_data
	output1 = []
	output2 = []
	final_output = []
	$partner_allocated_capacities = Hash.new(0)
	qube_cinema_input_data.each do |input_row|
		delivery_id = input_row[0]
		delivery_content_size = input_row[1].to_i
		delivery_theatre = input_row[2]
		eligible_delivery_partners = qube_cinema_find_eligible_delivery_partners(delivery_theatre, delivery_content_size)
		costs = []
		if eligible_delivery_partners.length != 0
			eligible_delivery_partners.each do |partner_row|
				costs << qube_cinema_calculate_delivery_cost(partner_row, delivery_content_size)
			end
			minimum_cost = costs.min
			minimum_cost_index = costs.index(minimum_cost)
			cheapest_partner = eligible_delivery_partners[minimum_cost_index]
			output1 << [delivery_id, true, cheapest_partner['Partner ID'], minimum_cost]
			until qube_cinema_partner_capacity?(cheapest_partner['Partner ID'], delivery_content_size)
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
	final_output << output1 << output2
	generationg_qube_output_for_cinema(final_output)
	

end

qube_cinema_solution