class CalculationService

  def initialize(partners, input_data)
    @partners = partners
    @input_data = input_data
    @result = []
  end

  def calculate_minimal_cost
    @input_data.each do|row|
      delivery_possible = false
      mincost = ""
      delivery_partner = ""
      @partners.each do|partner|
        next unless valid_partner?(partner, row)
        if row[1].to_i >= partner.min_size && row[1].to_i < partner.max_size
          total_cost = calculate_total_cost(partner, row)
          if mincost.to_i == 0 || total_cost < mincost
            mincost = total_cost
            delivery_partner = partner.partner_id
            delivery_possible = true
          end
        end
      end
      assign_capacity_to_partner(delivery_partner, row[1])
      @result << build_result(row, mincost, delivery_partner, delivery_possible)
    end
    @result
  end

  def calculate_minimal_cost_by_capacity(capacities)
    input_map = build_input_map
    partner_delivery_map = build_partner_delivery_map
    partner_delivery_map.each do|partner_id, results|
      capacity_map, assigned_capacity_map = build_partner_capacity_map
      next if assigned_capacity_map[partner_id].nil? || capacity_map[partner_id].nil?
      if capacity_exceeded?(partner_id, capacity_map, assigned_capacity_map)
        reassign_partner(input_map, partner_delivery_map, assigned_capacity_map, partner_id)
      end
    end
    @result
  end
  
  private
  def valid_partner?(partner, row)
    partner.theatre_id == row[2]
  end

  def update_partner_capacity(partner_id, row)
    @partners.each do|partner|
      next unless partner.partner_id == partner_id
      partner.max_capacity -= row[1].to_i
    end
  end

  def capacity_exceeded?(partner_id, capacity_map, assigned_capacity_map)
    capacity_map[partner_id] < assigned_capacity_map[partner_id]
  end

  def reassign_partner(input_map, partner_delivery_map, assigned_capacity_map, partner_id)
    delivery_data = partner_delivery_map[partner_id]
    min_diff = new_min_cost = prev_diff = 0
    partner_total_cost = get_total_cost_for_partner(delivery_data)
    delivery_possible = false
    delivery_partner_id = delivery_id = ""
    delivery_data.each do|row|
      input_row = input_map[row[:delivery]]
      @partners.each do|partner|
        next unless partner.theatre_id == input_row[2] && partner.partner_id != partner_id
        next if partner.max_capacity < (assigned_capacity_map[partner.partner_id].to_i + input_row[1].to_i)
        if input_row[1].to_i >= partner.min_size && input_row[1].to_i < partner.max_size
          total_cost = calculate_total_cost(partner, input_row)
          if new_min_cost.to_i == 0 || total_cost < new_min_cost
            new_final_cost = partner_total_cost - row[:min_cost] + total_cost
            if min_diff == 0 || min_diff < prev_diff
              min_diff = (new_final_cost - partner_total_cost).abs
              new_min_cost = total_cost
              delivery_partner_id = partner.partner_id
              delivery_possible = true
              delivery_id = input_row[0]
              prev_diff = min_diff
            end
          end
        end  
      end
    end
    update_result(new_min_cost, delivery_possible, delivery_partner_id, delivery_id)
  end

  def assign_capacity_to_partner(partner_id, capacity)
    @partners.each do|partner|
      if partner.partner_id == partner_id
        partner.capacity_assigned += capacity.to_i
      end
    end
  end

  def update_result(new_min_cost, delivery_possible, delivery_partner, delivery_id)
    @result.each do|result|
      next unless result[:delivery] == delivery_id
      result[:delivery_possible] = delivery_possible
      result[:min_cost] = new_min_cost
      result[:delivery_partner] = delivery_partner
    end
  end

  def calculate_total_cost(partner, row)
    total_cost = row[1].to_i * partner.cost_per_gb
    total_cost < partner.min_cost ? partner.min_cost : total_cost 
  end

  def get_total_cost_for_partner(delivery_data)
    delivery_data.inject(0) {|sum, hash| sum + hash[:min_cost]}
  end

  def build_result(row, mincost, delivery_partner, delivery_possible)
    {
      delivery: row[0],
      delivery_possible: delivery_possible,
      min_cost: mincost,
      delivery_partner: delivery_partner
    }
  end

  def build_input_map
    input_map = {}
    @input_data.each do|row|
      input_map[row[0]] = row
    end
    input_map
  end

  def build_partner_capacity_map
    capacity_map = {}
    assigned_capacity_map = {}
    @partners.each do|partner|
      if capacity_map[partner.partner_id].nil?
        capacity_map[partner.partner_id] = partner.max_capacity
        assigned_capacity_map[partner.partner_id] = partner.capacity_assigned
      end
    end
    return capacity_map, assigned_capacity_map
  end

  def build_partner_delivery_map
    partner_delivery_map = {}
    @result.each do|row|
      if partner_delivery_map[row[:delivery_partner]].nil?
        partner_delivery_map[row[:delivery_partner]] = [row]
      else
        partner_delivery_map[row[:delivery_partner]] << row
      end
    end
    partner_delivery_map
  end


  def partner_capacity_mapper
    mapping = {}
    @partners.each do|partner|
      next if mapping[partner.partner_id]
      mapping[partner.partner_id] = partner.max_capacity
    end
    mapping
  end

end