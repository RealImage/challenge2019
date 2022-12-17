class Input
  attr_accessor :delivery_id, :delivery_size, :theatre_id

  def initialize(delivery_id:, delivery_size:, theatre_id:)
    @delivery_id = delivery_id
    @delivery_size = delivery_size
    @theatre_id = theatre_id
  end

  def self.get_inputs
    inputs_arr = []
    CSV.foreach(File.join(Dir.pwd, 'input.csv'), headers: false) do |row|
      input = Input.new(
        delivery_id: row[0].strip,
        delivery_size: row[1].strip.to_i,
        theatre_id: row[2].strip
        )
      inputs_arr << input
    end
    inputs_arr
  end
end