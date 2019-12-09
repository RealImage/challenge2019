
def get_partners_list(theater)
  {
    'T1': [
      { partner: 'P1', size_slab_in_gb: 0..100, minimum_cost: 1500, cost_per_gb: 20 },
      { partner: 'P1', size_slab_in_gb: 101..200, minimum_cost: 2000, cost_per_gb: 13 },
      { partner: 'P1', size_slab_in_gb: 201..300, minimum_cost: 2500, cost_per_gb: 12 },
      { partner: 'P1', size_slab_in_gb: 301..400, minimum_cost: 3000, cost_per_gb: 10 },
      
      { partner: 'P2', size_slab_in_gb: 0..200, minimum_cost: 1000, cost_per_gb: 20 },
      { partner: 'P2', size_slab_in_gb: 201..400, minimum_cost: 2500, cost_per_gb: 15 },

      { partner: 'P3', size_slab_in_gb: 100..200, minimum_cost: 800, cost_per_gb: 25 },
      { partner: 'P3', size_slab_in_gb: 201..600, minimum_cost: 1200, cost_per_gb: 30 }
    ],
    'T2': [
      { partner: 'P1', size_slab_in_gb: 0..100, minimum_cost: 1500, cost_per_gb: 20 },
      { partner: 'P1', size_slab_in_gb: 101..200, minimum_cost: 2000, cost_per_gb: 13 },
      { partner: 'P1', size_slab_in_gb: 201..300, minimum_cost: 2500, cost_per_gb: 12 },
      { partner: 'P1', size_slab_in_gb: 301..400, minimum_cost: 3000, cost_per_gb: 10 },

      { partner: 'P2', size_slab_in_gb: 0..200, minimum_cost: 2500, cost_per_gb: 20 },
      { partner: 'P2', size_slab_in_gb: 201..400, minimum_cost: 3500, cost_per_gb: 10 },

      { partner: 'P3', size_slab_in_gb: 100..200, minimum_cost: 900, cost_per_gb: 15 },
      { partner: 'P3', size_slab_in_gb: 201..400, minimum_cost: 1000, cost_per_gb: 12 }

    ]
  }[theater.to_sym]
end


def partner_cost_evaluation(partner, size_slab_in_gb)
  consumption_cost = partner[:cost_per_gb] * size_slab_in_gb
  [consumption_cost, partner[:minimum_cost]].max
end


def evaluate_best_partner(partners_available, delivery)
  best_partner = partners_available.first
  best_partner_cost = partner_cost_evaluation(best_partner, delivery[:size_slab_in_gb])
  for i in 1..(partners_available.length - 1)
    partner_cost = partner_cost_evaluation(partners_available[i], delivery[:size_slab_in_gb])
    if partner_cost < best_partner_cost
      best_partner = partners_available[i]
      best_partner_cost = partner_cost
    end
  end
  puts "#{delivery[:delivery_id]}, true, #{best_partner[:partner]}, #{best_partner_cost}"
end

def find_best_partner(theater, delivery)
  return "Please input theater id" if theater.nil?
  partners = get_partners_list(theater)
  return "Theatre not found" if partners.nil?
  return "There are no partners for given theater" if partners.empty?
  partners_available = partners.select {|x| x[:size_slab_in_gb].include?(delivery[:size_slab_in_gb])}
  if partners_available.empty?
    puts "#{delivery[:delivery_id]}, false, '', ''"
    return
  end
  if partners_available.size == 1
    partner = partners_available.first
    partner_cost = partner_cost_evaluation(partner, delivery[:size_slab_in_gb])
    puts "#{delivery[:delivery_id]}, true, #{partner[:partner]}, #{partner_cost}"
    return
  end
  evaluate_best_partner(partners_available, delivery)
end

deliveries = [
  {delivery_id: 'D1', size_slab_in_gb: 50, theater: 'T1'},
  {delivery_id: 'D2', size_slab_in_gb: 325, theater: 'T2'},
  {delivery_id: 'D3', size_slab_in_gb: 510, theater: 'T1'},
  {delivery_id: 'D4', size_slab_in_gb: 700, theater: 'T2'}
]


deliveries.each do |delivery|
  find_best_partner(delivery[:theater], delivery)
end


# D1, true, P2, 1000
# D2, true, P1, 3250
# D3, true, P3, 15300
# D4, false, '', ''
