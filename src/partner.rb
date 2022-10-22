
class Partner
  attr_accessor :id, :theatre, :size_slab, :min_cost, :cost_per_gb, :partner_id

  def initialize(id:, theatre:, size_slab:, min_cost:, cost_per_gb:, partner_id:)
    @id = id
    @theatre = theatre
    @size_slab = size_slab
    @min_cost = min_cost
    @cost_per_gb = cost_per_gb
    @partner_id = partner_id
  end

  class << self
    def partners_data
      count = 0
      partners  = []
      filepath = File.join(Dir.pwd, 'partners.csv')
      CSV.foreach(filepath, headers: true) do |row|
        partner = Partner.new(
          id: (count += 1),
          theatre: row["Theatre"].strip,
          size_slab: row["Size Slab (in GB)"].strip,
          min_cost: row["Minimum cost"].strip.to_i,
          cost_per_gb: row["Cost Per GB"].strip.to_i,
          partner_id: row["Partner ID"].strip,
        )
        partners << partner
      end
      partners
    end
  end
end