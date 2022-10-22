

class Capacity
  attr_accessor :id, :capacity, :partner_id

  def initialize(id:, capacity:, partner_id:)
    @id = id
    @capacity = capacity
    @partner_id = partner_id
  end

  def self.capacity_data
    capacities = []
    count = 0
    filepath = File.join(Dir.pwd, 'capacities.csv')
    CSV.foreach(filepath, headers: true) do |row|
      cap = Capacity.new(
        id: (count += 1),
        capacity: row["Capacity (in GB)"]&.strip.to_i,
        partner_id: row["Partner ID"]&.strip
      )
      capacities << cap
    end
    capacities
  end
end