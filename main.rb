require './fileIO/reader.rb'
require './services/parser.rb'
require './services/calculation_service.rb'
require './models/partner.rb'
require './fileIO/writer'

#Problem1
data = Reader.new('./fixtures/partners.csv', headers: true).read
parser = Parser.new(data)
partners = parser.parse_partners
input_data = Reader.new('./fixtures/input.csv').read
input_data = parser.clean_input_data(input_data)
calculation_service = CalculationService.new(partners, input_data)
results = calculation_service.calculate_minimal_cost
Writer.new('./fixtures/output1.csv', results).write

#Problem2
data = Reader.new('./fixtures/partners.csv', headers: true).read
parser = Parser.new(data)
partners = parser.parse_partners
capacities = Reader.new('./fixtures/capacities.csv', headers: true).read
parser.parse_partner_capacities(capacities)
input_data = Reader.new('./fixtures/input.csv').read
input_data = parser.clean_input_data(input_data)
calculation_service = CalculationService.new(partners, input_data)
calculation_service.calculate_minimal_cost
final_results = calculation_service.calculate_minimal_cost_by_capacity(capacities)
Writer.new('./fixtures/output2.csv', final_results).write

puts "CSV Processed Successfully. Please refer the output csv files in fixtures folder for results."
