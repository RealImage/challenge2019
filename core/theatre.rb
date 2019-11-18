class Theatre
  def initialize(name)
    @name = name
  end

  # protected

  def possible_partners_list(partners, size)
    partners.select do |p|
      p[:theatre] == @name && p[:size_slab_lower] <= size && p[:size_slab_higher] >= size
    end
  end

  def cost_of_delivery(partners, size)
    ps = []

    partners.each do |p|
      exp_cost = size * p[:cost_per_gb]

      tmp = {
        theatre: p[:theatre],
        partner_id: p[:partner_id],
        size: size,
        cost_of_delivery: exp_cost > p[:min_cost] ? exp_cost : p[:min_cost]
      }

      ps << tmp
    end

    ps
  end

  def minimum_cost_of_delivery(partners)
    partners.min_by do |p|
      raise 'cost_of_delivery is required' if p[:cost_of_delivery].nil?

      p[:cost_of_delivery]
    end
  end
end
