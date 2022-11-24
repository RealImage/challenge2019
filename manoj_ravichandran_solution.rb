require 'csv'

def partners_csv
  CSV.open("partners.csv")
end

def input_csv 
  CSV.open("input.csv").reject { |row| row.all?(&:nil?) }.map(&:compact) 
end

def capacity_csv 
  CSV.open("capacities.csv").reject { |row| row.all?(&:nil?) }.map(&:compact) 
end

def write_output1(output) 
  CSV.open("output1.csv", "w") do |csv|
    output.each do |out|
      csv << out
    end
  end
end

def write_output2(output) 
  CSV.open("output2.csv", "w") do |csv|
    output.each do |out|
      csv << out
    end
  end
end

def populatePartnerHash
  partner_hash = {}
  partners_csv.each_with_index do |row, i|
    next if i == 0
    partner_hash[row[0]] ||= []
    partner_hash[row[0]] << {
      min: row[1].split("-")[0],
      max: row[1].split("-")[1],
      min_cost: row[2],
      costpg: row[3],
      pid: row[4] 
    }
  end
  partner_hash
end

def populateCapacityHash
  capacity = {}
  capacity_csv.each_with_index do |row, i|
    next if i == 0
    capacity[row[0]] = row[1].to_i
  end
  capacity
end

def calculateMinimumCostDelivery(partner_hash)
  output = []
  input_csv.each do |inp|
    tid = inp[2]
    size = inp[1].to_i
    min_cost_for_delivery = 100000000
    current_cost = 0
    min_cost_partner = nil
    current_theatre = partner_hash[tid]

    current_theatre.each do |partner|
      if size >= partner[:min].to_i && size <= partner[:max].to_i
        current_cost = [size * partner[:costpg].to_i, partner[:min_cost].to_i].max
        if current_cost < min_cost_for_delivery
          min_cost_for_delivery = current_cost 
          min_cost_partner = partner[:pid]
        end
      end
    end

    if min_cost_partner
      output << [tid, true, min_cost_partner, min_cost_for_delivery, size]
    else
      output << [tid, false, "", "", size]
    end
  end
  output
 end

def findMinPartnerTheatre(output, partner)
  minSize = 1000000000
  curSize = 0
  minIndex = nil
  output.each_with_index do |out, index|  
    if partner == out[2]
      curSize = out[4]
      if curSize <= minSize
        minSize = curSize
        minIndex = index
      end 
    end
  end
  minIndex
end

def recomputeMinPartnerTheatre(output, index, skip_partner, partner_hash, input_csv)
  theatre = output[0]
  current_theatre = partner_hash[theatre[0]]
  min_cost_for_delivery = 100000000
  current_cost = 0

  input_csv.each do |inp|
    size = inp[1].to_i
    min_cost_for_delivery = 100000000
    current_cost = 0
    min_cost_partner = nil
    next if inp[2] != theatre[0]
    current_theatre.each do |partner|
      next if partner[:pid] == skip_partner
      if size >= partner[:min].to_i && size <= partner[:max].to_i
        current_cost = [size * partner[:costpg].to_i, partner[:min_cost].to_i].max
        if current_cost < min_cost_for_delivery
          min_cost_for_delivery = current_cost 
          min_cost_partner = partner[:pid]
        end
      end
    end
    if min_cost_partner
      output[index] = [theatre[0], true, min_cost_partner, min_cost_for_delivery, size]
    else
      output[index] = [theatre[0], false, "", "", size]
    end
    break
  end
  output
end

def calculateMinimumCostDeliveryWithCapacity(output, partner_hash, capacity)
  total_size = {}
  output.each do |out|
    partner = out[2]
    size = out[4]
    delivery_possible = out[1]

    if delivery_possible
      total_size[partner] ||= 0
      total_size[partner] += size
      if total_size[partner] > capacity[partner]
        minIndex = findMinPartnerTheatre(output, partner)
        recomputeMinPartnerTheatre(output, minIndex, partner, partner_hash, input_csv)
      end
    end
  end
end

partner_hash = populatePartnerHash
capacity = populateCapacityHash
output = calculateMinimumCostDelivery(partner_hash)
write_output1(output)
output = calculateMinimumCostDeliveryWithCapacity(output, partner_hash, capacity)
write_output2(output)

