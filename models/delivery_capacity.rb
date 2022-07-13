require "./csv/import_csv"

class DeliveryCapacity
  extend ImportCsv
  
  attr_accessor :partner_id, :capacity

  FILE_NAME = "capacities.csv"
  
  def initialize(arr)
    @partner_id = arr[0].strip
    @capacity = arr[1].strip.to_i
  end

  def self.all
    import(FILE_NAME, { headers: true })
  end
end