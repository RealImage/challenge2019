require 'csv'

def main
  partners = read_constant_data('partners')
  results = input_content.map do |deliver|
    get_partner(deliver, partners)
  end

  File.write('output1.csv', results.map(&:to_csv).join)
end

def input_content
  CSV.open('./input.csv').each.to_a.compact.map do |content|
    content.map(&:strip)
  end
end

def read_constant_data(file_name)
  data = []
  CSV.foreach("./#{file_name}.csv", headers: true) do |row|
    data << {
      theatre: row['Theatre'].strip,
      size_lab_min: row['Size Slab (in GB)'].strip.split('-')[0].to_i,
      size_lab_max: row['Size Slab (in GB)'].strip.split('-')[1].to_i,
      minimun_cost: row['Minimum cost'].to_i,
      cost_per_gb: row['Cost Per GB'].to_i,
      partner_id: row['Partner ID'].strip
    }
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
        partner_id: theatre[:partner_id],
        cost: const_temp >= theatre[:minimun_cost] ? const_temp : theatre[:minimun_cost]
      }
    end

    delivery_costs.min { |a, b| a[:cost] <=> b[:cost] }.values
  else
    [deliver[0], false, '', '']
  end
end

main
