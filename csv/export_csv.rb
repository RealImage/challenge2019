require "csv"

module ExportCsv
  def self.csv_string(data)
    CSV.generate do |csv|
      data.each do |row|
        csv << row
      end

      csv
    end
  end

  def self.export(file_name, data)
    File.write(file_name, csv_string(data))
  end
end
