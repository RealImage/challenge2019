require 'csv'
require 'active_support/all'
require 'byebug'

class SolutionHelper
  def self.read_details
    partner_details = self.csv_to_hash("partners.csv")
    input_details = CSV.read("input.csv")
    capacity_limits = CSV.read("capacities.csv")
    capacity_limits.shift
    return partner_details, input_details, capacity_limits
  end
    
  def self.get_minimum_price(rate, partner)
    cost = partner.map{|i| i["Cost Per GB"].strip.to_i * rate.to_i }.min
    minimum_provide = partner.select{|val| val["Cost Per GB"].strip.to_i ==  (cost/rate.to_i)}
  end

  def self.output_csv(solution)
    CSV.open("output2.csv", "w") do |csv|
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
end