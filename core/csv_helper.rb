require 'csv'

module CSVHelper
  def parse_partners(file)
    csv = CSV.parse(File.read(file), headers: true)

    csv.map do |row|
      {
        theatre: row['Theatre'].strip,
        size_slab_lower: row['Size Slab (in GB)'].strip.split('-').first.to_i,
        size_slab_higher: row['Size Slab (in GB)'].strip.split('-').last.to_i,
        min_cost: row['Minimum cost'].to_i,
        cost_per_gb: row['Cost Per GB'].to_i,
        partner_id: row['Partner ID'].strip
      }
    end
  end

  def parse_capacities(file)
    csv = CSV.parse(File.read(file), headers: true)

    csv.map do |row|
      [
        row['Partner ID'].strip,
        row['Capacity (in GB)'].to_i
      ]
    end.to_h
  end

  def parse_input(file)
    csv = CSV.parse(File.read(file))

    csv.map do |row|
      {
        distribution_id: row[0].strip,
        required_size: row[1].to_i,
        theatre: row[2].strip
      }
    end
  end

  def convert_to_csv(content)
    directory_name = 'name'
    Dir.mkdir(directory_name) unless File.exist?(directory_name)

    CSV.open('tmp/output.csv', 'w') do |csv|
      content.each do |line|
        csv << line
      end
    end
  end
end
