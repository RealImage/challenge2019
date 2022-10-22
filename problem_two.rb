# frozen_string_literal: true

require './shared'

module Solution
  class ProblemTwo
    def self.execute
      partners, capacities = Shared.initialize_csv
      input = Shared.initialize_input
      result = []
      balance = capacities.map(&:values).to_h

      input.sort_by { |x| x[1] }.reverse.each do |input_row|
        size = input_row[1]
        theatre_id = input_row.last

        eligible_partners = Shared.eligible_partners(partners, theatre_id, size).sort_by { |x| x['cost'] }

        if eligible_partners.any?
          cheapest_partner, balance = find_cheapest_partner(eligible_partners, balance, size)

          final_price = Shared.calucate_final_price(cheapest_partner, size)
          result << [input_row.first, 'true', cheapest_partner['partner'], final_price]
        else
          result << [input_row.first, 'false', '', '']
        end
      end

      Shared.output_csv(result.sort, 'output2.csv')
    end

    def self.find_cheapest_partner(eligible_partners, balance, size)
      cheapest_partner = nil

      eligible_partners.each do |partner|
        next unless balance[partner['partner']] >= size

        cheapest_partner = partner
        balance[partner['partner']] = balance[partner['partner']] - size
        break
      end
      [cheapest_partner, balance]
    end
  end
end

Solution::ProblemTwo.execute
