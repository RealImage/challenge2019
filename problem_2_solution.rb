require "csv"
class Problem
  def solve
    partner = {}
    capacity = {}
    input = {}
    output = {}
    output_r = {}
    CSV.new(open('./partners.csv'), headers: :first_row).each do |element|
      partner[element["Theatre"].strip] = [] if partner[element["Theatre"].strip].nil?
      partner[element['Theatre'].strip] << {
        limit: element["Size Slab (in GB)"].strip,
        min_cost: element["Minimum cost"].strip,
        cost_per_gb: element["Cost Per GB"].strip,
        partner_id: element["Partner ID"].strip
      }
    end
    CSV.new(open('./capacities.csv'), headers: :first_row).each do |element|
      capacity[element["Partner ID"].strip]= element["Capacity (in GB)"].strip
    end
    CSV.new(open('./input.csv')).each do |row|
      input[row.first] = {size: row[1].to_i, theatre: row.last}
    end
    a = input
    a.reverse_each do |row|
      mincost = nil
      partner_id = nil
      partner[row.last[:theatre].strip].each do |data|
          element= data[:limit].split('-')
          if (element.first.to_i..element.last.to_i).include?(row.last[:size].to_i)&& row.last[:size].to_i <= capacity[data[:partner_id]].to_i  && !capacity[data[:partner_id]].nil? 
            price = data[:cost_per_gb].to_i * row.last[:size].to_i
            price = data[:min_cost].to_i if price < data[:min_cost].to_i
            if mincost.nil? || price < mincost
              mincost = price
              partner_id = data[:partner_id]
              output_r[row.first] = {
                  mincost: mincost,
                  partner_id: partner_id,
                  mincost_pre: true
                }                        
                end
            end
        end
        if !capacity[partner_id].nil?
          capacity[partner_id] = capacity[partner_id].to_i - row.last[:size].to_i
        end
        if mincost.nil?
          output_r[row.first] = {
            mincost:"\"\"",
            partner_id: "\"\"",
            mincost_pre: false
          }
        end

    end
    output = output_r.reverse_each
    output.each do |data|
      puts "#{data.first}, #{data.last[:mincost_pre]}, #{data.last[:partner_id]}, #{data.last[:mincost]}"
    end
end
end

Problem.new.solve 