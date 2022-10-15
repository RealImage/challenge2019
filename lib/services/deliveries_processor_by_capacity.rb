require_relative 'deliveries_processor'

module Services
  class DeliveriesProcessorByCapacity < DeliveriesProcessor
    CAPACITIES_FILENAME = 'capacities.csv'
    RESULT_FILENAME = 'result_2.csv'.freeze

    attr_reader :capacities_parser, :result

    def initialize(file_name = INPUT_FILENAME)
      super

      @capacities_parser = Parsers::Csv::Parser.new(
        file_name: CAPACITIES_FILENAME,
        schema: Parsers::Csv::Schemas::CAPACITY_SCHEMA,
        headers: true
      )
      @result = []
    end

    def call
      delete_csv(RESULT_FILENAME)

      deliveries.each do |delivery_row|
        proposals_by_theatre = proposals_by_theatre(delivery_row, partners)
        delivery_row[:proposals] = proposals_by_theatre.sort_by { |proposal| proposal[:price] }

        result << delivery_row
      end

      process_proposals_by_capacity

      result.each do |delivery|
        write_to_csv(delivery)
      end
    end

    private

    def capacities
      @capacities ||= capacities_parser.parse
    end

    def process_proposals_by_capacity
      capacities.each do |hash|
        partner_deliveries = result.select do |delivery|
          delivery[:proposals].any? &&
            delivery[:proposals][0][:partner_id] == hash[:partner_id]
        end

        next if partner_deliveries.empty?

        sum_delivery_sizes = partner_deliveries.sum { |h| h[:delivery_size] }

        if sum_delivery_sizes <= hash[:capacity]
          next
        else
          partner_deliveries[0][:proposals].delete_if do |proposal|
            proposal[:partner_id] == hash[:partner_id]
          end

          process_proposals_by_capacity
        end
      end
    end

    def proposals_by_theatre(delivery_row, partners_rows)
      partners_rows.each_with_object([]) do |partner_row, partner_row_with_price|
        next if !proposal_available?(partner_row, delivery_row)

        delivery_price = delivery_row[:delivery_size] * partner_row[:gb_cost]
        delivery_price = partner_row[:min_cost] if delivery_price < partner_row[:min_cost]
      
        partner_row_with_price << partner_row.merge(price: delivery_price)
      end
    end

    def write_to_csv(delivery)
      CSV.open('result_2.csv', 'a+') do |csv|
        csv << build_row(delivery)
      end
    end

    def build_row(delivery)
      return [delivery[:delivery_id], false, '', ''] if delivery[:proposals].empty?

      best_proposal = delivery[:proposals][0]

      [delivery[:delivery_id], true, best_proposal[:partner_id], best_proposal[:price]]
    end

    def proposal_available?(partner_row, delivery_row)
      partner_row[:theatre_id] == delivery_row[:theatre_id] &&
        partner_row[:slab_size].include?(delivery_row[:delivery_size])
    end
  end
end
