require 'csv'
require 'active_support/all'
require 'byebug'
require './solution_helper'

class SolutionOne

  def self.calculate
    partners, inputs = SolutionHelper.read_details
    solution = []
    inputs.map do |input|
      matched_partners = partners.select{|partner| partner["Theatre"].strip == input[2] && input[1].to_i.between?(partner["Size Slab (in GB)"].strip.split("-")[0].to_i, partner["Size Slab (in GB)"].strip.split("-")[1].to_i)}
      matched_partners = SolutionHelper.get_minimum_price(input[1], matched_partners)
      if matched_partners.empty?
        solution << { delivery_id: input[0], deliverable: false, partner_id: "", cost: "" }
      else
        cost = SolutionHelper.get_cost(input[1].to_i, matched_partners[0]["Cost Per GB"].strip.to_i, matched_partners[0]["Minimum cost"].strip.to_i)
        solution << { delivery_id: input[0], deliverable: true, partner_id: matched_partners[0]["Partner ID"], cost: cost }
      end
    end
    SolutionHelper.output_csv(solution, "output1.csv")
  end
      
end

