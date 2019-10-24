require 'csv'
class Reader

  def initialize(filepath, headers: false)
    @filepath = filepath
    @headers = headers
  end

  def read
    CSV.parse(File.read(@filepath), headers: @headers)
  end

  private
  def build_filepath
    File.join(File.dirname(__FILE__), @filepath)
  end

end