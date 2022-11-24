require 'csv'
require 'byebug'

$theatre_partner_list = {}
$delivery_list = []
$capacities = {}

def get_partners(filename)
  # read from csv file
  csv_data = CSV.read(filename)
  csv_data[1..].each do |line_data|
    # if theatre is not there in theatre_partner_list then we add it , if it's present we just add another partner detail
    if !$theatre_partner_list.key?(line_data[0].strip.downcase)
      $theatre_partner_list[line_data[0].strip.downcase] = []
      $theatre_partner_list[line_data[0].strip.downcase] << {
        'partner'=> line_data[-1].strip.downcase,
        'cost_per_gb'=> line_data[-2].strip.to_i,
        'min_cost'=> line_data[-3].strip.to_i,
        'lower_end'=> line_data[1].strip.split('-')[0].to_i,
        'higher_end'=> line_data[1].strip.split('-')[1].to_i
      }
    else
      $theatre_partner_list[line_data[0].strip.downcase] << {
        'partner'=> line_data[-1].strip.downcase,
        'cost_per_gb'=> line_data[-2].strip.to_i,
        'min_cost'=> line_data[-3].strip.to_i,
        'lower_end'=> line_data[1].strip.split('-')[0].to_i,
        'higher_end'=> line_data[1].strip.split('-')[1].to_i
      }
    end
  end
end
 

def get_input(filename)
  # get the delivery input
  csv_data = CSV.read(filename)
  csv_data.each do |line_data|
    $delivery_list << {
      'id' => line_data[0].strip.downcase,
      'gb_to_send' => line_data[1].strip.to_i,
      'theatre' => line_data[2].strip.downcase
    }
  end
end

def get_capacities(filename)
  # get data
  csv_data = CSV.read(filename)
  csv_data = csv_data[1..]
  csv_data.each do |line_data|
    partner = line_data[0].strip.downcase
    capacity = line_data[1].strip.downcase
    $capacities[partner] = capacity.to_i 
  end
end

def calculate_delivery_cost1
  result = []
  $delivery_list.each do |delivery|
    # if theatre to which delivery is assigned doesn't exist then retun false for that
    if $theatre_partner_list.key?(delivery['theatre'])
      cheapest_rate = {'partner' => 'false', 'cost' => ''}
      $theatre_partner_list[delivery['theatre']].each do |partner_details|
        if partner_details['lower_end'] <= delivery['gb_to_send'] && delivery['gb_to_send'] <= partner_details['higher_end']
          cost = delivery['gb_to_send'] * partner_details['cost_per_gb']
          if cost < partner_details['min_cost'] then cost = partner_details['min_cost'] end
          if cheapest_rate['cost'].to_i > cost || cheapest_rate['cost'].to_i == 0
            cheapest_rate['partner'] = partner_details['partner']
            cheapest_rate['cost'] = cost            
          end
        end
      end
      if cheapest_rate['partner'] == 'false'
        result << [delivery['id'].upcase,false,'','']
      else
        result << [delivery['id'].upcase,'true',cheapest_rate['partner'].upcase,cheapest_rate['cost']]
      end
    end
  end
  make_csv_file('output1',result)
end

def calculate_delivery_cost2
  final_result = [] # final result
  # get all possible delivery - GREEDY ALGO
  possible_ways = []
  lists = $delivery_list
  lists.each do |delivery|
    # if theatre to which delivery is assigned doesn't exist then retun false for that
    if $theatre_partner_list.key?(delivery['theatre'])
      $theatre_partner_list[delivery['theatre']].each do |partner_details|
        if partner_details['lower_end'] <= delivery['gb_to_send'] && delivery['gb_to_send'] <= partner_details['higher_end']
          cost = delivery['gb_to_send'] * partner_details['cost_per_gb']
          if cost < partner_details['min_cost'] then cost = partner_details['min_cost'] end
          possible_ways << [delivery['id'].upcase,'true',partner_details['partner'].upcase,cost]
        end
      end
    end
  end
  # now we know possible delievery methods and have stored them in array of arrays
  # but the one that can not be delivered should be added to final result so that i know it can't be delivered
  all_possible_delivery_id = []
  uniq_deliverey_id_list = []
  lists.each{|delivery| uniq_deliverey_id_list << delivery['id'].upcase}
  possible_ways.each { |possible_way|
    all_possible_delivery_id << possible_way[0]
  }
  not_deliverable = uniq_deliverey_id_list - all_possible_delivery_id
  not_deliverable.each { |delivery_id|
    final_result << [delivery_id,false,'','']
    # updating the delievery list too, that this item can't be delievered
    lists.reject! { |delivery| delivery['id'].upcase == delivery_id }
  }
  # now we have added those that cannot be delivered to the list of final result , and updated current delivery lis too
  # now let's make a list of possible ways by key = delivery id and value = number of ways you can deliver that delivery
  temp_list = {}
  possible_ways.each { |way|
    if temp_list.key?(way[0])
      temp_list[way[0]] << way
    else
      temp_list[way[0]] = []
      temp_list[way[0]] << way
    end
  }

  # now we just need to sort our temp_list (which is remaining possible ways) , by number of ways they can be sorted
  temp = temp_list.sort_by {|k,v| v.length}
  temp_list = {}
  temp.each{|value|
    temp_list[value[0]] = value[1]
  }
  # now temp_list contains of possible ways of delivery with least possible way being first in hash, and highest number of ways to be in last
  temp_list.each { |key,value|
    # sort them by highest to lowest cost to minimise cost
    value = value.sort_by {|item| item[-1]}
    value.each {|delivery|
      # debugger
      if $capacities[delivery[-2].downcase] >= lists.select{|item| item['id'] == delivery[0].downcase }[0]['gb_to_send'] 
        final_result << delivery
        # updating capacity of that partner 
        $capacities[delivery[-2].downcase] -= lists.select{|item| item['id'] == delivery[0].downcase }[0]['gb_to_send'] 
        lists.reject! {|del| del['id'] == key.downcase}
        break
      end
    }
  }  
  # now iterate over remaining items in list and add them to our result as they could not be completed
  lists.each {|delivery|
    final_result << [delivery['id'],false,'',''] 
  }
  make_csv_file('output2',final_result)
end

def make_csv_file(name,results)
  CSV.open("#{name}.csv", "w") do |csv|
    results.each do |results|
      csv << results
    end
  end  
end

partner_file = 'partners.csv'
input_file = 'input.csv'
capacities_file = 'capacities.csv'


get_partners(partner_file)
get_input(input_file)
get_capacities(capacities_file)

calculate_delivery_cost1
calculate_delivery_cost2