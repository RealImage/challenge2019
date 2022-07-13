require "./csv/import_csv"

class DeliveryPartner
  extend ImportCsv
  
  attr_accessor :theater_id, :slab_size, :minimum_cost, :cost_per_unit, :partner_id

  FILE_NAME = "partners.csv"

  def initialize(arr)
    @theater_id = arr[0].strip
    l, g = arr[1].strip.split("-").map(&:to_i)
    @slab_size = Range.new(l, g)
    @minimum_cost = arr[2].strip.to_i
    @cost_per_unit = arr[3].strip.to_i
    @partner_id = arr[4].strip
  end

  def self.all
    import(FILE_NAME, { headers: true })
  end
end