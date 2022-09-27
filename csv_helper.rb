require 'csv'

module CsvHelper
  def cook_csv_data(filename, header)
    csv = CSV.read(
      filename,
      headers: header,
      converters: [:numeric, lambda {|f| f ? f.strip : nil}],
      header_converters: [:symbol, lambda {|f| f ? f.strip : nil}]
    )
  end

  def write_to_csv_and_print_content(filename, contents, message = "")
    CSV.open(filename, "w") do |csv|
      contents.each do |content|
        csv << content
      end
    end

    p "#{message} #{contents}"
  end

end
