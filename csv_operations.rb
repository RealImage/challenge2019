# frozen_string_literal: true

require 'English'
require 'csv'
module CsvOperations
  def load_csv(file:)
    csv = File.read file.strip
    head, *rest = csv.split($INPUT_RECORD_SEPARATOR).map do |row|
      row.split(/["']*\s*,\s*['"]*/).map do |col|
        col.tr("\r\n", '').delete('\\"')
      end
    end
    rest.map { |row| head.zip(row).to_h }
  end

  def load_input_csv
    inputs = []
    CSV.foreach('input.csv', headers: false) do |row|
      input_row = {}
      input_row['Delivery ID'] = row[0]
      input_row['Delivery Size'] = row[1]
      input_row['Theatre ID'] = row[2]
      inputs.push(input_row)
    end
    inputs
  end
end
