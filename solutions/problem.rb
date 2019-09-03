module Problem
  def problem1
    delivery_patners = get_delivery_patners
    deliveries = get_deliveries
    output = []
    deliveries.each_with_index do |delivery|
      minimum_value_cost = ["", Float::INFINITY]
      delivery_status = false
      delivery_patners[delivery[2]].each do |delivery_patner|
        if delivery_patner[0] < delivery[1] && delivery[1] < delivery_patner[1];
          delivery_status = true
          actual_cost = delivery[1] * delivery_patner[3]
          cost = actual_cost > delivery_patner[2] ? actual_cost : delivery_patner[2]
          minimum_value_cost = [delivery_patner.last, cost] if cost < minimum_value_cost[1]
        end
      end
      if delivery_status
        output << [delivery[0], delivery_status] + minimum_value_cost
      else
        output << [delivery[0], delivery_status, "", ""]
      end
    end
    write_output("output1.csv", output)
    output
  end

  def problem2
    capacities = get_capacities
    delivery_patners = get_delivery_patners
    deliveries = get_deliveries
    output = []
    total_cost = []
    deliveries.each do |delivery|
      minimum_value_cost = ["", Float::INFINITY]
      delivery_status = false
      temp = []
      delivery_patners[delivery[2]].each do |delivery_patner|
        if delivery_patner[0] < delivery[1] && delivery[1] < delivery_patner[1];
          delivery_status = true
          actual_cost = delivery[1] * delivery_patner[3]
          cost = actual_cost > delivery_patner[2] ? actual_cost : delivery_patner[2]
          temp << [delivery[0], delivery_status, delivery_patner.last, cost, delivery[1]]
          minimum_value_cost = [delivery_patner.last, cost] if cost < minimum_value_cost[1]
        end
      end
      if delivery_status
        output << [delivery[0], delivery_status] + minimum_value_cost
        total_cost << temp if temp.any?
      else
        output << [delivery[0], delivery_status, "", ""]
      end
    end

    possible_array = output_with_capacity(possible_combination(total_cost))
    output_array = []
    index = 0
    deliveries.each do |delivery|
      if delivery.first == possible_array[index]&.first
        output_array << possible_array[index].first(4)
        index += 1
      else
        output_array << [delivery[0], false, "", ""]
      end
    end
    write_output("output2.csv", output_array)
    output_array
  end

  def possible_combination(total_cost)
    combinations = []
    total_cost.each do |array|
      unless combinations.any?
        combinations = array
      else
        temp = []
        array.each do |a|
          combinations.each do |x|
            temp << x + a
          end
        end
        combinations = temp
      end
    end
    combinations
  end

  def output_with_capacity(combination_array)
    minimum_value = Float::INFINITY
    output_array = []
    combination_array.each do |arrays|
      value = 0
      status = true
      capacities = get_capacities
      arrays.each_slice(5) do |array|
        capacities[array[2]] -= array[4]
        if capacities[array[2]] < 0
          status = false
          break;
        else
          value += array[3]
        end
      end
      if status
        if value < minimum_value
          minimum_value = value
          output_array = arrays
        end
      end
    end
    output_array.each_slice(5).to_a
  end

  def get_capacities
    capacities_csv = File.read('capacities.csv')
    capacities_array = CSV.parse(capacities_csv, :headers => true )
    capacities = {}
    capacities_array.each { |array|  capacities[array["Partner ID"].strip] = array["Capacity (in GB)"].to_i }
    capacities
  end

  def get_delivery_patners
    partners_csv = File.read('partners.csv')
    partners_array = CSV.parse(partners_csv, :headers => true )
    delivery_patners = {}
    partners_array.each do |array|
      size = array["Size Slab (in GB)"].strip.split('-')
      delivery_patners[array["Theatre"].strip] = [] if delivery_patners[array["Theatre"].strip].nil?
      delivery_patners[array["Theatre"].strip] << [size.first.to_i, size.last.to_i, array["Minimum cost"].to_i, array["Cost Per GB"].to_i, array["Partner ID"].strip]
    end
    delivery_patners
  end

  def get_deliveries
    deliveries_csv = File.read('input.csv')
    deliveries_array = CSV.parse(deliveries_csv, :headers => false )
    deliveries_array.map { |array| [array[0], array[1].to_i, array[2]] }
  end

  def write_output(name, arrays)
    output_csv = name
    CSV.open( output_csv, 'w' ) do |writer|
      arrays.each do |array|
        writer << array
      end
    end
  end
end

include Problem
require 'csv'
problem1
problem2
