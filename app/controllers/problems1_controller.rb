class Problems1Controller < ApplicationController
    def index
        partners = get_partners_data
        final_arr = get_input_data.map do |d_info|
          find_partner_info(d_info, partners)
        end
        # below render is giving output on web page as json and saving data to csv named output1.csv in csv format
        render json:final_arr
        File.write('output1.csv', final_arr.map(&:to_csv).join)
    end
    
    def find_partner_info(deliver_arr_info, partners)
        # below checks if partners info contains the following condition and provide resultant partners array
        theatres_info = partners.select { |arr| arr[:theatre] == deliver_arr_info[2] && arr[:size_lab_min] <= deliver_arr_info[1].to_i && arr[:size_lab_max] >= deliver_arr_info[1].to_i }
      
        # checks if partners are present i.e partners surpasses above condition
        if theatres_info.count > 0   
          delivery_costs = theatres_info.map do |theatre|
            {
              deliver_id: deliver_arr_info[0],
              possible: true,
              partner_id: theatre[:partner_id],
              cost: (deliver_arr_info[1].to_i * theatre[:cost_per_gb]) >= theatre[:minimun_cost] ? (deliver_arr_info[1].to_i * theatre[:cost_per_gb]) : theatre[:minimun_cost]
            }
          end
      
           # select record having min cost in delivery cost
          return delivery_costs.min { |a, b| a[:cost] <=> b[:cost] }.values
        else
            # return nil in case no theatre_info found and marking false for delivery impossible
          return [deliver_arr_info[0], false, '', '']
        end
      end
end
