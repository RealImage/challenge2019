# frozen_string_literal: true

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
require_relative 'csv_helper'

class QubeChallenge
  include CsvHelper

  def initialize(problem_type)
    @deliveries = format_data('input', false)
    @partners_data = cook_csv_data('partners.csv', true)

    case problem_type
    when 'one'
      problem_one
    when 'two'
      @capacities = format_data('capacities', true)
      problem_two
    end
  end

  def problem_one
    delivery_details = @deliveries.map do |delivery_id, delivery|
      find_optimal_priced_partner(delivery_id, delivery)
    end

    write_to_csv_and_print_content('output1.csv', delivery_details, 'Optimal priced partners: ')
  end

  def problem_two
    delivery_details = @deliveries.map do |delivery_id, delivery|
      find_optimal_priced_partner(delivery_id, delivery, { consider_capacity: true })
    end

    write_to_csv_and_print_content('output2.csv', delivery_details, 'Optimal priced partners according to capacity: ')
  end

  private

  def find_optimal_priced_partner(delivery_id, delivery, options = {})
    # Filter partners by availability
    eligible_partners = get_eligible_partners(delivery)
    delivery_possible = eligible_partners.any?
    @optimal_partner = {}

    if delivery_possible
      @optimal_partner = find_partner(eligible_partners, options[:consider_capacity],
                                      delivery[:size_of_delivery])
    end

    [delivery_id, delivery_possible, @optimal_partner[:partner_id] || '', @optimal_partner[:delivery_cost] || '']
  end

  def get_eligible_partners(delivery)
    eligible_partners = @partners_data.filter do |partner|
      size_array = partner[:size_slab_in_gb].split('-').map(&:to_i)
      delivery_size_range = size_array[0]..size_array[1]

      (partner[:theatre] == delivery[:theatre_id]) &&
        (delivery_size_range.include? delivery[:size_of_delivery])
    end

    # Find delivery cost for each eligible partner
    eligible_partners.map do |partner|
      # Take minimum cost into consideration
      final_cost = [partner[:cost_per_gb] * delivery[:size_of_delivery], partner[:minimum_cost]].max
      { partner_id: partner[:partner_id], delivery_cost: final_cost }
    end
  end

  def find_partner(eligible_partners, consider_capacity, delivery_size)
    # Sort based on the optimal price
    sorted_partners = eligible_partners.sort_by { |partner| partner[:delivery_cost] }

    # Set Optimal partner based on cheapest price
    @optimal_partner = sorted_partners.first

    # Consider capacity of partners
    if consider_capacity
      sorted_partners.each do |partner|
        partner_id = partner[:partner_id]
        next unless delivery_size <= @capacities[partner_id]

        @optimal_partner = partner
        @capacities[partner_id] -= delivery_size
        break
      end
    end

    @optimal_partner
  end

  def format_data(data_type, header)
    data = cook_csv_data("#{data_type}.csv", header)
    formatted_data = {}

    data.each do |d|
      case data_type
      when 'input'
        formatted_data[d[0]] = { size_of_delivery: d[1], theatre_id: d[2] }
      when 'capacities'
        formatted_data[d[:partner_id]] = d[:capacity_in_gb]
      end
    end

    formatted_data
  end
end

QubeChallenge.new('one')
QubeChallenge.new('two')
