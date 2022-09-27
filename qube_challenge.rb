# Sample Input (input.csv)

# D1,150,T1
# D2,325,T2
# D3,510,T1
# D4,700,T2

# Sample Output for problem 1 (output1.csv)

# D1,true,P1,2000
# D2,true,P1,3250
# D3,true,P3,15300
# D4,false,"",""

# Sample Output for problem 2 (output2.csv)

# D1,true,P1,2000
# D2,true,P2,3500
# D3,true,P3,15300
# D4,false,"",""

# require 'csv'
require_relative 'csv_helper.rb'

include CsvHelper

class QubeChallenge
  def initialize(problem_type)
    @deliveries = get_formatted_input_data
    @partners_data = cook_csv_data("partners.csv", true)

    if problem_type == "one"
      problem_one
    elsif problem_type == "two"
      @capacities = get_formatted_capacities_data
      problem_two
    end
  end

  def problem_one
    delivery_details = @deliveries.map do |delivery_id, delivery|
      find_optimal_priced_partner(delivery_id, delivery)
    end

    write_to_csv_and_print_content("output1.csv", delivery_details, "Optimal priced partners: ")
  end

  def problem_two
    delivery_details = @deliveries.map do |delivery_id, delivery|
      find_optimal_priced_partner(delivery_id, delivery, {consider_capacity: true})
    end

    write_to_csv_and_print_content("output2.csv", delivery_details, "Optimal priced partners according to capacity: ")
  end

  def find_optimal_priced_partner(delivery_id, delivery, options = {})
    # Filter partners by availability
    eligible_partners = get_eligible_partners(delivery)
    is_delivery_possible = eligible_partners.any?

    if is_delivery_possible
      final_partners = []
      # Format eligible partners data
      formatted_eligible_partners = eligible_partners.map do |partner|
        delivery_cost = partner[:cost_per_gb] * delivery[:size_of_delivery]
        # Take minimum cost into consideration
        final_cost = [delivery_cost, partner[:minimum_cost]].max
        {partner_id: partner[:partner_id], delivery_cost: final_cost}
      end

      # Sort based on the optimal price
      sorted_partners = formatted_eligible_partners.sort_by{|partner| partner[:delivery_cost]}
      # Set Optimal partner based on cheapest price
      final_list = {sorted_partners: sorted_partners, optimal_partner: sorted_partners.first}

      # Consider capacity of partners
      if options[:consider_capacity]
        sorted_partners.each do |partner|
          partner_id = partner[:partner_id]
          if delivery[:size_of_delivery] <= @capacities[partner_id]
            final_list[:optimal_partner] = partner
            @capacities[partner_id] -= delivery[:size_of_delivery]
            break
          end
        end
      end

      partner_id = final_list[:optimal_partner][:partner_id]
      final_delivery_cost = final_list[:optimal_partner][:delivery_cost]
    else
      partner_id = ""
      final_delivery_cost = ""
    end

    [delivery_id, is_delivery_possible, partner_id, final_delivery_cost]
  end

  private

  def get_formatted_input_data
    input_data = cook_csv_data("input.csv", false)
    formatted_input = {}

    input_data.each do |input|
      formatted_input[input[0]] = {size_of_delivery: input[1], theatre_id: input[2]}
    end

    formatted_input
  end

  def get_formatted_capacities_data
    capacities_data = cook_csv_data("capacities.csv", true)
    capacities = {}

    capacities_data.each do |capacity|
      capacities[capacity[:partner_id]] = capacity[:capacity_in_gb]
    end

    capacities
  end

  def get_eligible_partners(delivery)
    @partners_data.filter do |partner|
      size_array = partner[:size_slab_in_gb].split("-").map(&:to_i)
      delivery_size_range = size_array[0]..size_array[1]

      (partner[:theatre] == delivery[:theatre_id]) &&
      (delivery_size_range.include? delivery[:size_of_delivery])
    end
  end

end

QubeChallenge.new("one")
QubeChallenge.new("two")
