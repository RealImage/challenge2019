require 'csv'
require 'pry'

PROBLEM_STATEMENT_ONE_OUTPUT_FILE = 'output1.csv'
PROBLEM_STATEMENT_TWO_OUTPUT_FILE = 'output2.csv'
PARTNERS_CSV = 'partners.csv'
INPUT_CSV = 'input.csv'
CAPACITY_CSV = 'capacities.csv'

class ProcessCsv
  attr_accessor :file, :klass, :header

  def initialize(file_name, klass: 'challenge.csv', header: true)
    self.file = file_name
    self.klass = klass
    self.header = header
  end

  def read
    datas = []
    CSV.read(file, headers: header).each{ |row| datas << klass.new(row) }
    datas
  end

  def write(data)
    puts data.join(', ')
    CSV.open(file, 'a+'){ |row| row << data }
  end

  def reset_file
    File.truncate(file, 0)
  end
end

class Partner
  attr_accessor :theatre, :min_size_slab, :max_size_slab, :min_cost, :cost_per_gb, :partner_id, :delivery_cost, :available_slab

  def initialize(opts = {})
    self.theatre = opts[0].strip
    size_slab = opts[1].strip.split('-')
    self.min_size_slab = size_slab[0].to_i
    self.max_size_slab = size_slab[1].to_i
    self.min_cost = opts[2].strip.to_i
    self.cost_per_gb = opts[3].strip.to_i
    self.partner_id = opts[4].strip
  end

  def slab_within_limit?(delivery_size)
    delivery_size.between?(min_size_slab, max_size_slab)
  end

  def set_cost_of_delivery(size)
    cost = size * cost_per_gb
    self.delivery_cost = (cost > min_cost ? cost : min_cost)
  end

  def set_available_slab
    self.available_slab = max_size_slab - min_size_slab
  end
end

class Capacity
  attr_accessor :partner_id, :limit

  def initialize(opts = {})
    self.partner_id = opts[0].strip
    self.limit = opts[1].strip.to_i
  end
end

class Input
  attr_accessor :delivery_id, :size, :theatre_id

  def initialize(opts = {})
    self.delivery_id = opts[0].strip
    self.size = opts[1].to_i
    self.theatre_id = opts[2].strip
  end
end

class Challenge
  attr_accessor :partners, :inputs, :group_partners_by_theatre

  def initialize
    self.partners = ProcessCsv.new(PARTNERS_CSV, klass: Partner).read
    self.inputs = ProcessCsv.new(INPUT_CSV, klass: Input, header: false).read
    group_partners_by_theatre
  end

  private

  def group_partners_by_theatre
    self.group_partners_by_theatre = partners.group_by{|partner| partner.theatre }
  end
end

class ProblemStatementOne < Challenge
  attr_accessor :output

  def initialize
    super
    self.output = ProcessCsv.new(PROBLEM_STATEMENT_ONE_OUTPUT_FILE)
  end

  def result
    puts "#{'*' * 10} Problem statement one"
    inputs.map do |input|
      begin
        partner = group_partners_by_theatre[input.theatre_id]&.select do |partner|
          partner.set_cost_of_delivery(input.size)
          partner.slab_within_limit?(input.size)
        end&.min_by(&:delivery_cost)
        self.output.write([input.delivery_id, !!partner, (partner ? [partner.partner_id, partner.delivery_cost] : ['', ''])].flatten)
      rescue => e
        puts 'Error!!! ', e.inspect
      end
    end
  end
end

class ProblemStatementTwo < Challenge
  attr_accessor :capacities, :output, :grouped_theatre_inputs

  def initialize
    super
    self.capacities = ProcessCsv.new(CAPACITY_CSV, klass: Capacity).read
    self.output = ProcessCsv.new(PROBLEM_STATEMENT_TWO_OUTPUT_FILE)
    group_inputs_by_theatre
  end

  def result
    puts "#{'*' * 10} Problem statement two"
    inputs.map do |input|
      begin
        row = if total_inputs_size(input.theatre_id) > available_slab(input.theatre_id)
          [input.delivery_id, false, '', '']
        else
          partner = filter_by_capacity(input)
          [input.delivery_id, !!partner, (partner ? [partner.partner_id, partner.delivery_cost] : ['', ''])].flatten
        end
        self.output.write(row)
      rescue => e
        puts 'Error!!! ', e.inspect
      end
    end
  end

  private

  def available_slab(theatre_id)
    group_partners_by_theatre[theatre_id].sum(&:max_size_slab)
  end

  def total_inputs_size(theatre_id)
    grouped_theatre_inputs[theatre_id].sum(&:size)
  end

  def filter_by_capacity(input)
    result = group_by_delivery_cost(input)
    result.select do |_k, v|
      v.select{|partner| partner.available_slab > capacity_limit(partner) }
    end&.values.flatten.min_by(&:available_slab)
  end
  
  def capacity_limit(partner)
    capacities.find{|capacity| capacity.partner_id == partner.partner_id}.limit
  end

  def filtered_partners(input)
    group_partners_by_theatre[input.theatre_id]&.select do |partner|
      partner.set_cost_of_delivery(input.size)
      partner.set_available_slab
      partner.slab_within_limit?(input.size)
    end
  end

  def group_by_delivery_cost(input)
    result = {}
    filtered_partners = filtered_partners(input)
    filtered_partners.permutation(filtered_partners.size).to_a.collect do |partner|
      result[partner.sum(&:delivery_cost)] = partner
    end
    result
  end

  def group_inputs_by_theatre
    self.grouped_theatre_inputs = inputs.group_by{|partner| partner.theatre_id }
  end
end


ProcessCsv.new(PROBLEM_STATEMENT_ONE_OUTPUT_FILE).reset_file
ProcessCsv.new(PROBLEM_STATEMENT_TWO_OUTPUT_FILE).reset_file

ProblemStatementOne.new.result
ProblemStatementTwo.new.result
