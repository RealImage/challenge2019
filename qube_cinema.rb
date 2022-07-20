require 'csv'

class QubeCinema

  # NOTE: Solution to problem statement 1 from the README.md
  def find_partner_for_delivery
    results = []

    CSV.foreach('./input.csv') do |delivery|
      partner_payload = [nil, false, '', Float::INFINITY]
      delivery_status = false

      CSV.foreach('./partners.csv', headers: true, header_converters: :symbol) do |partner|
        next unless partner[:theatre].strip.eql?(delivery[2].strip)

        min_slab, max_slab = partner[:size_slab_in_gb].split('-').map(&:to_i)
        input_gb = delivery[1].to_i

        if (min_slab..max_slab).cover?(input_gb)
          delivery_status = true

          total_cost = partner[:cost_per_gb].to_i * input_gb
          cost = total_cost > partner[:minimum_cost].to_i ? total_cost : partner[:minimum_cost].to_i

          partner_payload = [delivery[0], delivery_status, partner[:partner_id].strip, cost] if cost < partner_payload.last
        end
      end

      results << (delivery_status ? partner_payload : [delivery[0], false, '', ''])
    end

    generate_output('./output.csv', results)
  end

  private

  def generate_output(path, rows)
    CSV.open(path, 'w') do |csv|
      rows.each do |row|
        csv << row
      end
    end
  end
end

QubeCinema.new.find_partner_for_delivery
