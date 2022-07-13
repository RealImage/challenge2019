require "./csv/import_csv"
require "./models/delivery_partner"

class Delivery
  extend ImportCsv
  
  attr_accessor :id, :size, :theater_id

  FILE_NAME = "input.csv"

  def initialize(arr)
    @id = arr[0].strip
    @size = arr[1].strip.to_i
    @theater_id = arr[2].strip
  end

  def self.all
    import FILE_NAME
  end

  def estimation(delivery_partners)
    delivery_partner = delivery_partners.select { |dp| dp.theater_id == theater_id && dp.slab_size.include?(size) && (dp.minimum_cost <= dp.cost_per_unit * size) }
                                        .min { |dp, dq| dp.cost_per_unit <=> dq.cost_per_unit }

    if delivery_partner.nil?
      [id, false, "", ""]
    else
      [id, true, delivery_partner.partner_id, (delivery_partner.cost_per_unit.to_i * size)]
    end
  end

  def estimation_upon_capacity(delivery_partners, delivery_capacities)
    available_capacities = delivery_capacities.select { |dc| dc.capacity >= size }
    available_capacity_partners = available_capacities.map(&:partner_id)

    valid_partners = delivery_partners
                      .select { |dp| available_capacity_partners.include?(dp.partner_id) && (dp.theater_id == theater_id) && dp.slab_size.include?(size) && (dp.minimum_cost <= dp.cost_per_unit * size) }

    minimum_cost_partner = valid_partners.min { |dp, dq| dp.cost_per_unit <=> dq.cost_per_unit }
    
    if minimum_cost_partner.nil?
      [id, false, "", ""]
    else
      delivery_capacity = available_capacities.find { |c| c.partner_id == minimum_cost_partner.partner_id }
      
      if delivery_capacity.nil?
        [id, false, "", ""]
      else
        delivery_capacity.capacity -= size
        [id, true, minimum_cost_partner.partner_id, (minimum_cost_partner.cost_per_unit * size)]
      end
    end
  end
end