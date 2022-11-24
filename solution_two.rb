require 'csv'
require 'active_support/all'
require 'byebug'
require './solution_helper'

class SolutionTwo

  def self.calculate
    partners, inputs, capacity_limits = SolutionHelper.read_details
    solution = []
    inputs.map do |input|
      matched_partners = partners.select{|partner| partner["Theatre"].strip == input[2] && input[1].to_i.between?(partner["Size Slab (in GB)"].strip.split("-")[0].to_i, partner["Size Slab (in GB)"].strip.split("-")[1].to_i)}
      matched_partners = SolutionHelper.get_minimum_price(input[1], matched_partners)
      have_capacity = matched_partners.empty? ? false : self.get_capacity(input, matched_partners[0]["Partner ID"], capacity_limits)

      unless !matched_partners.empty? && have_capacity
        solution << { delivery_id: input[0], deliverable: false, partner_id: "", cost: "" }
      else
        cost = SolutionHelper.get_cost(input[1].to_i, matched_partners[0]["Cost Per GB"].strip.to_i, matched_partners[0]["Minimum cost"].strip.to_i)
        solution << { delivery_id: input[0], deliverable: true, partner_id: matched_partners[0]["Partner ID"], cost: cost }
      end
    end
    SolutionHelper.output_csv(solution, "output2.csv")
  end

  def self.get_capacity(input, partner_id, capacity_limits)
    capacity = capacity_limits.select{|limit| limit[0].include?(partner_id)}
    return input[1].to_i < capacity[0][1].to_i
  end
 
end