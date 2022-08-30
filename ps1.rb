## IMPORTANT
## FOR CONSISTENCY ADDED HEADER TO DELIVERY FILE ##
# FILE STRUCTURE
# INPUT FILE => "Delivery ID","Size","Theatre"
# Partners => Theatre,Size Slab (in GB),Minimum cost,Cost Per GB,Partner ID
# Capacity => "Partner ID","Capacity (in GB)"

## IMPORTANT
class FindPartner

    @@final_delivery = []
    def initialize(partner, input, capacity)
      @partners  = get_csv(partner,true) if !partner.empty?
      @input  = get_csv(input,true)  if !input.empty?
      @@capacity  = get_csv(capacity,true) if !capacity.empty?
    end

    def processing_logic()
      get_minimum_delivery_price()
    end

  private
    @partners = []
    @input = []
    @@capacity = []
  # Input file read and prepare struct format
  def get_csv(filename, header)
      require 'csv'
      file_datum = []
      # table = CSV.parse(File.read(filename), headers: header,  :header_converters => :symbol)
      CSV.foreach(File.open(filename), headers: header, :converters => :numeric, :header_converters => :symbol){ |row|
        case
        when filename.downcase().include?('partners')
          row[:min_slab] = row[:size_slab_in_gb].split()[0].split("-")[0].to_i
          row[:max_slab] = row[:size_slab_in_gb].split()[0].split("-")[1].to_i
          row[:theatre] = row[:theatre].split[0]
        when filename.downcase().include?('capacities')
          row[:partner_id] = row[:partner_id].split[0]
        else
          #nothing
        end
        file_datum.append(row) }
        file_datum
  end

  # "get possible vendor quotes"
  def get_vendor(delivery, theatre)
    list_price = {}
    all = []
    possible = []
    row = Array.new
    partners = @partners.filter{ |partner| partner[:theatre] == theatre && partner[:min_slab] <= delivery[:size] && partner[:max_slab] >= delivery[:size]}

    partners.each {  |partner|
      # partner[:theatre] == theatre && partner[:min_slab] <= delivery[:size] && partner[:max_slab] >= delivery[:size]
      if ( partner[:theatre] == theatre && partner[:min_slab] <= delivery[:size] && partner[:max_slab] >= delivery[:size] )
        max_cost =  delivery[:size] * partner[:cost_per_gb] <= partner[:minimum_cost] ? partner[:minimum_cost] : delivery[:size] * partner[:cost_per_gb]
        possible.append(partner[:partner_id],max_cost, partner[:cost_per_gb])
        all.append(possible)
        possible=[]
      end
      }
      if !all.empty?
        list_price.store(delivery[:delivery_id],all)
      end
      list_price
  end

  def get_minimum_delivery_price()
    input = @input.to_a
    @input.each do |delivery|
      found_capacity = false
      # "call" logic
      @found = get_vendor(delivery,delivery[:theatre])
      found = @found.to_a
      p "Preferred"
      if !found.empty?
        print delivery[:delivery_id], "\t", delivery[:theatre], "\t", delivery[:size], "\t", 'True', "\t", @found[delivery[:delivery_id]][0][0], "\t", @found[delivery[:delivery_id]][0][1]
      else
        print delivery[:delivery_id], "\t", delivery[:theatre], "\t", delivery[:size], "\t", 'False', "\t", '', "\t", ''
      end
      puts
        @@final_delivery.append(@found)
    end

    binding.break
     @@capacity.each do |cap|

    end
  end



end

require 'debug'

# "start of processing"
ab = FindPartner.new("partners.csv","input.csv","capacities.csv")
# "pS1" - minimum cost
ab.processing_logic()

# ab.find_optimal_cost()




