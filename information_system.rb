# frozen_string_literal: true

require './database'
require './csv_operations'
require 'pp'
# Information system to display menu and output results
class InformationSystem
  include CsvOperations
  def initialize
    @partners = Database.new(load_csv(file: 'partners.csv'))
    @capacities = Database.new(load_csv(file: 'capacities.csv'))
    @input = Database.new(load_input_csv)
  end

  def display_menu
    puts "What would you like to do?
      1: Display partners
      2: Display capacities
      3: Display input
      4: Problem Statement 1
      5: Problem Statement 2
      Any other key to exit"
    choice = gets
    process_choice(choice)
  end

  private

  def process_choice(choice)
    case choice.to_i
    when 1
      pp @partners.data
      display_menu
    when 2
      pp @capacities.data
      display_menu
    when 3
      pp @input.data
      display_menu
    when 4
      pp problem_statement_1
      display_menu
    when 5
      pp problem_statement_2
      display_menu
    else
      exit
    end
  end

  def problem_statement_1
    output = []
    @input.data.each do |row|
      partner_with_cost = cost_effective_partner(row['Theatre ID'], row['Delivery Size'])
      output_row = []
      output_row[0] = row['Delivery ID']
      output_row[1] = !partner_with_cost.nil?
      output_row[2] = partner_with_cost.nil? ? '' : partner_with_cost.first
      output_row[3] = partner_with_cost.nil? ? '' : partner_with_cost.last
      output.push(output_row)
    end
    File.write('output1.csv', output.map(&:to_csv).join)
    output
  end

  def problem_statement_2
    output = []
    balance = @capacities.data.map(&:values).to_h

    @input.data.sort_by { |x| x['Delivery Size'] }.each do |input_row|
      size = input_row['Delivery Size'].to_i
      theatre_id = input_row['Theatre ID']

      eligible_partners = possible_partners(theatre_id, size)
      if eligible_partners.any?
        cheapest_partner, balance = cost_effective_partner_with_balance(eligible_partners, balance, size)

        final_price = cost_for_size(cheapest_partner, size)
        output.push([input_row['Delivery ID'], 'true', cheapest_partner['Partner ID'], final_price])
      else
        output.push([input_row['Delivery ID'], 'false', '', ''])
      end
    end
    File.write('output2.csv', output.map(&:to_csv).join)
    output
  end

  def cost_effective_partner_with_balance(eligible_partners, balance, size)
    cheapest_partner = nil

    eligible_partners.each do |partner|
      next unless balance[partner['Partner ID']].to_i >= size

      cheapest_partner = partner
      balance[partner['Partner ID']] = balance[partner['Partner ID']].to_i - size
      break
    end
    [cheapest_partner, balance]
  end

  def cost_effective_partner(theatre_id, size)
    possible_partner_plans = possible_partners(theatre_id, size)
    return if possible_partner_plans.empty?

    cost_per_partner = {}
    possible_partner_plans.each { |hsh| cost_per_partner[hsh['Partner ID']] = cost_for_size(hsh, size) }
    cost_per_partner.min_by(&:last)
  end

  def possible_partners(theatre_id, size)
    partners_for_theatre = @partners.find(key_value_pair: { 'Theatre' => theatre_id })
    return nil if partners_for_theatre.empty?

    partners_for_theatre.select do |x|
      Range.new(*x['Size Slab (in GB)'].split('-').map(&:to_i)).include?(size.to_i)
    end
  end

  def cost_for_size(plan, size)
    total_cost = plan['Cost Per GB'].to_i * size.to_i
    return plan['Minimum cost'].to_i if total_cost < plan['Minimum cost'].to_i

    total_cost
  end
end

