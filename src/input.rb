class Input
  attr_accessor :delivery_id, :req_slab, :theatre
  def initialize(delivery_id:, req_slab:, theatre:)
    @delivery_id = delivery_id
    @req_slab = req_slab
    @theatre = theatre
  end

  def self.input_data
    inputs  = []
    filepath = File.join(Dir.pwd, 'input.csv')
    CSV.foreach(filepath, headers: false) do |row|
    input = Input.new(
        delivery_id: row[0].strip,
        req_slab: row[1].strip.to_i,
        theatre: row[2].strip,
      )
      inputs << input
    end
    inputs
  end
end