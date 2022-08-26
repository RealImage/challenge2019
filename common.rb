
class Common
  def Common.partners_hash
    partners_hash = []
    CSV.foreach(("partners.csv"), headers: true) do |row|
      partners_hash << { theater_id: row['Theatre'].strip, slab: row['Size Slab (in GB)'].strip, min_cost: row['Minimum cost'].strip, cost: row['Cost Per GB'].strip, partner_id: row['Partner ID'].strip }
    end
    partners_hash
  end

  def Common.capacities_hash
    capacities_hash = {}
    CSV.foreach(("capacities.csv"), headers: true) do |row|
      capacities_hash[row['Partner ID'].strip] = row['Capacity (in GB)'].strip.to_i
    end
    capacities_hash
  end

  def Common.input_arr
    input_arr = []
    CSV.foreach(("input.csv"), headers: false) do |row|
      input_arr << row
    end
    input_arr
  end
end
