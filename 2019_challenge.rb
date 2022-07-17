require 'csv'

options = { headers: true, skip_blanks: true, skip_lines: /^(?:,\s*)+$/ }

add_entry_to_csv = []
CSV.foreach('input.csv').with_index do |input_row, input_index|
  minimum_cost = ''
  minimum_parter = ''
  CSV.foreach('partners.csv', options).with_index do |partner_row, parter_index|
    if partner_row['Theatre'].strip == input_row[2] && partner_row['Size Slab (in GB)'].strip.split('-')[0].to_i <= input_row[1].to_i && partner_row['Size Slab (in GB)'].strip.split('-')[1].to_i >= input_row[1].to_i
      multiple_cost = input_row[1].to_i * partner_row['Cost Per GB'].to_i
      actual_cost = multiple_cost > partner_row['Minimum cost'].to_i ? multiple_cost : partner_row['Minimum cost'].to_i
      if minimum_cost == ''
        minimum_cost = actual_cost
        minimum_parter = partner_row['Partner ID']
      elsif actual_cost < minimum_cost
        minimum_parter = partner_row['Partner ID']
        minimum_cost = actual_cost
      end
    end
  end
  add_entry_to_csv << [input_row[0], minimum_parter != '', minimum_parter, minimum_cost]
end
CSV.open('output.csv', 'wb') do |csv|
  add_entry_to_csv.each do |entry|
    csv << entry
  end
end
