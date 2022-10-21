require 'csv'
require 'byebug'
class QubeChallenge
    def solution1
        all_partners = load_csv_data
        output_array = []
        CSV.new(open('./input.csv')).each do |row|
            # row[0] is delivery ID, row[1] is size row[2] id threatre
            theatre_partners = all_partners.select{|p| p["Theatre"]== row[2]}

            theatre_partners.each_with_index do |tp,index|
                slab_size = row[1]&.to_i
                d_id = row[0]
                size_min,size_max = tp["Size Slab (in GB)"].split("-")
                minimum_cost = tp["Minimum cost"]&.to_i
                price = tp["Cost Per GB"]&.to_i
                total_cost = slab_size * price
                final_cost = total_cost < minimum_cost ? minimum_cost : total_cost
                if (size_min.to_i..size_max.to_i).include?(slab_size)
                    output_array << [d_id,true,tp["Partner ID"],final_cost]
                    break
                elsif index == (theatre_partners.size - 1)
                    output_array << [d_id,false,"",""]
                end
            end
            output_array << [d_id,false,"",""] if theatre_partners.empty?
        end
        generate_csv(output_array)
    end

    def load_csv_data
        csv = CSV.new(open('./partners.csv'), headers: :first_row, converters: ->(f) { f&.strip })
        csv =  csv.to_a.map {|row| row.to_hash }
    end

    def generate_csv(csv_data)
        CSV.open('output.csv', 'w') do |csv|
          csv_data.each { |ar| csv << ar }
        end
    end
end
QubeChallenge.new.solution1