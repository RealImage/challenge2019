require "./models/delivery"
require "./models/delivery_capacity"
require "./models/delivery_partner"
require "./csv/export_csv"

PROBLEM1_OUTPUT_FILE = "output1.csv"
PROBLEM2_OUTPUT_FILE = "output2.csv"

def delivery_infos(delivery_partners, deliveries)
  deliveries.map do |delivery|
    delivery.estimation(delivery_partners)
  end
end

def minimum_cost_delivery_infos(delivery_partners, deliveries, delivery_capacities)  
  deliveries_desc_size = deliveries.sort_by { |d| -d.size }
 
  delivery_infos = deliveries_desc_size.map do |delivery|
    delivery.estimation_upon_capacity(delivery_partners, delivery_capacities)
  end

  delivery_infos.reverse
end

# problem statement-1
data = delivery_infos(DeliveryPartner.all, Delivery.all)
ExportCsv.export(PROBLEM1_OUTPUT_FILE, data)

# problem statement-2
data = minimum_cost_delivery_infos(DeliveryPartner.all, Delivery.all, DeliveryCapacity.all)
ExportCsv.export(PROBLEM2_OUTPUT_FILE, data)
