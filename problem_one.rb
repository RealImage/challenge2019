# frozen_string_literal: true

require './shared'

module Solution
  class ProblemOne
    def self.execute
      partners, _capacities = Shared.initialize_csv
      input = Shared.initialize_input
      result = []

      input.each do |input_row|
        size = input_row[1]
        theatre_id = input_row.last

        eligible_partners = Shared.eligible_partners(partners, theatre_id, size)

        if eligible_partners.any?
          cheapest_partner = eligible_partners.min_by { |x| x['cost'] }
          final_price = Shared.calucate_final_price(cheapest_partner, size)
          result << [input_row.first, 'true', cheapest_partner['partner'], final_price]
        else
          result << [input_row.first, 'false', '', '']
        end
      end

      Shared.output_csv(result, 'output1.csv')
    end
  end
end

Solution::ProblemOne.execute
