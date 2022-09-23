# frozen_string_literal: true

require './database'
require 'pp'
# Information system to display menu and output results
class InformationSystem < Database
  def initialize
    super
  end

  def display_menu
    puts "What would you like to do?
      1:Problem Statement 1
      2:Problem Statement 2
      Any other key to exit"
    choice = gets
    process_choice(choice)
  end

  private

  def process_choice(choice)
    case choice.to_i
    when 1
      pp problem_statement_1
      display_menu
    when 2
      pp problem_statement_2
      display_menu
    else
      exit
    end
  end

  def problem_statement_1
    
  end

  def problem_statement_2
   
  end
end

c1 = InformationSystem.new
c1.display_menu