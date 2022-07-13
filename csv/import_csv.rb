require "csv"

module ImportCsv
  def import(file_name, options = {})
    CSV.read(file_name, headers: options[:headers]).map do |row|
      new(row)
    end
  end
end