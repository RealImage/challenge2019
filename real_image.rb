require 'open-uri'
require 'rubygems'
require 'nokogiri'

class RealImage
  def initialize
    @partner_api = "https://github.com/RealImage/challenge2019/blob/master/partners.csv".freeze
    @input_api = "https://github.com/RealImage/challenge2019/blob/master/input.csv".freeze
    @capacities_api = "https://github.com/RealImage/challenge2019/blob/master/capacities.csv".freeze
  end

  def deliver_content_to_theatre
    puts "*********************************************************************"
    puts "Results fot the First Problem"
    puts "---------------------------------------------------------------------"
    problem_1_sent_content
    puts "---------------------------------------------------------------------"
    puts "Results fot the Second Problem"
    puts "*********************************************************************"
    problem_2_send_content
    puts "---------------------------------------------------------------------"
  end

  private

  def problem_2_send_content
    capacity_details = parse_and_read_capacities_table
    partner_details = parse_and_read_partner_table
    sorted_input = sort_input_table parse_and_read_input_table
    final_out_put_for_problem_2 = []
    check_limit = {}
    sorted_input.each do | input |
      temp_hash = {}
      sent_detail = send_content_to_theater input['delivery_id'], input['content_size'], input['theatre_id'], partner_details
      if !sent_detail.empty?
        sent_detail.group_by{|k, v| v}.min.last.each do |sending_detail|
          if check_limit[sending_detail[0]] == nil
            check_limit[sending_detail[0]] = input['content_size']
          else
            check_limit[sending_detail[0]] =  check_limit[sending_detail[0]] + input['content_size']
          end
          if check_limit[sending_detail[0]] < parse_and_read_capacities_table[sending_detail[0]]
            temp_hash['delivery_id'] = input['delivery_id']
            temp_hash['patner_name'] = sending_detail[0]
            temp_hash['minimum_cost'] = sending_detail[1]
            temp_hash['possible'] = 'true'
          else
            check_for_remaining = send_content_to_theater input['delivery_id'], input['content_size'], input['theatre_id'], partner_details.delete_if{ |h| h["patner_name"] == sending_detail[0] }
            check_for_remaining.group_by{|k, v| v}.min.last.each do |sending_detail|
              if check_limit[sending_detail[0]] == nil
                check_limit[sending_detail[0]] = input['content_size']
              else
                check_limit[sending_detail[0]] =  check_limit[sending_detail[0]] + input['content_size']
              end
              if check_limit[sending_detail[0]] < parse_and_read_capacities_table[sending_detail[0]]
                temp_hash['delivery_id'] = input['delivery_id']
                temp_hash['patner_name'] = sending_detail[0]
                temp_hash['minimum_cost'] = sending_detail[1]
                temp_hash['possible'] = 'true'
              end
            end
          end
        end
      else
        temp_hash['delivery_id'] = input['delivery_id']
        temp_hash['patner_name'] = '""'
        temp_hash['minimum_cost'] = '""'
        temp_hash['possible'] = 'false'
      end
      final_out_put_for_problem_2 << temp_hash
    end
    final_out_put_for_problem_2.each do | out_put |
      puts "DeliveryID :#{out_put['delivery_id']} DeliveryPossible :#{out_put['possible']} PartnerId :#{out_put['patner_name']} MinimumCost :#{out_put['minimum_cost']}"
    end
  end

  def sort_input_table input_data_to_sort
    input_data_to_sort.sort_by{ |h| [-h['content_size']] }
  end

  def parse_and_read_capacities_table
    capacities_html_document = Nokogiri::HTML.parse(open(@capacities_api))
    capacitie_hash = {}
    capacities_html_document.xpath("//tr[@id != 'LC1']").each do |capacity|
      patner_name = capacity.css('td[2]').text.partition(' ')[0]
      capacitie_hash[patner_name] = capacity.css('td[3]').text.to_i
    end
    capacitie_hash
  end

  def problem_1_sent_content
    partner_details = parse_and_read_partner_table
    input_details = parse_and_read_input_table
    problem_first_out_put = []
    input_details.each do |input_detail|
      temp_hash = {}
      send_details = send_content_to_theater input_detail['delivery_id'], input_detail['content_size'], input_detail['theatre_id'], partner_details
      if !send_details.empty?
        send_details.group_by{|k, v| v}.min.last.each do |send_detail|
          temp_hash['delivery_id'] = input_detail['delivery_id']
          temp_hash['possible'] = 'true'
          temp_hash['patner_name'] = send_detail[0]
          temp_hash['minimum_cost'] = send_detail[1]
        end
      else
        temp_hash['delivery_id'] = input_detail['delivery_id']
        temp_hash['possible'] = 'false'
        temp_hash['patner_name'] = ""
        temp_hash['minimum_cost'] = ""
      end
      problem_first_out_put << temp_hash
    end
    problem_first_out_put.each do | out_put |
      puts "DeliveryID :#{out_put['delivery_id']} DeliveryPossible :#{out_put['possible']} PartnerId :#{out_put['patner_name']} MinimumCost :#{out_put['minimum_cost']}"
    end
  end

  def parse_and_read_input_table
    # getting input from github public remote pfile
    input_html_document = Nokogiri::HTML.parse(open(@input_api))
    input_array = []
    i = 0
    input_html_document.xpath("//tr").each do |input|
      input_hash = {}
      # i == 0 why?; because in input.csv find first data is in 'th' and others are in 'td'
      if i == 0
        delivery_id = input.css('th[2]').text
        content_size = input.css('th[3]').text.to_i
        theatre_id = input.css('th[4]').text
      else
        delivery_id = input.css('td[2]').text
        content_size = input.css('td[3]').text.to_i
        theatre_id = input.css('td[4]').text
      end
      input_hash['delivery_id'] = delivery_id
      input_hash['content_size'] = content_size
      input_hash['theatre_id'] = theatre_id
      input_array << input_hash unless input_hash.empty?
      i+=1
    end
    input_array
  end

  def parse_and_read_partner_table
    partner_html_document = Nokogiri::HTML.parse(open(@partner_api))
    partners_hash = []
    # looping file though rows execpt first row because first row contain header name
    partner_html_document.xpath("//tr[@id !='LC1']").each do |partner|
      temp_hash = {}
      temp_hash['theatre_name'] = partner.css('td[2]').text.partition(' ')[0]
      # making size_slabe as Range for checking content size while sending content
      size_slab = partner.css('td[3]').text.partition('-')[0].to_i..partner.css('td[3]').text.partition('-')[2].to_i
      temp_hash['size_slab'] = size_slab
      temp_hash['minimum_cost'] = partner.css('td[4]').text.to_i
      temp_hash['cost_per_gb'] = partner.css('td[5]').text.to_i
      temp_hash['patner_name'] = partner.css('td[6]').text
      partners_hash << temp_hash
    end
    partners_hash
  end

  def send_content_to_theater delivery_id, content_size, theatre_id, partner_details
    partner_and_cost_to_send = {}
    # setting flag to flase to check if delivery is not possible
    @flag = false
    partner_details.each do | partner |
      content_size_boolean = partner['size_slab'].include?content_size
      theatre_name_boolean = partner['theatre_name'] == theatre_id
      if content_size_boolean && theatre_name_boolean
        @flag = true
        cost = content_size * partner['cost_per_gb']
        sending_cost = cost < partner['minimum_cost'] ? partner['minimum_cost'] : cost
        if partner_and_cost_to_send[partner['patner_name']] == nil
          partner_and_cost_to_send[partner['patner_name']] = sending_cost
        else
          # assigning the value which is less to hash( example: "p1"=> 10 ,  "p1"=>9 , result : "p1"=>9)
          partner_and_cost_to_send[partner['patner_name']] = sending_cost if sending_cost < partner_and_cost_to_send[partner['patner_name']]
        end
      end
    end
    partner_and_cost_to_send
  end
end

RealImage.new.deliver_content_to_theatre

