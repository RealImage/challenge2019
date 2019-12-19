require 'csv'

partners = CSV.read("partners.csv")
input = CSV.read("input.csv")

solution = Array.new(input.count){Array.new(4,0)}


i=0
input.each do | d_id|

	partners.each do |partner|

		solution[i][0] = d_id[0]

		if d_id[2].strip == partner[0].strip then
			#p d_id[2] + " " +partner[0]

			if d_id[1].to_i > partner[1].split("-")[0].to_i && d_id[1].to_i < partner[1].split("-")[1].to_i then
			#	p d_id[1] + " " + partner[1]

				cost = d_id[1].to_i * partner[3].to_i

				if cost < partner[2].to_i then cost = partner[2].to_i end

				if cost < solution[i][3] || solution[i][3]==0 then
					solution[i][1] = "true "
					solution[i][2] = partner[4]
					solution[i][3] = cost 
				end
			end

		end

	end
		if solution[i][1] == 0 then solution[i][1] = false; solution[i][2]="";solution[i][3]=""; end
	i = i + 1
end


	CSV.open("output.csv", "w") do |csv|
		solution.each do |s|
  		csv << [s[0],s[1],s[2],s[3]]
		end
	end

