require 'csv'
require 'active_support/all'
require 'byebug'

class SolutionHelper
  def self.read_details
    partner_details = self.csv_to_hash("partners.csv")
    input_details = CSV.read("input.csv")
    capacity_limits = CSV.read("capacities.csv")
    return partner_details, input_details, capacity_limits
  end
    
  def self.get_minimum_price(rate, partner)
    cost = partner.map{|i| i["Cost Per GB"].strip.to_i * rate.to_i }.min
    minimum_provide = partner.select{|val| val["Cost Per GB"].strip.to_i ==  (cost/rate.to_i)}
  end

  def self.output_csv(solution, file)
    CSV.open(file, "w") do |csv|
      solution.each do |sol|
        csv << sol.values
      end
    end
  end

  def self.csv_to_hash(data_file)
    data = []
    CSV.foreach(data_file, headers: true) do |row|
        data << row.to_hash
    end
    return data
  end

  def self.get_cost(size, cost_per_gb, minimum_cost)
    temp_cost = size * cost_per_gb
    cost = minimum_cost < temp_cost ? temp_cost : minimum_cost
  end
end