
import csv




partnerss = "E:/challenge2019-master/partners.csv"
input_file = "E:/challenge2019-master/input1.csv"


output = []
cost = 0
theatre = []
theaterid = []

with open(partnerss,'r') as file:
                csvafile = csv.reader(file,delimiter=',')
                partners = list(csvafile)

with open(input_file,'r') as csvfile:
            csd = csv.reader(csvfile,delimiter=',')
            input=list(csd)
            
            
for i in range(len(input)):
            a = input[i]
            delivery=a[0]
            
            contentsize = a[1]
            theaterid = a[2]
            provider = ""
            finalcost = 0
            isdeliverypossible = False
            
            
            for j in range(len(partners)):
                b = partners[j]
                sizeslab=b[1]
                sizeslab = sizeslab.split('-')
                theatre = b[0]
                costpergb = b[3]
                partnerid = b[4]
                minimumcost = b[2]
                
                if theaterid in theatre:
                        if contentsize >= sizeslab[0] and contentsize <= sizeslab[1]:
                           
                            if int(minimumcost) > int(contentsize)*int(costpergb):
                                cost=minimumcost
                                
                            else:
                                cost = int(contentsize)*int(costpergb)
                                
                                
                            if int(finalcost) > int(cost) or not isdeliverypossible:
                                    finalcost=cost
                                    provider=partnerid
                                    isdeliverypossible=True
                                
                            
                                    
            if finalcost==0:
                finalcost = ""
                provider = ""

            output = [(delivery,str(isdeliverypossible),provider,str(finalcost))] 
            print(output)                           
            
            file = open('E:/challenge2019-master/output1.csv', 'a+',newline='')
            with file:
                write = csv.writer(file)
                write.writerows(output)











            


                            


                        
                      

    

        
   
