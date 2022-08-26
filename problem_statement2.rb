require 'csv'
require './common'
class ProblemStatment2
  def ProblemStatment2.verify_partners(partners, theatre_id, slab_size)
    partner_arr = []
    partners.sort_by { |p| p[:cost] }.each do |partner|
      start_slab, end_slab = partner[:slab].split('-')
      partner_arr << partner if partner[:theater_id].to_s == theatre_id.to_s && (start_slab.to_i..end_slab.to_i).include?(slab_size)
    end
    partner_arr
  end

  def problem_solution
    output_arr = []
    partners_hash = Common.partners_hash
    capacities_hash = Common.capacities_hash
    Common.input_arr.reverse.each do |row|
      slab_size = row[1].to_i
      verified_partners = ProblemStatment2.verify_partners(partners_hash, row[2], slab_size)
      if verified_partners.any?
        low_cost_partner = nil
        verified_partners.each do |partner|
          next if (capacities_hash[partner[:partner_id]].to_i <= slab_size)
          low_cost_partner = partner
          break;
        end
        if low_cost_partner.any?
          capacities_hash[low_cost_partner[:partner_id]] -= slab_size
          cal_cost = low_cost_partner[:cost].to_i * slab_size
          cost = cal_cost < low_cost_partner[:min_cost].to_i ? low_cost_partner[:min_cost].to_i : cal_cost
          output_arr << [row[0], 'true', low_cost_partner[:partner_id], cost]
        else
          output_arr << [row[0], 'false', '', '']
        end
      else
        output_arr << [row[0], 'false', '', '']
      end
    end
    output_arr.reverse.each do |data|
      puts "#{data[0]}, #{data[1]}, #{data[2]}, #{data[3]}"
    end
  end
end
ProblemStatment2.new.problem_solution
