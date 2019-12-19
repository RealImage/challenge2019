require 'csv'

partners = CSV.read("partners.csv")
input = CSV.read("input.csv")
capacities = CSV.read("capacities.csv")


solution = Array.new(input.count){Array.new(5,0)}
solutions = Array.new(input.count){Array.new(5,0)}

solutions_matrix = Array.new(input.count) { Array.new(1)  }

# Finding all the solutions
i=0
j=0
input.each do | d_id|

	partners.each do |partner|

		solution[i][0] = d_id[0]

		if d_id[2].strip == partner[0].strip then

			if d_id[1].to_i > partner[1].split("-")[0].to_i && d_id[1].to_i < partner[1].split("-")[1].to_i then

				cost = d_id[1].to_i * partner[3].to_i

				if cost < partner[2].to_i then cost = partner[2].to_i end
					solutions[j] = d_id[0] +" "+ d_id[1] + " true " +partner[4]+" "+cost.to_s
					j = j + 1
				if cost < solution[i][3] || solution[i][3]==0 then
					#solution[i][0] = d_id[0]
					solution[i][1] = true
					solution[i][2] = partner[4]
					solution[i][3] = cost 
					solution[i][4] = d_id[1].to_i
				end
			end

		end

	end
		if solution[i][1] == 0 then solutions[j] = d_id[0] +" "+ d_id[1] + " false"; j=j+1 end
	i = i + 1
end
# testing the above code, evaluating the code
# solutions.each {|s| p s}


# Creating Matrix of all found Solutions
	i = 0
	input.each do |input|
		solutions_matrix[i].shift
		solutions.grep(/^#{input[0]}/).each {|s| solutions_matrix[i]<<s}
		i = i + 1
	end
	# testing the above code, evaluating the code
	# solutions_matrix.each {|s| p s}



# Getting solutions under Partner Capacity	
capacities.each	do |capacity|

	par = capacity[0].strip
	par_cap = capacity[1].to_i

		x = Array.new(0);  i=0;
	solutions_matrix.each {|s| if s[0].split(" ").grep(/#{par}/).count==1 then x<<i end; i = i+1}

	 par_load = 0
	 x.each { |pos| par_load = par_load + solutions_matrix[pos][0].split(" ")[1].to_i }

	

	 if par_load > par_cap then
	 	if x.count == 2 then
	 		if solutions_matrix[x[0]][0].split(" ")[1].to_i	> solutions_matrix[x[1]][0].split(" ")[1].to_i then solutions_matrix[x[1]].shift else solutions_matrix[x[0]].shift end	
	 	end
	 	if x.count == 3 then
	 			x_0 = solutions_matrix[x[0]][0].split(" ")[1].to_i
	 			x_1 = solutions_matrix[x[1]][0].split(" ")[1].to_i
	 			x_2 = solutions_matrix[x[2]][0].split(" ")[1].to_i
	 		if ((x_0+x_1) < par_cap) && ((x_0+x_1)> (x_0+x_2)) && ((x_0+x_1)>(x_2+x_1)) then solutions_matrix[x[2]].shift else solutions_matrix[x[1]].shift; solutions_matrix[x[0]].shift end
	 		if ((x_0+x_2) < par_cap) && ((x_0+x_2)> (x_0+x_1)) && ((x_0+x_2)>(x_2+x_1)) then solutions_matrix[x[1]].shift else solutions_matrix[x[0]].shift; solutions_matrix[x[2]].shift end
			if ((x_2+x_1) < par_cap) && ((x_2+x_1)> (x_0+x_2)) && ((x_2+x_1)>(x_0+x_1)) then solutions_matrix[x[0]].shift else solutions_matrix[x[1]].shift; solutions_matrix[x[2]].shift end
	 	end

	 end


end	

# testing the above code, evaluating the code
#solutions_matrix.each {|s| p s}

str = Array.new(0)
solutions_matrix.each{|s| str << s[0].split(" ")}


 # Formatting as per the Output requirement 
 i = 0
 str.each do |s|
 	if s[2]=="true" then str[i][2]="true "  else str[i][3]="";str[i][4]="" end
 		i = i+1
 end

 
# testing the above code, evaluating the code
# str.each {|s| p s}


CSV.open("output.csv", "w") do |csv|
		str.each do |s|
  		csv << [s[0],s[2],s[3],s[4]]
		end
	end


