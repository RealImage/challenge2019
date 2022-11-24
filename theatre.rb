require 'csv'

class Qube
  def resolve_problem_1
    partners = get_partners
    deliveries = input_data
    output = []

    deliveries.each do |delivery|
      minimum_cost_partner = ["", Float::INFINITY]
      delivery_status = false
      # [0, 100, 1500, 20, "P1"]
      partners[delivery[2]].each do |partner_delivery|
        if delivery[1] > partner_delivery[0] && delivery[1] < partner_delivery[1]
          delivery_status = true

          total_cost = partner_delivery[3] * delivery[1]
          cost = total_cost > partner_delivery[2] ? total_cost : partner_delivery[2]

          minimum_cost_partner = [partner_delivery.last, cost] if cost < minimum_cost_partner.last
        end
      end

      if delivery_status
        output << [delivery[0], true] + minimum_cost_partner
      else
        output << [delivery[0], false, "", ""]
      end
    end

    write_output("output.csv", output)
  end

  def input_data
    array_deliveries = CSV.parse(File.read('input.csv'), :headers => false)
    array_deliveries.map { |array| [array[0], array[1].to_i, array[2]] }
  end

  def get_partners
    table_partners = CSV.parse(File.read('partners.csv'), :headers => true)
    theatre_partners = {}

    table_partners.each do |row|
      size = row["Size Slab (in GB)"]&.strip&.split("-")
      theatre_partners[row["Theatre"].strip] = [] if theatre_partners[row["Theatre"].strip].nil?
      theatre_partners[row["Theatre"].strip] << [size.first.to_i, size.last.to_i,
        row["Minimum cost"].to_i, row["Cost Per GB"].to_i, row["Partner ID"].strip]
    end

    theatre_partners
  end

  def write_output(path_csv, arrays)
    CSV.open( path_csv, 'w' ) do |f|
      arrays.each do |array|
        f << array
      end
    end
  end

  # [["P1", 350], ["P2", 500], ["P3", 1500]]
  def get_capacities
    table_capacities = CSV.parse(File.read('capacities.csv'), :headers => true)
    capacities = []

    table_capacities.each do |capacity|
      capacities << [capacity["Partner ID"].strip, capacity["Capacity (in GB)"].strip.to_i]
    end

    capacities
  end
end

qube = Qube.new
resolve_1 = qube.resolve_problem_1
