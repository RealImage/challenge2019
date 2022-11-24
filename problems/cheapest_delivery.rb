require 'csv';

module Problems
  class CheapestDelivery
    def initialize
      @data = formatted_data
    end

    def report
      CSV.open('output1.csv', 'w') do |csv|
        fetch_required_rows.each do |required_row|
          csv << required_row
        end
      end
    end

    private

    def fetch_required_rows
      required_data = []
      CSV.parse(File.read('input.csv')) do |delivery_row|
        deliveries = filter_data_for_delivery(delivery_row)
        default_row = [delivery_row[0], false, '', '']
        if deliveries.values.empty?
          required_data.push(default_row)
          next
        end

        cheapest_record = find_cheapest_of_deliveries(delivery_row[1], deliveries.values)
        if cheapest_record
          required_data.push([delivery_row[0], true, cheapest_record[:partner], cheapest_record[:actual_cost]])
        else
          required_data.push(default_row)
        end
      end
      required_data
    end

    def filter_data_for_delivery(delivery)
      @data.select do |key, value|
        key[0] == delivery[2] && delivery[1].between?(key[1], key[2])
      end
    end

    def find_cheapest_of_deliveries(size, filtered_data)
      filtered_data.map do |data|
        cost = find_cost_for_given_size(size, data[:cost_per_gb], data[:minimum_cost])
        data.merge!({ actual_cost: cost })
      end

      filtered_data.sort_by! {|x| [x[:partner], x[:actual_cost]]}
      filtered_data[0]
    end

    def find_cost_for_given_size(size, cost_per_gb, minimum_cost)
      total_cost = size.to_i * cost_per_gb.to_i
      total_cost <= minimum_cost.to_i ? minimum_cost.to_i : total_cost
    end

    def formatted_data
      hash = {}
      table = CSV.parse(File.read('partners.csv'), headers: true) do |row|
        size_lab = row[1].split('-')
        hash[[row[0], size_lab[0], size_lab[1]]] = {
          minimum_cost: row[2],
          cost_per_gb: row[3],
          partner: row[4]
        }
      end
      hash
    end
  end
end

Problems::CheapestDelivery.new.report
