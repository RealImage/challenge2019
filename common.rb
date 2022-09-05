
class Common
  def Common.partners_hash
    partners_hash = []
    CSV.foreach(("partners.csv"), headers: true) do |row|
      partners_hash << { theater_id: row['Theatre'].strip, slab: row['Size Slab (in GB)'].strip, min_cost: row['Minimum cost'].strip, cost: row['Cost Per GB'].strip, partner_id: row['Partner ID'].strip }
    end
    partners_hash
  end

  def Common.capacities_hash
    capacities_hash = {}
    CSV.foreach(("capacities.csv"), headers: true) do |row|
      capacities_hash[row['Partner ID'].strip] = row['Capacity (in GB)'].strip.to_i
    end
    capacities_hash
  end

  def Common.input_arr
    input_arr = []
    CSV.foreach(("input.csv"), headers: false) do |row|
      input_arr << row
    end
    input_arr
  end

  def Common.theater_with_partners_hash
    partners_hash = {}
    CSV.foreach(("partners.csv"), headers: true) do |row|
      partners_hash[row['Theatre'].strip] ||= {}
      partners_hash[row['Theatre'].strip][row['Partner ID'].strip] ||= []
      partners_hash[row['Theatre'].strip][row['Partner ID'].strip] << { slab: row['Size Slab (in GB)'].strip, min_cost: row['Minimum cost'].strip, cost: row['Cost Per GB'].strip }
    end
    partners_hash
  end

  def Common.print_delivery_partner_details(partner_data_hash = {})
    partner_data_hash&.each do |key, value|
      final_min_cost = value&.values&.sort.reject{ |k| k==0}.first
      is_can_deliver = !final_min_cost.nil?
      final_min_cost_partner = is_can_deliver ? value&.invert[final_min_cost] : ''
      puts "#{key}, #{is_can_deliver}, #{final_min_cost_partner}, #{final_min_cost}"
    end
  end

  def Common.get_partner_missed_slab_details(partners = [], input_slab = 0)
    is_partner_has_no_starting_cost_val = []
    starting_slab = []
    available_slab = 0
    previous_end_slab = 0
    partners&.each do |val|
      first_slab, last_slab = val[:slab].split('-')
      is_partner_has_no_starting_cost_val << (first_slab.to_i..last_slab.to_i).include?(0)
      starting_slab << first_slab.to_i
      f_slab = first_slab.to_i == previous_end_slab ? first_slab.to_i + 1 : first_slab.to_i
      available_slab += (f_slab..last_slab.to_i).include?(input_slab) ? 1 : 0
      previous_end_slab = last_slab.to_i
    end
    has_more_available_partner = available_slab > 1
    return [is_partner_has_no_starting_cost_val, starting_slab, has_more_available_partner]
  end

  def Common.get_calculation_details(partner = {}, is_partner_has_no_starting_cost_val = [])
    start_slab, end_slab = partner[:slab].split('-')
    gb_count = (start_slab..end_slab).to_a.count - 1
    is_partner_has_no_starting_cost = is_partner_has_no_starting_cost_val.uniq.include?(false) && !is_partner_has_no_starting_cost_val.uniq.include?(true)
    return [start_slab, end_slab, gb_count, is_partner_has_no_starting_cost]
  end
end
