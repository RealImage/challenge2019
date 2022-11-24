class Problem2
  require 'csv'

  def solution
    _hash = {}
    _capacity_hash = {}
    _input_hash = {}
    _output_hash = {}
    _total_capacity = 0
    _row_count = 0

    CSV.new(open('./partners.csv'), headers: :first_row).each do |row|
      _hash[row['Theatre'].rstrip.lstrip] = [] if _hash[row['Theatre'].rstrip.lstrip].nil?
      _hash[row['Theatre'].rstrip.lstrip] << {
        range: row['Size Slab (in GB)'].rstrip.lstrip,
        min_price: row['Minimum cost'].rstrip.lstrip,
        cost_per_gb: row['Cost Per GB'].rstrip.lstrip,
        partner_id: row['Partner ID'].rstrip.lstrip
     }
    end

    CSV.new(open('./capacities.csv'), headers: :first_row).each do |row|
      _capacity_hash[row['Partner ID'].rstrip.lstrip] = row['Capacity (in GB)'].rstrip.lstrip
      _total_capacity += row['Capacity (in GB)'].rstrip.lstrip.to_i
    end

    CSV.new(open('./input.csv')).each do |row|
      _total_capacity -= row[1].to_i
      _input_hash[row.first] = {size: row[1].to_i, theatre: row.last} if _total_capacity >= 0
      _row_count += 1
    end

    _input_arr = _input_hash.sort_by{ |_, v| v[:size] }

    _input_arr.reverse!.each do |input|
      _min_price = nil
      _partner = ''
      _hash[input.last[:theatre].lstrip.rstrip].each do |data|
        _int_arr = data[:range].split('-')
        if (_int_arr.first.to_i.._int_arr.last.to_i).include?(input.last[:size]) && !_capacity_hash[data[:partner_id]].nil? && input.last[:size].to_i <= _capacity_hash[data[:partner_id]].to_i
          price = data[:cost_per_gb].to_i * input.last[:size].to_i
          price = data[:min_price].to_i if price < data[:min_price].to_i
          if _min_price.nil? || price < _min_price
            _min_price = price
            _partner = data[:partner_id]
          end
          _output_hash[input.first] = {
              partner: _partner,
              min_price: _min_price
          }
        end
      end

      if !_capacity_hash[_partner].nil?
        _capacity_hash[_partner] = _capacity_hash[_partner].to_i - input.last[:size].to_i
      end
    end


    _row_count.times do |i|
      if !_output_hash["D#{i+1}"].nil?
        puts "D#{i+1}, true, #{_output_hash["D#{i+1}"][:partner]}, #{_output_hash["D#{i+1}"][:min_price]}"
      else
        puts "D#{i+1}, false, '', ''"
      end
    end
  end

end

Problem2.new.solution