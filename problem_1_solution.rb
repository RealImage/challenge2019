require 'csv'
class Problem
    def solve
        mincost= nil
        partner_id = nil
        partner = {}
        CSV.new(open('./partners.csv'), headers: :first_row).each do |element|

            partner[element["Theatre"].strip] = [] if partner[element["Theatre"].strip].nil?
            partner[element['Theatre'].strip] << {
                limit: element["Size Slab (in GB)"].strip,
                min_cost: element["Minimum cost"].strip,
                cost_per_gb: element["Cost Per GB"].strip,
                partner_id: element["Partner ID"].strip
            }

        end
        CSV.new(open('./input.csv')).each do |row|
            a= row.map{ |element| element.strip }
            mincost = nil
            partner_id = nil
            partner[row.last].each do |data|
                element= data[:limit].split('-')
                if (element.first.to_i..element.last.to_i).include?(a[1].to_i)
                    price = data[:cost_per_gb].to_i * a[1].to_i
                    price = data[:min_cost].to_i if price < data[:min_cost].to_i
                    if mincost.nil? || price < mincost
                        mincost = price
                        partner_id = data[:partner_id]
                    end
                end
            end
            if !mincost.nil?
                puts " #{a[0]}, #{!mincost.nil?}, #{partner_id}, #{mincost}"
            else
                puts " #{a[0]}, #{!mincost.nil?}, \"\", \"\"  "
            end
        end
    end
end
Problem.new.solve