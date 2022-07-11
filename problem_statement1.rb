require 'csv' 

 
partners = CSV.foreach('partners.csv', headers: :first_row).map do |p|
	row = {}
	p.each {|k,v| row[k.strip.downcase.gsub(' ','_').to_sym] = v.strip}
	row
end

input = CSV.foreach('input.csv').map(&:to_a)


output = {}

input.each do |d, s, t|
	output[d] = [d, false, "", ""]
	theatres = partners.select {|p| p[:theatre] == t}
	theatres.each do |row|
	 	slab = row[:"size_slab_(in_gb)"].split('-')
	 	if s.to_i >= slab[0].to_i && s.to_i <= slab[1].to_i
	 		cost = s.to_i * row[:cost_per_gb].to_i
	 		cost = row[:minimum_cost].to_i if cost < row[:minimum_cost].to_i
	 		flag = output[d][3].to_i == 0 || output[d][3] > cost

		 	output[d] = [d, true, row[:partner_id], cost] if flag	 	
	 	end
	end
end

p output.values
