require 'csv'

def squish string
  string.gsub!(/\s+/, '')
  string
end

def lowest_cost_deleviry
  array = []
  partners_file = Dir.pwd  + '/partners.csv'
  input_file = Dir.pwd  + '/input.csv'
  optimal_partner = []
  @partners_csv = CSV.read(partners_file)
  @input_csv = CSV.read(input_file)
  @input_csv.each do |input_row|
    avalible_partners = []
    (1 .. @partners_csv.count-1).each do |i|
      row = @partners_csv[i].map{|a| squish(a)}
      range = row[1].split('-').map(&:to_i)
      if (range[0] .. range[1]).include?(input_row[1].to_i) && input_row[2] == row[0]
        if input_row[1].to_i*row[3].to_i >= row[2].to_i
          avalible_partners << row +[input_row[1].to_i*row[3].to_i ]
        else
          avalible_partners << row +[ row[2].to_i ]
        end
      end
    end
    if avalible_partners.count > 0
      optimal_partner = avalible_partners.min_by(&:last)
      array << [input_row[0], 'true', optimal_partner[4], optimal_partner[5] ,input_row[2], input_row[1].to_i]
    else
      array << [input_row[0], 'false', '', '',input_row[2], input_row[1].to_i]
    end
  end
  array
end


def get_optimal input_row
  avalible_partners = []
  (1 .. @partners_csv.count-1).each do |i|
    row = @partners_csv[i].map{|a| squish(a)}
    range = row[1].split('-').map(&:to_i)
    # binding.pry
    if (range[0] .. range[1]).include?(input_row[1].to_i) && input_row[2] == row[0]
      # p row +[input_row[1].to_i*row[3].to_i ]
      if input_row[1].to_i*row[3].to_i >= row[2].to_i
        avalible_partners << row +[input_row[1].to_i*row[3].to_i ]
      else
        avalible_partners << row +[ row[2].to_i ]
      end
    end
  end
  avalible_partners
end


def problem_1
  my_output_file1 = Dir.pwd  + '/my_output1.csv'
  system("rm #{my_output_file1}")
  lowest_cost_deleviry = lowest_cost_deleviry()
  CSV.open(my_output_file1,"a+") do |row|
    lowest_cost_deleviry.each do |a|
      row << a[0..-3]
    end
  end
end

def problem_2
  capacity_file = Dir.pwd + '/capacities.csv'
  my_output_file2 = Dir.pwd  + '/my_output2.csv'
  optimal_partner = []
  row1 = []
  capacity_csv = CSV.read(capacity_file)
  lowest_cost_deleviry = lowest_cost_deleviry()
  lowest_cost_deleviry_with_limit = []
  (1..capacity_csv.count-1).each do |i|
    capacity_row = capacity_csv[i].map{|a| squish(a)}
    avl_partners = lowest_cost_deleviry.select{|a| a[2] == capacity_row[0]} #pluck partners
    sum = 0
    avl_partners.map{ |a| sum+=a.last }
    if sum > capacity_row[1].to_i
      total =0
      total = avl_partners.map{ |a| total+=a[3] }
      array = []
      avl_partners.each do |avl_partner|
        input = [avl_partner[0], avl_partner[5], avl_partner[4] ]
        s = get_optimal(input)
        s1 = s.select{ |a| a[4] != avl_partner[2]}.min_by(&:last)
        array << s1.last.to_i - avl_partner[3].to_i
      end
      indx = array.index(array.max)
      avl_partners_v2 =avl_partners
      input = [avl_partners_v2[indx][0], avl_partners_v2[indx][5], avl_partners_v2[indx][4]]
      s = get_optimal(input)
      s1 = s.select{ |a| a[4] != avl_partners_v2[indx][2]}.min_by(&:last)
      avl_partners_v2[indx][2] = s1[4]
      avl_partners_v2[indx][3] = s1[5]
      lowest_cost_deleviry_with_limit << avl_partners_v2
    else
      lowest_cost_deleviry_with_limit << avl_partners
    end
    
  end
  lowest_cost_deleviry_with_limit << lowest_cost_deleviry.select{|a| a[1] == 'false'}
  system("rm #{my_output_file2}")
  CSV.open(my_output_file2,"a+") do |row|
    lowest_cost_deleviry_with_limit.flatten(1).uniq.each do |a|
      row << a[0..-3]
    end
  end
end

problem_1()
problem_2()
