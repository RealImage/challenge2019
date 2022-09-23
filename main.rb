# frozen_string_literal: true

require './information_system'
begin
  puts 'Initializing information system'
  info = InformationSystem.new
  info.display_menu
rescue StandardError => e
  puts "Exception: #{e}"
ensure
  puts 'Exiting the program'
end
