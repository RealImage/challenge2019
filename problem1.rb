class Problem1
  require 'csv'

  def solution
    _hash = {}
    _capacity_hash = {}
    _input_hash = {}
    _total_capacity = 0

    CSV.new(open('./partners.csv'), headers: :first_row).each do |row|
      _hash[row['Theatre'].rstrip.lstrip] = [] if _hash[row['Theatre'].rstrip.lstrip].nil?
      _hash[row['Theatre'].rstrip.lstrip] << {
        range: row['Size Slab (in GB)'].rstrip.lstrip,
        min_price: row['Minimum cost'].rstrip.lstrip,
        cost_per_gb: row['Cost Per GB'].rstrip.lstrip,
        partner_id: row['Partner ID'].rstrip.lstrip
      }
    end

    CSV.new(open('./input.csv')).each do |row|
      row.map!{ |r| r.rstrip.lstrip }
      unless _hash.keys.include?(row.last)
        puts row.first, false, '', ''
        next
      end

      _min_price = nil
      _partner = ''
      _hash[row.last].each do |data|
        _int_arr = data[:range].split('-')
        if (_int_arr.first.to_i.._int_arr.last.to_i).include?(row[1].to_i)
          _price = data[:cost_per_gb].to_i * row[1].to_i
          _price = data[:min_price].to_i if _price < data[:min_price].to_i
          if _min_price.nil? || _price < _min_price
            _min_price = _price
            _partner = data[:partner_id]
          end
        end
      end
      puts "#{row[0]}, #{!_min_price.nil?}, #{_partner}, #{_min_price}"
    end
  end

end

Problem1.new.solution