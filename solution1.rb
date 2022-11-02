require 'csv'

CSV::Converters[:range] = ->(s) { 
  begin 
    Range.new(*s.split('-').map(&:to_i))
  rescue ArgumentError
    s
  end
}

CSV::Converters[:strip] = ->(s) { s&.strip }

PARTNERS = CSV.parse(File.read("partners.csv"), converters: [:integer, :range, :strip], headers: true)
THEATRE = "Theatre"
SIZE_SLAB =  "Size Slab (in GB)"
MIN_COST = "Minimum cost"
COST_PER_GB = "Cost Per GB"
PARTNER_ID = "Partner ID"


DELIVERIES = CSV.parse(File.read("input.csv"), converters: [:numeric, :strip])

result = []

DELIVERIES.each do |delivery|
	result << [delivery[0], false, '', '' ]

	PARTNERS.each do |partner|
		next if partner[THEATRE] != delivery[2]
		next unless partner[SIZE_SLAB].include?(delivery[1])

    price = [delivery[1] * partner[COST_PER_GB], partner[MIN_COST]].max
	  result[-1] = [delivery[0], true, partner[PARTNER_ID], price] if result[-1][-1].is_a?(String) || price < result[-1][-1]
	end
end

CSV.open("output1.csv", "w") do |csv|
  result.each {|r| csv << r}
end
