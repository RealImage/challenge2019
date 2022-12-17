class Partner
  attr_accessor :id, :theatre_id , :size_slab, :min_cost, :cost_per_gb, :partner_id

  def initialize(id:, theatre_id: , size_slab:, min_cost:, cost_per_gb:, partner_id:)
    @id = id
    @theatre_id = theatre_id
    @size_slab = size_slab
    @min_cost = min_cost
    @cost_per_gb = cost_per_gb
    @partner_id = partner_id
  end

  def self.get_partners
    partners_arr = []
    count = 0
    CSV.foreach(File.join(Dir.pwd, 'partners.csv'), headers: true) do |row|
      partner = Partner.new(
        id: (count += 1),
        theatre_id: row["Theatre"].strip,
        size_slab: row["Size Slab (in GB)"].strip,
        min_cost: row["Minimum cost"].strip.to_i,
        cost_per_gb: row["Cost Per GB"].strip.to_i,
        partner_id: row["Partner ID"].strip,
        )
      partners_arr << partner
    end
    partners_arr
  end
end