require 'csv'

partners = CSV.parse(File.read("partners.csv"))

inputs = CSV.parse(File.read("input.csv"))

CSV.open("output.csv", "w") do |csv|
    inputs.each do |input|
        distributor = input[0]
        size = input[1].to_i
        flag = false
        output = [distributor]
        min_cost = 999999999999999
        final_partner = " "
        flag = false
        partners[1..-1].each do |partner|
            if (input[2].strip == partner[0].strip)
                size_low = partner[1].strip.split("-").first.to_i
                size_high = partner[1].strip.split("-").last.to_i
                if (size >= size_low) && (size <= size_high)
                    new_min_cost = [size * partner[3].to_i, partner[2].to_i].max
                    if new_min_cost < min_cost
                        min_cost = new_min_cost
                        final_partner = partner[4]
                        flag = true
                    end
                end
            end
        end
        min_cost = " " if min_cost == 999999999999999
        output << flag
        output << final_partner
        output << min_cost
        csv << output
    end
end

