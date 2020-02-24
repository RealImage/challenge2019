require 'csv'
    csv_data = CSV.read(ARGV[0])
    data_set = []
    csv_data.each_with_index do |data, index|
      next if index.zero?
      data_set << {
          theatre: data[0].strip,
          size_slab: data[1].strip.split('-').map(&:to_i),
          min_cost: data[2].strip.to_i,
          cost_per_gb: data[3].strip.to_i,
          partner_id: data[4].strip
      }
    end
    capacities = CSV.read(ARGV[1])

    partner_capacities = []
    capacities.each_with_index do |cap_data, index1|
      next if index1.zero?
      partner_capacities << {
          partner_id: cap_data[0].strip,
          capacities_gb: cap_data[1]
      }
    end

    input_data_1 = CSV.read(ARGV[2])
    deliveries = []
    input_data_1.each do |input|
      deliveries << {
          delivery: input[0],
          size: input[1].to_i,
          theatre: input[2]
      }
    end
    CSV.open(ARGV[3], "a+") do |output|
      output << %w[Delivery Possibility Partner Cost]
      deliveries.each_with_index do |delivery|
        cost = ""
        possibility = false
        partner = ""
        valid_by_theatre = data_set.select { |key| key[:theatre] == "#{delivery[:theatre]}" }
        data_size = delivery[:size]
        @valid_by_size = valid_by_theatre.select { |key| data_size.between?(key[:size_slab][0], key[:size_slab][1]) }
        if !@valid_by_size.empty?
          cost_for_each = @valid_by_size.map { |data| [{"#{data[:partner_id]}": data[:cost_per_gb] * delivery[:size]}, {:mincost => data[:min_cost]}] }
          first_min_cost = cost_for_each.first
          final_min_cost = cost_for_each.map { |a| a.select { |b| b.values.last < first_min_cost.map { |a| a.values }.first.last if b.keys == [:mincost] } }.select { |c|  c if !c.empty? }.first
          final_pair = cost_for_each.select { |a|  a.include?(final_min_cost.first) }.flatten.first
          if partner_capacities.select { |a| a if a.values.first == final_pair.keys.last.to_s }.map { |a| a[:capacities_gb] }.first.to_i > delivery[:size].to_i
            partner = final_pair.keys.last
            possibility = true
            cost = final_pair.values.last
          end
        end
        output << ["#{delivery[:delivery]}", "#{possibility}", "#{partner}", "#{cost}"]
      end
    end

