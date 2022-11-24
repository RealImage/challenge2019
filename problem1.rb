class Problem1
  require 'csv'

  @hash = {}
  @capacity_hash = {}
  @input_hash = {}
  @total_capacity = 0
  ###### Form hash with Theaters list ##########
  CSV.new(open('partners.csv'),  liberal_parsing: true, headers: :first_row).each do |row|
    @hash[row['Theatre'].rstrip.lstrip] = [] if @hash[row['Theatre'].rstrip.lstrip].nil?
    @hash[row['Theatre'].rstrip.lstrip] << { range: row['Size Slab (in GB)'].rstrip.lstrip, min_price: row['Minimum cost'].rstrip.lstrip,
                               cost_per_gb: row['Cost Per GB'].rstrip.lstrip, partner_id: row['Partner ID'].rstrip.lstrip }
  end



  ###### Process the output ########
  CSV.new(open('input.csv'),  liberal_parsing: true).each do |row|
    row.map!{ |r| r.rstrip.lstrip }

    unless @hash.keys.include?(row.last)
      puts row.first, false, '', ''
      next
    end

    min_price = nil
    partner = ''
    @hash[row.last].each do |ele|
      int_arr = ele[:range].split('-')
      if (int_arr.first.to_i..int_arr.last.to_i).include?(row[1].to_i)
        price = ele[:cost_per_gb].to_i * row[1].to_i
        price = ele[:min_price].to_i if price < ele[:min_price].to_i
        if min_price.nil? || price < min_price
          min_price = price
          partner = ele[:partner_id]
        end
      end
    end
    puts "#{row[0]}, #{!min_price.nil?}, #{partner}, #{min_price}"
  end
end