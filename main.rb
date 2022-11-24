require 'csv'
require 'pry'

class Main
    def csv
        csv_output_arr = main
        CSV.open("myfile.csv", "w") do |csv|
            csv_output_arr.each do |p|
                csv << p
            end
        end
    end

    def main
        final_output_arr = []
        CSV.parse(File.read("input.CSV"), headers: false).each do |row|
            log = logic(row[1],row[2])  
            log = log[0] == "true" ? log.unshift(row[0]) : [row[0],"false", "", ""]
            final_output_arr << log
        end
        final_output_arr
    end

    def logic(slab, theatre)
        output_arr = []
        CSV.parse(File.read("partners.CSV"), headers: false).drop(1).each do |row|
            break if !(output_arr.empty?)
            next if row[0].strip != theatre.strip
            low,high = row[1].split("-").map(&:to_i)
            next if !(slab.to_i.between?(low, high))
            total_cost = slab.to_i * row[3].to_i 
            final_cost = total_cost < row[2].to_i ? row[2].to_i : total_cost 
            output_arr[0], output_arr[1], output_arr[2] = "true", row[4].strip, final_cost
        end
        output_arr
    end
end

obj1 = Main.new
obj1.csv
