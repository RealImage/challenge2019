module Utility
  def no_delivery(input)
    {
      cost: nil, 
      delivery_id: input.delivery_id,
      partner_id: nil,
      possible: false
    }
  end

  def condition(range, input, capacity)
    (range[0].to_i..range[1].to_i).include?(input.req_slab)
  end

  def get_capacity(partner_id)
    capacities.find { |x| x.partner_id == partner_id }
  end

  def print(outputs)
    outputs.reverse.each do |output|
      if output[:cost]
        puts "#{output[:delivery_id]}, #{output[:possible]}, #{output[:cost].to_s} #{output[:partner_id].to_s}"
      else
        puts "#{output[:delivery_id]}, #{output[:possible]}, " +  "\"\"" + ", " +  "\"\""
      end
    end
  end

  def capacity_partner(selected_partners, req_slab)
    min_partner = nil

    selected_partners.each do |partner|
      capacity = get_capacity(partner.partner_id)

      next unless capacity.capacity >= req_slab
      min_partner = partner
      capacity.capacity = capacity.capacity - req_slab
      break
    end
    min_partner
  end

  def minimum_partner(selected_partners, req_slab)
    min_partner = nil
    min_cost = nil 
    selected_partners.each_with_index do |partner, index|
      cost = req_slab * partner.cost_per_gb
      cost = partner.min_cost > cost ? partner.min_cost : cost
      if index == 0
        min_cost = cost
        min_partner = partner
      elsif  min_cost >= cost
        min_cost = cost
        min_partner = partner
      end
    end 
    [min_partner, min_cost]
  end

  def fetch_partners(input)
    partners.select do |partner|
      range = partner.size_slab.split("-")
      partner.theatre == input.theatre && (range[0].to_i..range[1].to_i).include?(input.req_slab)
    end
  end
end
