
require 'pry'

def partner_capacity
    {
      'P1': 350,
      'P2': 500,
      'P3': 1500
    }
end


def get_partners_list_by_theater(theater)
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

def get_todo_deliveries

  # [
  #   {delivery_id: 'D1', size_slab_in_gb: 100, theater: 'T1'},
  #   {delivery_id: 'D2', size_slab_in_gb: 240, theater: 'T1'},
  #   {delivery_id: 'D3', size_slab_in_gb: 260, theater: 'T1'},
  # ]
  
  [
    {delivery_id: 'D1', size_slab_in_gb: 150, theater: 'T1'},
    {delivery_id: 'D2', size_slab_in_gb: 325, theater: 'T2'},
    {delivery_id: 'D3', size_slab_in_gb: 510, theater: 'T1'},
    {delivery_id: 'D4', size_slab_in_gb: 700, theater: 'T2'}
  ]
end

def partner_cost_evaluation(partner, size_slab_in_gb)
  consumption_cost = partner[:cost_per_gb] * size_slab_in_gb
  [consumption_cost, partner[:minimum_cost]].max
end

def partner_can_accomadate_delivery?(partner, delivery, capacities)
  capacities[partner[:partner].to_sym] > delivery[:size_slab_in_gb]
end

def find_best_partner(partners, delivery, capacities)
  best_partner_cost = best_partner = nil
  partners.each do |partner|
    next unless partner_can_accomadate_delivery?(partner, delivery, capacities)
    partner_cost = partner_cost_evaluation(partner, delivery[:size_slab_in_gb])
    if best_partner.nil? 
      best_partner = partner
      best_partner_cost = partner_cost
    elsif partner_cost < best_partner_cost
      best_partner = partner
      best_partner_cost = partner_cost
    end
  end
  best_partner[:best_partner_cost] = best_partner_cost
  best_partner
end


# Proecss the deliveries in descending order, so that we can make sure maximum deliveries are taken care
def find_delivery_partners
  todo_deliveries = get_todo_deliveries.sort {|x, y | y[:size_slab_in_gb] <=> x[:size_slab_in_gb]}
  capacities = partner_capacity
  delivery_partners = []
  todo_deliveries.each do |delivery|
    theater_partners = get_partners_list_by_theater(delivery[:theater])
    partners = theater_partners.select {|x| x[:size_slab_in_gb].include?(delivery[:size_slab_in_gb])}
     if partners.empty?
      puts "#{delivery[:delivery_id]}, false, '', ''"
      next
    end
    best_partner = find_best_partner(partners, delivery, capacities)
    puts "#{delivery[:delivery_id]}, true, #{best_partner[:partner]}, #{best_partner[:best_partner_cost]}"
    capacities[best_partner[:partner].to_sym] -= delivery[:size_slab_in_gb]
  end
  # puts "partner capacities left out: "
  # puts capacities
end

find_delivery_partners


# D4, false, '', ''
# D3, true, P3, 15300
# D2, true, P1, 3250
# D1, true, P2, 3000
