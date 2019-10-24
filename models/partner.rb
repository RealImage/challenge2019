class Partner

  attr_accessor :min_size, :max_size, :min_cost, :cost_per_gb, :partner_id, :theatre_id, :max_capacity, :capacity_assigned

  def initialize(min_size, max_size, min_cost, cost_per_gb, partner_id, theatre_id)
    @min_size = min_size
    @max_size = max_size
    @min_cost = min_cost
    @cost_per_gb = cost_per_gb
    @partner_id = partner_id
    @theatre_id = theatre_id
    @max_capacity = 0
    @capacity_assigned = 0
  end

end