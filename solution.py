from asyncore import read
import pandas as pd
import sys
import csv


'''
    @author Arjun Balasubramanian
    The general idea is to first find which all partners deliver to the target theater
    Once this is found, convert the range into intervals and sort the 2d list thus obtained by the start value of each interval
    Eg: ('0-100') becomes the interval[0, 100] and this comes before [200, 300] so on and so forth.
    Now find all the intervals that can contain the target amount of data. This will be the candidate list.
    Now calculate the cost for each candidate and then finally, return the candidate with minimum possible cost
'''

partners_data = pd.read_csv("partners.csv", skipinitialspace=True)
input_data = pd.read_csv("input.csv", header=None, skipinitialspace=True)

def start():
    output_list = []
    for i in range(len(input_data)):
        output_list.append(find_minimum_cost(i))
    output_to_csv(output_list)

def output_to_csv(output_list):
    fields = ['input', 'is_possible', 'partner', 'cost']
    with open('output.csv', 'w') as csvfile:
        writer = csv.DictWriter(csvfile, fields, delimiter=',', quotechar="\'", quoting=csv.QUOTE_NONE)
        writer.writerows(output_list)
        

def find_minimum_cost(i):
    partners_who_deliver_to_target = []
    target_theatre = input_data.iloc[i,2]
    amount_of_data_to_be_sent = input_data.iloc[i,1]
    candidate_partners = []
    return_map = {}

    for partner_index in range(len(partners_data)):
        if(partners_data.iloc[partner_index,0].strip() == target_theatre.strip()):     
            m = {}
            m['index'] = partner_index
            m['data'] = convert_to_range(partners_data.iloc[partner_index,1])
            partners_who_deliver_to_target.append(m)

    #sort the list based on the start value of each interval
    partners_sorted_based_on_data_range = sorted(partners_who_deliver_to_target, key=lambda x: x['data'][0])

    for k in partners_sorted_based_on_data_range:
        if(amount_of_data_to_be_sent >= k['data'][0] and amount_of_data_to_be_sent <= k['data'][1]):
            candidate_partners.append(k)
    
    if(len(candidate_partners) == 0):
        return_map['is_possible'] = "false"
        return_map['input'] = input_data.iloc[i,0]
        return_map['partner'] = '\"\"'
        return_map['cost'] = '\"\"'
        return return_map
    else:
        return_map['is_possible'] = "true"

    for partner_index in candidate_partners:
        min_cost_for_this_partner = partners_data.iloc[partner_index['index'], 2]
        calculated_cost = partners_data.iloc[partner_index['index'], 3] * amount_of_data_to_be_sent
        if (calculated_cost < min_cost_for_this_partner):
            partner_index['cost'] = min_cost_for_this_partner
        else:
            partner_index['cost'] = calculated_cost

    min_cost_partner = None
    min_cost = 99999999
       
    for g in candidate_partners:
        if(g['cost'] < min_cost):
            min_cost = g['cost']
            min_cost_partner = g
    return_map['input'] = input_data.iloc[i,0]
    return_map['partner'] = partners_data.iloc[min_cost_partner['index'],4]
    return_map['cost'] = min_cost
    return return_map
    
'''
    given a string range such as '0-100', convert it to the list [0, 100]
'''
def convert_to_range(str_range):
    elem = []
    elem.append(int(str_range.strip().split('-')[0]))
    elem.append(int(str_range.strip().split('-')[1]))
    return elem;


if __name__ == "__main__":
    start()