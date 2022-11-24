class Problem2
  require 'csv'

  @hash = {}
  @capacity_hash = {}
  @input_hash = {}
  @output_hash = {}
  @total_capacity = 0
  row_count = 0

  CSV.new(open('partners.csv'),  liberal_parsing: true, headers: :first_row).each do |row|
    @hash[row['Theatre'].rstrip.lstrip] = [] if @hash[row['Theatre'].rstrip.lstrip].nil?
    @hash[row['Theatre'].rstrip.lstrip] << { range: row['Size Slab (in GB)'].rstrip.lstrip, min_price: row['Minimum cost'].rstrip.lstrip,
                                             cost_per_gb: row['Cost Per GB'].rstrip.lstrip, partner_id: row['Partner ID'].rstrip.lstrip }
  end

  CSV.new(open('capacities.csv'),  liberal_parsing: true, headers: :first_row).each do |row|
    @capacity_hash[row['Partner ID'].rstrip.lstrip] = row['Capacity (in GB)'].rstrip.lstrip
    @total_capacity += row['Capacity (in GB)'].rstrip.lstrip.to_i
  end

  CSV.new(open('input.csv'),  liberal_parsing: true).each do |row|
    @total_capacity -= row[1].to_i
    if @total_capacity >= 0
      @input_hash[row.first] = {size: row[1].to_i, theatre: row.last}
    end
    row_count += 1
  end

  @input_arr = @input_hash.sort_by{ |_, v| v[:size] }

  @input_arr.reverse!.each do |input|
    min_price = nil
    partner = ''
    @hash[input.last[:theatre].lstrip.rstrip].each do |ele|
      int_arr = ele[:range].split('-')
      if (int_arr.first.to_i..int_arr.last.to_i).include?(input.last[:size]) && !@capacity_hash[ele[:partner_id]].nil? && input.last[:size].to_i <= @capacity_hash[ele[:partner_id]].to_i
        price = ele[:cost_per_gb].to_i * input.last[:size].to_i
        price = ele[:min_price].to_i if price < ele[:min_price].to_i
        if min_price.nil? || price < min_price
          min_price = price
          partner = ele[:partner_id]
        end
        @output_hash[input.first] = {
            partner: partner,
            min_price: min_price
        }
      end
    end

    if !@capacity_hash[partner].nil?
      @capacity_hash[partner] = @capacity_hash[partner].to_i - input.last[:size].to_i
    end
  end


  row_count.times do |i|
    if !@output_hash["D#{i+1}"].nil?
      puts "D#{i+1}, true, #{@output_hash["D#{i+1}"][:partner]}, #{@output_hash["D#{i+1}"][:min_price]}"
    else
      puts "D#{i+1}, false, '', '' "
    end
  end
end