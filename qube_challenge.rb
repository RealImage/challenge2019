require "csv"

class QubeCinema
    
    # Method to check if delivery is possible or not. If possible then returns the lowest cost and the partner
    def check_delivery_status
        # Gets the partner details from csv file
        partner_details = CSV.read("./partners.csv")
        processed_partner_details = []
        final_output = []
        partner_details.shift
        partner_details.each do |details|
            processed_partner_details << {"Theatre"=>details[0].strip, "Details"=>details[1..4].map{|data| data.strip}}
        end
        # Gets the delivery details from csv file
        delivery_details = CSV.read("./input.csv")
        delivery_details.each do |details|
            total_data = details[1].to_i
            theatres = processed_partner_details.select {|data| data["Theatre"] == details[2]}
            if theatres.count > 0
                output = []
                lowest_rate = 0
                theatres.each do |theatre|
                    theatre_data = theatre["Details"]
                    minimum_size, maximum_size = theatre_data[0].split("-").map{|size| size.to_i}
                    if minimum_size < total_data && total_data < maximum_size
                        minimum_cost, per_cost = theatre_data[1..2].map(&:to_i)
                        rate = total_data * per_cost
                        current_rate = (rate > minimum_cost ? rate : minimum_cost)
                        if lowest_rate == 0 || lowest_rate > current_rate
                            lowest_rate = current_rate
                            output = [details[0].strip, true, theatre_data[3], current_rate]
                        end
                    elsif output.length == 0
                        output = [details[0].strip, false, "", ""]
                    end
                end
                final_output << output if output.length > 0
            else
                final_output << [details[0].strip, false, "", ""]
            end
        end
        # Generates the result in a csv file
        CSV.open("./output.csv", "wb") do |csv|
            final_output.each do |row|
                csv << row
            end
        end
    end
    
end

QubeCinema.new.check_delivery_status