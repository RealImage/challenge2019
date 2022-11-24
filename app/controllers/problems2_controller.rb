class Problems2Controller < ApplicationController
    require 'csv'

    def index
        capacity_file = Dir.pwd + '/capacities.csv'
        my_output_file2 = Dir.pwd  + '/my_output2.csv'
        optimal_partner = []
        row1 = []
        capacity_csv = CSV.read(capacity_file)
        low_cost_delivery = low_cost_delivery()
        low_cost_delivery_with_limit = []
        (1..capacity_csv.count-1).each do |i|
          capacity_row = capacity_csv[i].map{|a| strip_all_whitespaces(a)}
          avl_partners = low_cost_delivery.select{|a| a[2] == capacity_row[0]} #pluck partners
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
            low_cost_delivery_with_limit << avl_partners_v2
          else
            low_cost_delivery_with_limit << avl_partners
          end
      
        end
        low_cost_delivery_with_limit << low_cost_delivery.select{|a| a[1] == 'false'}
        render json: low_cost_delivery_with_limit
        CSV.open('./output2.csv',"a+") do |row|
          low_cost_delivery_with_limit.flatten(1).uniq.each do |a|
            row << a[0..-3]
          end
        end
    end

    def strip_all_whitespaces string
      string.gsub!(/\s+/, '')
      string
    end
    
    def low_cost_delivery
      array = []
      partners_file = Dir.pwd  + '/partners.csv'
      input_file = Dir.pwd  + '/input.csv'
      optimal_partner = []
      @partners_csv = CSV.read(partners_file)
      @input_csv = CSV.read(input_file)
      @input_csv.each do |input_row|
        avl_partners = []
        (1 .. @partners_csv.count-1).each do |i|
          row = @partners_csv[i].map{|a| strip_all_whitespaces(a)}
          range = row[1].split('-').map(&:to_i)
          if (range[0] .. range[1]).include?(input_row[1].to_i) && input_row[2] == row[0]
            if input_row[1].to_i*row[3].to_i >= row[2].to_i
              avl_partners << row +[input_row[1].to_i*row[3].to_i ]
            else
              avl_partners << row +[ row[2].to_i ]
            end
          end
        end
        if avl_partners.count > 0
          optimal_partner = avl_partners.min_by(&:last)
          array << [input_row[0], 'true', optimal_partner[4], optimal_partner[5] ,input_row[2], input_row[1].to_i]
        else
          array << [input_row[0], 'false', '', '',input_row[2], input_row[1].to_i]
        end
      end
      array
    end
    
    
    def get_optimal input_row
      avl_partners = []
      (1 .. @partners_csv.count-1).each do |i|
        row = @partners_csv[i].map{|a| strip_all_whitespaces(a)}
        range = row[1].split('-').map(&:to_i)
        # binding.pry
        if (range[0] .. range[1]).include?(input_row[1].to_i) && input_row[2] == row[0]
          # p row +[input_row[1].to_i*row[3].to_i ]
          if input_row[1].to_i*row[3].to_i >= row[2].to_i
            avl_partners << row +[input_row[1].to_i*row[3].to_i ]
          else
            avl_partners << row +[ row[2].to_i ]
          end
        end
      end
      avl_partners
    end
    
end
