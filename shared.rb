# frozen_string_literal: true

require 'csv'

module Solution
  class Shared
    def self.initialize_csv
      %w[partners capacities].map do |file_name|
        csv = Solution::Shared.csv_parser("./#{file_name}.csv", headers: true)
        csv.map(&:to_h)
      end
    end

    def self.initialize_input
      Solution::Shared.csv_parser('./input.csv')
    end

    def self.csv_parser(file, headers: false)
      CSV.parse(
        File.read(file),
        headers: headers,
        header_converters: ->(h) { h.strip.split(' ').first.downcase },
        converters: [->(d) { d.strip }, :numeric]
      )
    end

    def self.output_csv(result, file)
      CSV.open(file, 'w') do |csv|
        result.each do |row|
          csv << row
        end
      end
    end

    def self.eligible_partners(partners, theatre_id, size)
      partners.select do |partner|
        partner['theatre'] == theatre_id && size.between?(partner['size'].split('-').first.to_i,
                                                          partner['size'].split('-').last.to_i)
      end
    end

    def self.calucate_final_price(partner, size)
      calucated_price = partner['cost'] * size
      calucated_price < partner['minimum'] ? partner['minimum'] : calucated_price
    end
  end
end
