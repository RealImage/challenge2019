require 'csv'
require './common'
class ProblemStatment1
  def problem_solution
    partners_hash = Common.theater_with_partners_hash
    partner_data_hash = {}
    CSV.foreach(("input.csv")) do |row|
      partner_obj = partners_hash[row[2]]
      input_slab = row[1].to_i
      input_slab_value = (0..row[1].to_i).to_a
      data = [row[0], false, '', '']
      partner_data_hash[row[0]] ||= {}
      partner_obj&.each do |parter_id, partners|
        subtracted_gb = input_slab
        partners_count = partners.count
        is_partner_has_no_starting_cost_val, starting_slab, has_more_available_partner = Common.get_partner_missed_slab_details(partners, input_slab)
        partners&.each_with_index do |partner, ind|
          start_slab, end_slab, gb_count, is_partner_has_no_starting_cost = Common.get_calculation_details(partner, is_partner_has_no_starting_cost_val)
          sub_gb = subtracted_gb - gb_count.to_i
          is_last_partner_has_gb = partners_count == (ind + 1) && (sub_gb != 0 && sub_gb.positive?)
          if is_partner_has_no_starting_cost && (ind + 1) == 1
            subtracted_gb = subtracted_gb - starting_slab.sort.first
          end
          if is_last_partner_has_gb
            partner_data_hash[row[0]][parter_id] = 0
          elsif end_slab.to_i >= input_slab && subtracted_gb != 0
            partner_cost = (subtracted_gb * partner[:cost].to_i)
            old_parner_cost = partner_data_hash[row[0]][parter_id].to_i
            subtracted_gb = has_more_available_partner ? input_slab : 0
            final_cost = partner_cost < partner[:min_cost].to_i ? partner[:min_cost].to_i : partner_cost
            new_cost = has_more_available_partner && old_parner_cost != 0 && old_parner_cost < final_cost ? old_parner_cost : final_cost
            final_cost = has_more_available_partner ? new_cost : (partner_data_hash[row[0]][parter_id].to_i + final_cost)
            partner_data_hash[row[0]][parter_id] = final_cost
          elsif (end_slab.to_i <= input_slab) && subtracted_gb != 0
            subtracted_gb = subtracted_gb - gb_count.to_i
            tot_cost = end_slab.to_i * partner[:cost].to_i
            tot_cost = tot_cost < partner[:min_cost].to_i ? partner[:min_cost].to_i : tot_cost
            partner_data_hash[row[0]][parter_id] = partner_data_hash[row[0]][parter_id].to_i + tot_cost.to_i
          end
        end
      end
    end
    Common.print_delivery_partner_details(partner_data_hash)
  end
end

ProblemStatment1.new.problem_solution
