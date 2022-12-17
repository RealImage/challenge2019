class Capacity
  attr_accessor :id, :capacity, :partner_id

  def initialize(id:, capacity:, partner_id:)
    @id = id
    @capacity = capacity
    @partner_id = partner_id
  end

  def self.get_capacities
    capacities_arr = []
    row_count = 0
    CSV.foreach(File.join(Dir.pwd, 'capacities.csv'), headers: true) do |row|
      capacity = Capacity.new(
        id: (row_count += 1),
        capacity: row['Capacity (in GB)'].strip.to_i,
        partner_id: row['Partner ID'].strip
        )
      capacities_arr << capacity
    end
    capacities_arr
  end
end
