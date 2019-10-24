class Parser

  def initialize(csv_data)
    @csv_data = csv_data
    @partners = []
  end

  def parse_partners
    @csv_data.each do|row|
      @partners << build_partner_data(row)
    end
    @partners
  end

  def parse_partner_capacities(capacities)
    capacities.each do|row|
      update_capacities_for_partners(row)
    end
    @partners
  end

  def clean_input_data(input_data)
    input_data.map { |row| row.map { |value| value.strip } }
  end

  private

  def build_partner_data(row)
    min_size, max_size = row["Size Slab (in GB)"].strip.split("-")
    Partner.new(
      min_size = min_size.to_i,
      max_size = max_size.to_i,
      min_cost = row["Minimum cost"].strip.to_i,
      cost_per_gb = row["Cost Per GB"].strip.to_i,
      partner_id = row["Partner ID"].strip,
      theater = row["Theatre"].strip
    )
  end

  def update_capacities_for_partners(row)
    @partners.each do|partner|
      if partner.partner_id == row["Partner ID"].strip
        partner.max_capacity = row["Capacity (in GB)"].strip.to_i
      end
    end
  end

end