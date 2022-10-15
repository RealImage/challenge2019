require_relative '../file_parser'
require 'csv'

module Parsers
  module Csv
    class Parser < Parsers::FileParser
      def parse
        CSV.foreach(file_name, parser_opts) do |row|
          content << parsed_row(row)
        end
        content
      end

      private

      def parser_opts
        {
          headers: !!opts[:headers],
          col_sep: opts[:col_sep] || ','
        }
      end

      def parsed_row(row)
        opts[:schema].each_with_object({}) do |(field_name, field_params), h|
          h[field_name] = format_field(row[field_params[:index]], field_params)
        end
      end

      def format_field(field, field_params)
        striped_field = field.strip

        formated_field = case field_params[:type].to_s
        when 'Integer'
          striped_field.to_i
        when 'Range'
          Range.new(*striped_field.split(/\-/).map(&:to_i))
        else
          striped_field
        end
      end
    end
  end
end
