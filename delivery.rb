require 'csv'
require_relative 'src/partner'
require_relative 'src/input'
require_relative 'src/capacity'
require_relative 'src/utility'

class Delivery
  include Utility
  attr_accessor :partners, :inputs, :capacities

  def initialize
    @partners = Partner.partners_data
    @inputs = Input.input_data
    @capacities = Capacity.capacity_data
  end

  def problem1
    puts "Problem 1"
    print(solution1())
  end

  def problem2
    puts "\n\nProblem 2"
    print(solution2())
  end

  def solution1
    outputs = []
    inputs.reverse.each do |input|
      selected_partners = fetch_partners(input)
    
      if selected_partners.any?
        partner, cost = minimum_partner(selected_partners, input.req_slab)
        outputs << { 
          cost: cost,
          delivery_id: input.delivery_id,
          partner_id: partner.partner_id,
          possible: true
        }
      else
        outputs << no_delivery(input)
      end
    end
    outputs
  end

  def solution2
    outputs = []
    inputs.reverse.each do |input|
      selected_partners = fetch_partners(input)

      if selected_partners.any?
        partner = capacity_partner(selected_partners, input.req_slab)
        cost = input.req_slab * partner.cost_per_gb
        cost = partner.min_cost > cost ? partner.min_cost : cost
        outputs << { 
          cost: cost,
          delivery_id: input.delivery_id,
          partner_id: partner.partner_id,
          possible: true
        }
      else
        outputs << no_delivery(input)
      end
    end
    outputs
  end
end

Delivery.new.problem1
Delivery.new.problem2