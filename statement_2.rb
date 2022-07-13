require 'csv'

def main
  partners = read_constant_data('partners')
  capacities = read_constant_data('capacities')

  packages = input_content.map do |deliver|
    get_partner(deliver, partners)
  end

  results = handle_capacity(packages, capacities)

  File.write('output2.csv', results.map(&:to_csv).join)
end

def input_content
  CSV.open('./input.csv').each.to_a.compact.map do |content|
    content.map(&:strip)
  end
end

def read_constant_data(file_name)
  data = []
  CSV.foreach("./#{file_name}.csv", headers: true) do |row|
    data << if file_name == 'partners' 
      {
        theatre: row['Theatre'].strip,
        size_lab_min: row['Size Slab (in GB)'].strip.split('-')[0].to_i,
        size_lab_max: row['Size Slab (in GB)'].strip.split('-')[1].to_i,
        minimun_cost: row['Minimum cost'].to_i,
        cost_per_gb: row['Cost Per GB'].to_i,
        partner_id: row['Partner ID'].strip
      }
    else
      {
        partner_id: row['Partner ID'].strip,
        capacity: row['Capacity (in GB)'].to_i
      }
    end
  end
  data
end

def get_partner(deliver, partners)
  csv_out_puts = []
  theatres = partners.select { |content| content[:theatre] == deliver[2] && content[:size_lab_min] <= deliver[1].to_i && content[:size_lab_max] >= deliver[1].to_i }

  if theatres.count > 0
    delivery_costs = theatres.map do |theatre|
      const_temp = deliver[1].to_i * theatre[:cost_per_gb]

      {
        deliver_id: deliver[0],
        possible: true,
        size: deliver[1].to_i,
        partner_id: theatre[:partner_id],
        cost: const_temp >= theatre[:minimun_cost] ? const_temp : theatre[:minimun_cost]
      }
    end

    delivery_costs.sort { |a| -a[:cost] }
  else
    [deliver[0], false, '', '']
  end
end

def handle_capacity(packages, capacities)
  output = []
  impossible_packages = packages.select { |result| result.include?(false) }
  only_partners = packages.select { |package| package.count == 1 }
  possible_packages = packages - impossible_packages - only_partners
  over_capacity_partners = []

  capacities.each do |capacity|
  end

  only_partners.flatten.map{ |only_partner| only_partner.delete(:size) }
  output << only_partners.flatten.map(&:values).flatten
  output << impossible_packages.first
  output
end

main
