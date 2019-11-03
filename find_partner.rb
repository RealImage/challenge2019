require 'csv'

#Partner = Struct.new(:partner_id, :capacity)
#TheaterPartnerCost = Struct.new(:theater_id, :slab_start, :slab_end, :minimum_cost, :cost_per_gb, :partner_id, :capacity)

# class Partner
# 	attr_accessor :partner_id, :capacity

# 	def initialize(partner_id, capacity)
# 		self.partner_id = partner_id
# 		self.capacity = capacity
# 	end

# 	def partner
# end

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
		self.delivery_options = Array.new #[DeliveryOption.new(self.delivery_id, false)]
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
		puts "======= @deliverables = #{@deliverables.inspect} ==========="
		@deliverables.each do |deliverable|
			set_delivery_options(deliverable)
		end
		puts "======= after cost @deliverables = #{@deliverables.inspect} ==========="
		output1 = @deliverables.map { |del| del.lowest_delivery_cost  }
		puts "======= after output1 = #{output1.inspect} ==========="
		puts "======= output1 ============================"
		write_csv('output1', output1)
	end

	def low_cost_delivery_within_capacity

		puts "======= ====================================="
		@deliverables.each{|del| del.sort_delivery_options}
		puts "======= after sort @deliverables = #{@deliverables.inspect} ==========="

		puts "======= deliverables ============================"
		@deliverables.each do |del|
			puts "======= del.size_of_delivery = #{del.size_of_delivery} ============================"
			del.delivery_options.each do |del_opt|
				puts "=== #{del_opt.delivery_id}, #{del_opt.is_deliverable}, #{del_opt.partner_id}, #{del_opt.cost_of_delivery} ==="
			end
		end

		delivery_options_array = @deliverables.map { |e| e.delivery_options  } #.map { |p| p.partner_id }
		delivery_options_array.select!{|e| e.size>0}

		puts "======= delivery_options_array = #{delivery_options_array} ============================"
		combinations = delivery_options_array.first.product(*delivery_options_array.drop(1))
		final_hash = Hash.new
		last_combo_cost = nil
		combinations.each_with_index do |combo, index|
			puts "============================== = #{index} ================================="
			available_capacity = @partners.dup
			puts "============= first loop available_capacity = #{available_capacity} =========="
			puts "============= first loop @partners = #{@partners} =========="
			current_hash = Hash.new
			current_combo_cost = 0
			combo.each do |del_opt|
				puts "============= second loop available_capacity = #{available_capacity} =========="
				puts "============= second loop @partners = #{@partners} =========="
				puts "============= current_combo_cost = #{current_combo_cost}, last_combo_cost = #{last_combo_cost} =========="
				puts "============= second loop del_opt.cost_of_delivery = #{del_opt.cost_of_delivery} =========="
				if available_capacity[del_opt.partner_id] >= del_opt.size_of_delivery
					current_hash[del_opt.delivery_id] = del_opt
					available_capacity[del_opt.partner_id] -= del_opt.size_of_delivery
					current_combo_cost += del_opt.cost_of_delivery
				else
					puts "============= loop break = #{current_combo_cost} =========="
					current_hash = final_hash.dup
					current_combo_cost = last_combo_cost
					break
				end
			end
			puts "============= current_combo_cost = #{current_combo_cost} =========="
			if last_combo_cost.nil? || last_combo_cost > current_combo_cost
				last_combo_cost = current_combo_cost
				final_hash = current_hash
			end
			puts "======= #{index} final_hash = #{final_hash} ============================"
		end

		puts "======= final_hash = #{final_hash} ============================"
		output2 = @deliverables.map { |del| final_hash[del.delivery_id].nil? ? DeliveryOption.new(del.delivery_id, false) : final_hash[del.delivery_id]  }
		puts "======= output2 = #{output2} ============================"
		write_csv('output2', output2)
	end

	def read_csv(file_name, have_headers=true)
		file = "#{file_name}.csv"
		csv_text = File.read(file)
		csv_data = CSV.parse(csv_text, :headers => have_headers)
		csv_data
	end

	def write_csv(file_name, output)
		CSV.open("#{file_name}.csv", "wb") do |csv|
			output.each do |delivery|
				puts "=== #{delivery.delivery_id}, #{delivery.is_deliverable}, #{delivery.partner_id}, #{delivery.cost_of_delivery} ==="
				csv << [delivery.delivery_id, delivery.is_deliverable, delivery.partner_id, delivery.cost_of_delivery]
			end
		end
	end

	def fetch_details
		partners_capacity = read_csv('capacities')
		puts "======= parners_capacity = #{partners_capacity.inspect} ==========="
		initialize_partners(partners_capacity)
		puts "======= @partners = #{@partners.inspect} ==========="
		partners_cost = read_csv('partners')
		puts "======= partners_cost = #{partners_cost.inspect} ==========="
		initialize_partners_costs(partners_cost)
		puts "======= @theater_partner_cost = #{@theater_partner_cost.inspect} ==========="

		t1_user = find_by_partner_id(@theater_partner_cost, 'P3')
		puts "======= t1_user = #{t1_user.inspect} ==========="

		

		#parners_capacity = read_csv('partners')
	end

	#@partners - Hash , key = partner_id, value - capacity
	def initialize_partners(partners_capacity)
		@partners = Hash.new
		partners_capacity.each do |partner|
			@partners[partner[0].strip] = partner[1].strip.to_f # << Partner.new(partner[0].strip, partner[1].strip.to_f)
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
# def update_quality(items)
#   items.each do |item|
# 	updateter = ItemUpdaterFactory.getUpdater(item.name)
# 	item = updateter.updateItem(item)
#   end
# end

# # Factory Class
# class ItemUpdaterFactory
# 	def self.getUpdater(name)		
# 		case name.downcase
# 		when "aged brie"
# 			AgedBrieUpdater.new
# 		when "backstage passes"
# 			BackstagePassesUpdater.new
# 		when "sulfuras"
# 			SulfurasUpdater.new
# 		when "conjured"
# 			ConjuredUpdater.new
# 		else
# 			NormalUpdater.new
# 		end
# 	end
# end

# class  NormalUpdater
# 	def updateItem(item)
# 		updated_item = update_sell_in(item)
# 		updated_item = update_quality(item)
# 		updated_item
# 	end
	
# 	def update_sell_in(item)
# 		item.sell_in += get_sell_in_degrade
# 	end
	
# 	def update_quality(item)
# 		item.quality += (get_quality_degrade(item) * get_quality_degrade_multiplier(item)) 
# 		normalize_item(item)
# 	end
	
# 	def normalize_item(item)
# 		item.quality = 0 if item.quality < 0 
# 		item.quality = getQualityMax if item.quality > getQualityMax
# 		item
# 	end
	
# 	def get_sell_in_degrade
# 		-1
# 	end
	
# 	def get_quality_degrade(item=nil)
# 		-1
# 	end
	
# 	def get_quality_degrade_multiplier(item)
# 		isExpired(item) ? 2 : 1
# 	end
	
# 	def isExpired(item)
# 		item.sell_in < 0
# 	end
	
# 	def getQualityMax
# 		50
# 	end
# end

# class  AgedBrieUpdater < NormalUpdater
	
# 	def get_quality_degrade(item=nil)
# 		1
# 	end
	
# 	def get_quality_degrade_multiplier(item)
# 		1
# 	end
# end

# class  BackstagePassesUpdater < AgedBrieUpdater
	
# 	def get_quality_degrade(item=nil)
# 		quality_loss = super
# 		quality_loss += 1 if item.sell_in < 10
# 		quality_loss += 1 if item.sell_in < 5
# 		quality_loss = (- item.quality) if item.sell_in < 0
# 		quality_loss
# 	end
	
# end

# class  SulfurasUpdater < NormalUpdater
	
# 	def updateItem(item)
# 		item
# 	end
# end

# class  ConjuredUpdater < NormalUpdater
	
# 	def get_quality_degrade_multiplier(item)
# 		qLoss = super * 2
# 		qLoss
# 	end
# end
######### DO NOT CHANGE BELOW #########
