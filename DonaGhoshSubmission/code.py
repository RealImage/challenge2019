import csv
with open('C:/Users/hp/Desktop/DonaGhoshSubmission/input.csv', 'r') as input:
    input_reader = csv.reader(input)
    rows=[]
    for irow in input_reader:
        with open('C:/Users/hp/Desktop/DonaGhoshSubmission/partners.csv', 'r') as partner:
            partner_reader = csv.reader(partner)
            next(partner_reader)
            min_output=[]
            min_cost=9999999999
            for prow in partner_reader:
                if int(irow[1])>=int(prow[1].split('-')[0]) and int(irow[1])<=int(prow[1].split('-')[1]):
                    if min_cost > int(prow[3])*int(irow[1]):
                        min_cost=int(prow[3])*int(irow[1])
                    #if int(prow[3])*int(irow[1])<int(prow[2]):
                        #min = [irow[0],'true',prow[4],prow[2]]
                    #else:
                        min_output= [irow[0],'true',prow[4],max(min_cost,int(prow[2]))]
            if len(min_output) == 0:
                min_output =[irow[0],'false','','']
            #print(irow[0]," ",min_cost)
        rows.append(min_output)
    print(rows)

filename = "C:/Users/hp/Desktop/DonaGhoshSubmission/output.csv"
with open(filename, 'w') as output: 
    output_writer = csv.writer(output) 
    output_writer.writerows(rows)
                 
    
