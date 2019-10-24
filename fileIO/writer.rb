require 'csv'
class Writer
  def initialize(filepath, results)
    @filepath = filepath
    @results = results
  end

  def write
    CSV.open(@filepath, "wb") do |csv|
      @results.each do|result|
        values_to_write = result.values
        csv << values_to_write
      end
    end
  end

  private
  def build_filepath
    File.join(File.dirname(__FILE__), @filepath)
  end
end