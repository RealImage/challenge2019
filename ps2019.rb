## IMPORTANT
## FOR CONSISTENCY ADDED HEADER TO DELIVERY FILE ##
# FILE STRUCTURE
# INPUT FILE => "Delivery ID","Size","Theatre"
# Partners => Theatre,Size Slab (in GB),Minimum cost,Cost Per GB,Partner ID
# Capacity => "Partner ID","Capacity (in GB)"

## IMPORTANT
class FindPartner
    @@theatre = {}
    @@final_delivery = {}

    def initialize(partner, input, capacity)
      @partners  = get_csv(partner,true) if !partner.empty?
      @input  = get_csv(input,true)  if !input.empty?
      get_csv(capacity,true) if !capacity.empty?
    end

    def processing_logic()
      get_minimum_delivery_price()
    end

  private
    @partners = []
    @input = []
    @@capacity = {}
  # Input file read and prepare struct format
  def get_csv(filename, header)
      require 'csv'
      file_datum = []
      index  = 0
      # table = CSV.parse(File.read(filename), headers: header,  :header_converters => :symbol)
      CSV.foreach(File.open(filename), headers: header, :converters => :numeric, :header_converters => :symbol){ |row|
      next if row.empty?
      case
        when filename.downcase().include?('partners')
          index = index + 1
          row[:pid]= index
          row[:min_slab] = row[:size_slab_in_gb].split()[0].split("-")[0].to_i
          row[:max_slab] = row[:size_slab_in_gb].split()[0].split("-")[1].to_i
          row[:theatre] = row[:theatre].split[0]
          @@theatre.store(row[:theatre],row[:theatre]) if row[:theatre] != nil and !@@theatre.has_key?(row[:theatre])
        when filename.downcase().include?('capacities')
          @@capacity.store(row[:partner_id].split[0],row[:capacity_in_gb])
        else
          #nothing
        end
        file_datum.append(row)
      }

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
        possible.append(partner[:pid], partner[:partner_id], theatre, delivery[:size], max_cost, partner[:cost_per_gb])
        all.append(possible)
        possible=[]
      end
      }
      # if !all.empty?
      #   # list_price.store(delivery[:delivery_id],all)
      # end
      all if !all.empty?
  end

  def get_minimum_delivery_price()
    input = @input.to_a
    p "Preferred Delivery"
    @input.each do |delivery|
      found_capacity = false
      # "call" logic
      found = get_vendor(delivery,delivery[:theatre])
      if !found.nil?
        print delivery[:delivery_id], "\t", delivery[:theatre], "\t", delivery[:size], "\t", 'True', "\t", found[0][0], "\t", found[0][1]
      else
        print delivery[:delivery_id], "\t", delivery[:theatre], "\t", delivery[:size], "\t", 'False', "\t", '', "\t", ''
      end
      puts
        @@final_delivery.store(delivery[:delivery_id],found)
    end

    #  process per Theatre  - 2. check total needed vs total capacity 3. sort by max size/max priority against least cost 4. mark
    found=[]
    available_delivery=[]
    @@theatre.each do |key,val|
      #"total Qty needed for the theatre
      # "sorted by decreasing size" - This can be altered to check PRIORITY IF NEEDED
      deliveries =  @input.filter{ |t| t[:theatre] == val }.sort_by{ |r| r[:size]}.reverse!
      # for theatre - D1  + D3 => 660
      total = deliveries.sum{ |r|  r[:size] }
      # "need to check against each delivery"
      deliveries.each do |delivery|
          preferred_deliveries = @@final_delivery[delivery[:delivery_id]]
          available=0
          found=[]
          needed = delivery[:size]
          if preferred_deliveries.nil?
            found.append(delivery[:delivery_id],false)
            available_delivery.append(found)
          else
           preferred_deliveries.each do |pref|
              available += @@capacity[pref[1]]
                if available > delivery[:size] and needed > 0
                  @@capacity[pref[1]] = available - delivery[:size]
                  available = available - delivery[:size]
                  needed = needed - delivery[:size]
                  found.append(delivery[:delivery_id],true, pref[1], pref[4])
                  available_delivery.append(found)
                  break
                end
            end
        end
       end
      # Total needed
      available = 0
      # needed = per_theatre.sum{ |r|  r[:size] }
    end
    puts "Possible Delivery"
    available_delivery.each{ |d| print d ,"\n" }
  end
end

require 'debug'

# "start of processing"
ab = FindPartner.new("partners.csv","input.csv","capacities.csv")
# "pS1" - minimum cost
ab.processing_logic()




