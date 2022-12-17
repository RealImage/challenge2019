require 'csv'
require_relative 'lib/partner'
require_relative 'lib/input'
require_relative 'lib/capacity'

class Delivery
  attr_accessor :partners, :inputs, :capacities

  def initialize
    @partners = Partner.get_partners
    @inputs = Input.get_inputs
    @capacities = Capacity.get_capacities
  end

  def solve_problem1
    puts "Get Each row containing delivery ID, indication if delivery is possible (true/false), cost of delivery and selected partner. "
    print_solution(soution1)
  end

  def solve_problem2
    puts "Get each row containing delivery ID, indication if delivery is possible (true/false), cost of delivery and selected partner. "
    print_solution(solution2)
  end

  def soution1
    output_arr = []
    inputs.reverse.each do |input|
      choose_partners = fetch_partners(input)
      if choose_partners.any?
        partner, cost = minimum_partner(choose_partners, input.delivery_size)
        output_arr << {
          delivery_id: input.delivery_id,
          possible: true,
          cost: cost,
          partner_id: partner.partner_id
        }
      else
        output_arr << {
          delivery_id: input.delivery_id,
          possible: false,
          cost: nil,
          partner_id: nil
        }
      end
    end
    output_arr
  end

  def solution2
    output_arr = []
    inputs.reverse.each do |input|
      choose_partners = fetch_partners(input)

      if choose_partners.any?
        partner = partner_capacity(choose_partners, input.delivery_size)
        cost = input.delivery_size * partner.cost_per_gb
        cost = partner.min_cost > cost ? partner.min_cost : cost
        output_arr << { 
          cost: cost,
          delivery_id: input.delivery_id,
          partner_id: partner.partner_id,
          possible: true
        }
      else
       output_arr << {
          delivery_id: input.delivery_id,
          possible: false,
          cost: nil,
          partner_id: nil
        }
      end
    end
    output_arr

  end

  def fetch_partners(input)
    partners.select do |partner|
      range = partner.size_slab.split("-")
      partner.theatre_id == input.theatre_id && (range[0].to_i..range[1].to_i).include?(input.delivery_size)
    end
  end

  def minimum_partner(selected_partners, delivery_size)
    min_partner = nil
    min_cost = nil

    selected_partners.each_with_index do |partner, i|
      cost = delivery_size * partner.cost_per_gb
      cost = partner.min_cost > cost ? partner.min_cost : cost

      if i == 0
        min_cost = cost
        min_partner = partner
      elsif min_cost >= cost
        min_cost = cost
        min_partner = partner
      end
    end
    [min_partner, min_cost]
  end

  def partner_capacity(selected_partners, delivery_size)
    min_partner = nil

    selected_partners.each do |partner|
      capacity = capacities.find { |c| c.partner_id == partner.partner_id }
      next unless capacity.capacity >= delivery_size
      min_partner = partner
      capacity.capacity = capacity.capacity - delivery_size
      break
    end
    min_partner
  end

  def print_solution(solutions)
    solutions.reverse.each do |output|
      if output[:cost]
        puts "#{output[:delivery_id].to_s}, #{output[:possible]}, #{output[:cost].to_s}, #{output[:partner_id].to_s}"
      else
        puts "#{output[:delivery_id].to_s}, #{output[:possible]}, " +  "\"\"" + ", " +  "\"\""
      end
    end
  end
end


Delivery.new.solve_problem1
Delivery.new.solve_problem2