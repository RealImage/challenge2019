require 'csv'

class TheaterPartnerCost
	attr_accessor :theater_id, :slab_start, :slab_end, :minimum_cost, :cost_per_gb, :partner_id

	def initialize(theater_id, slab_start, slab_end, minimum_cost, cost_per_gb, partner_id)
		self.theater_id = theater_id
		self.slab_start = slab_start
		self.slab_end = slab_end
		self.minimum_cost = minimum_cost
		self.cost_per_gb = cost_per_gb
		self.partner_id = partner_id
	end

	def with_in_slab?(delivery_size)
		delivery_size >= self.slab_start && delivery_size <= self.slab_end
	end

	def calculate_delivery_cost(deliverable)
		(deliverable.size_of_delivery * self.cost_per_gb) < self.minimum_cost ? self.minimum_cost : (deliverable.size_of_delivery * self.cost_per_gb)
	end
end

class DeliveryOption
	attr_accessor :delivery_id, :is_deliverable, :partner_id, :cost_of_delivery, :size_of_delivery

	def initialize(delivery_id, is_deliverable, partner_id=nil, cost_of_delivery=nil, size_of_delivery=nil)
		self.delivery_id = delivery_id
		self.is_deliverable = is_deliverable
		self.partner_id = partner_id
		self.cost_of_delivery = cost_of_delivery
		self.size_of_delivery = size_of_delivery
	end
end

class Deliverable
	attr_accessor :delivery_id, :size_of_delivery, :theater_id, :delivery_options
	
	def initialize(delivery_id, size_of_delivery, theater_id)
		self.delivery_id = delivery_id
		self.size_of_delivery = size_of_delivery
		self.theater_id = theater_id
		self.delivery_options = Array.new 
	end

	def sort_delivery_options
		self.delivery_options.sort_by!{ |opt| opt.cost_of_delivery }
	end

	def lowest_delivery_cost
		lowest = self.delivery_options.min_by{ |opt| opt.cost_of_delivery }
		lowest.nil? ? DeliveryOption.new(self.delivery_id, false) : lowest
	end
end


class FindPartner
	def initialize
		low_cost_delivery
		low_cost_delivery_within_capacity
	end

	def low_cost_delivery
		fetch_details
		input_data = read_csv('input', false)
		initialize_deliveries(input_data)
		@deliverables.each do |deliverable|
			set_delivery_options(deliverable)
		end
		output1 = @deliverables.map { |del| del.lowest_delivery_cost  }
		write_csv('output1', output1)
	end

	def low_cost_delivery_within_capacity
		delivery_options_array = @deliverables.map { |e| e.delivery_options  }
		delivery_options_array.select!{|e| e.size>0}

		combinations = delivery_options_array.first.product(*delivery_options_array.drop(1))
		final_hash = Hash.new
		last_combo_cost = nil
		combinations.each_with_index do |combo, index|
			available_capacity = @partners.dup
			current_hash = Hash.new
			current_combo_cost = 0
			combo.each do |del_opt|				
				if available_capacity[del_opt.partner_id] >= del_opt.size_of_delivery
					current_hash[del_opt.delivery_id] = del_opt
					available_capacity[del_opt.partner_id] -= del_opt.size_of_delivery
					current_combo_cost += del_opt.cost_of_delivery
				else
					current_hash = final_hash.dup
					current_combo_cost = last_combo_cost
					break
				end
			end
			if last_combo_cost.nil? || last_combo_cost > current_combo_cost
				last_combo_cost = current_combo_cost
				final_hash = current_hash
			end
		end

		output2 = @deliverables.map { |del| final_hash[del.delivery_id].nil? ? DeliveryOption.new(del.delivery_id, false) : final_hash[del.delivery_id]  }
		write_csv('output2', output2)
	end

	def read_csv(file_name, have_headers=true)
		file = "#{file_name}.csv"
		csv_text = File.read(file)
		csv_data = CSV.parse(csv_text, :headers => have_headers)
		csv_data
	end

	def write_csv(file_name, output)
		puts "================= #{file_name} =============="
		CSV.open("#{file_name}.csv", "wb") do |csv|
			output.each do |delivery|
				puts "=== #{delivery.delivery_id}, #{delivery.is_deliverable}, #{delivery.partner_id}, #{delivery.cost_of_delivery} ==="
				csv << [delivery.delivery_id, delivery.is_deliverable, delivery.partner_id, delivery.cost_of_delivery]
			end
		end
		puts "================= #{file_name}.csv updated =============="
	end

	def fetch_details
		partners_capacity = read_csv('capacities')
		initialize_partners(partners_capacity)
		partners_cost = read_csv('partners')
		initialize_partners_costs(partners_cost)
	end

	#@partners - Hash , key = partner_id, value - capacity
	def initialize_partners(partners_capacity)
		@partners = Hash.new
		partners_capacity.each do |partner|
			@partners[partner[0].strip] = partner[1].strip.to_f
		end
	end

	def initialize_partners_costs(partners_cost)
		@theater_partner_cost = Array.new
		partners_cost.each do |cost|
			slab_arr = cost[1].strip.split('-')
			@theater_partner_cost << TheaterPartnerCost.new(cost[0].strip, slab_arr[0].to_f, slab_arr[1].to_f, cost[2].strip.to_f, cost[3].strip.to_f, cost[4].strip)
		end
	end

	def initialize_deliveries(deliverables)
		@deliverables = Array.new
		@deliverable_size = Hash.new
		deliverables.each do |delivery|
			@deliverables << Deliverable.new(delivery[0].strip, delivery[1].strip.to_f, delivery[2].strip )
			@deliverable_size[delivery[0].strip] = delivery[1].strip.to_f
		end
	end

	def find_by_partner_id(entries, partner_id)
		entries.select { |e|  e.partner_id == partner_id}
	end

	def find_by_theater_id(entries, theater_id)
		entries.select { |e|  e.theater_id == theater_id}
	end

	def set_delivery_options(deliverable)
		available_partners = find_by_theater_id(@theater_partner_cost, deliverable.theater_id)
		available_partners.each do |t_cost|
			if t_cost.with_in_slab?(deliverable.size_of_delivery)
				deliverable.delivery_options << DeliveryOption.new(deliverable.delivery_id, true, t_cost.partner_id, t_cost.calculate_delivery_cost(deliverable), deliverable.size_of_delivery)
			end
		end
	end
end

FindPartner.new
