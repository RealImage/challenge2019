require 'csv'
require 'byebug'

partners = CSV.read('/home/development/Challenges/partners.csv')
limits = CSV.read('/home/development/Challenges/capacities.csv')
input_data = CSV.read('/home/development/Challenges/input.csv')

require 'smarter_csv'

IntegerConverter = Object.new

def IntegerConverter.convert(value)
  Integer(value)
end

def range_converter(data)
	data.map do |dat|
		range = dat[:"size_slab_(in_gb)"].split('-').map(&:to_i)
	    dat[:"size_slab_(in_gb)"]  = (range[0]..range[1])
	end
end

def best_price(rate, ans)
	price  =  ans.map{|i| i[:cost_per_gb] * rate.to_i }.min
	ans.select{|val|  val[:cost_per_gb] ==  (price/rate.to_i)}
end
result = []
data = SmarterCSV.process('/home/development/Challenges/partners.csv', value_converters: { minimum_cost: IntegerConverter })
range_converter(data)
input_data.map do |input|
	ans = data.select{|k| k[:theatre] == input[2] and k[:"size_slab_(in_gb)"].member? input[1].to_i}
	ans = best_price(input[1], ans)
    
    unless ans.empty?
     result << [input[0], TRUE, ans[0][:partner_id],  (input[1].to_i)*ans[0][:cost_per_gb] > ans[0][:minimum_cost] ? (input[1].to_i)*ans[0][:cost_per_gb] : ans[0][:minimum_cost]]
    else
     result << [input[0], FALSE, "" , ""]
    end
end


CSV.open("/home/development/Challenges/output1.csv", "w") do |csv|
	result.map {|row| csv << row }
	
end

