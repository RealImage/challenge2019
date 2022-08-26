require 'csv'
require './common'
class ProblemStatment1
  def problem_solution
    output_arr = []
    partners_hash = Common.partners_hash
    CSV.foreach(("input.csv")) do |row|
      partner_obj = partners_hash.select{ |p| p[:theater_id] == row[2]}
      input_slab = row[1].to_i
      data = [row[0], false, '', '']
      partner_obj&.each do |partner|
        start_slab, end_slab = partner[:slab].split('-')
        if (start_slab.to_i..end_slab.to_i).include?(input_slab)
          tot_cost = input_slab * partner[:cost].to_i
          data[1] = true
          data[2] = partner[:partner_id]
          data[3] = (tot_cost < partner[:min_cost].to_i) ? partner[:min_cost].to_i : tot_cost
          break;
        end
      end
      puts "#{data[0]}, #{data[1]}, #{data[2]}, #{data[3]}"
    end
  end
end

ProblemStatment1.new.problem_solution
