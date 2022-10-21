require_relative '../parsers/csv/parser'
require_relative '../parsers/csv/schemas'

module Services
  class DeliveriesProcessor
    INPUT_FILENAME = 'input.csv'
    PARTNERS_FILENAME = 'partners.csv'
    RESULT_FILENAME = 'result_1.csv'.freeze

    attr_reader :input_parser, :partners_parser

    def initialize(file_name = INPUT_FILENAME)
      @input_parser = Parsers::Csv::Parser.new(
        file_name: file_name,
        schema: Parsers::Csv::Schemas::INPUT_SCHEMA
      )
      @partners_parser = Parsers::Csv::Parser.new(
        file_name: PARTNERS_FILENAME,
        schema: Parsers::Csv::Schemas::PARTNERS_SCHEMA,
        headers: true
      )
    end

    def call
      delete_csv(RESULT_FILENAME)

      deliveries.each do |delivery_row|
        proposals_by_theatre = proposals_by_theatre(delivery_row, partners)
        proposal = proposals_by_theatre.min_by { |delivery_row| delivery_row[:price] }

        write_to_csv(delivery_row, proposal)
      end
    end

    private

    def deliveries
      @deliveries ||= input_parser.parse
    end

    def partners
      @partners ||= partners_parser.parse
    end

    def proposals_by_theatre(delivery_row, partners)
      partners.each_with_object([]) do |partner_row, partner_row_with_price|
        next if !proposal_available?(partner_row, delivery_row)

        delivery_price = delivery_row[:delivery_size] * partner_row[:gb_cost]
        delivery_price = partner_row[:min_cost] if delivery_price < partner_row[:min_cost]
      
        partner_row_with_price << partner_row.merge(price: delivery_price)
      end
    end

    def write_to_csv(delivery_row, partner_row)
      CSV.open('result_1.csv', 'a+') do |csv|
        csv << build_row(delivery_row, partner_row)
      end
    end

    def build_row(delivery_row, partner_row)
      return [delivery_row[:delivery_id], false, '', ''] if !partner_row

      [delivery_row[:delivery_id], true, partner_row[:partner_id], partner_row[:price]]
    end

    def proposal_available?(partner_row, delivery_row)
      partner_row[:theatre_id] == delivery_row[:theatre_id] &&
        partner_row[:slab_size].include?(delivery_row[:delivery_size])
    end

    def delete_csv(file_name)
      File.delete(file_name) if File.exists? file_name
    end
  end
end
