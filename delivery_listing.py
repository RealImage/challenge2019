import csv
import warnings
import pandas as pd

#: Avoid futurewarnings
warnings.simplefilter(action='ignore', category=Warning)


def splitStr(data):
    """Split the size in GB
    Function change the datatype from str to tuple
    for compraing the incoming size data.
    """
    value_split = data.split('-')
    min = int(value_split[0])
    max = int(value_split[1])
    return min, max


def d_frame(df, data_size, theatre):
    """Process the partners sheet based on the incoming
    data
    df : pandas datafram
    data_size : incoming data size
    theatre : theatre code
    """
    def checkData(data):
        """Incoming data size validity check
        This function check whether the incoming
        data belongs in a range.
        """
        if data_size in range(data[0], data[1]):
            return True
        else:
            return False

    #: Filter the data based on the incoming theare code
    t_data = df[df['Theatre'] == theatre]
    #: Add new column 'status' which set to Boolean
    t_data['status'] = t_data.sizeSplit.apply(checkData)
    #: Consider only those rows which are status = True
    #: This ensure only those rows are cosidered where incoming
    #: data size is with in the range
    filtered = t_data[t_data['status'] == True]
    #: No records found where incoming data size is in range
    #: for any of the threate code
    if len(filtered.index) == 0:
        status = False
        cost = ""
        partner = ""
    else:
        #: Calculate the total cost and assign to new column 'total
        filtered['total'] = filtered['CostPerGB'] * data_size
        #: Cosider only those rows which gives minimum total
        final = filtered[filtered['total'] == filtered['total'].min()]
        #: Check whether there are more rows with same minimum total
        if len(final.index) > 1:
            #: Consider the first entry amongs the rows
            final = final.iloc[[0]]
        status = True
        final.iloc[0]
        #: Check whether total cost is less than minimum Cost
        #: If true, consider minimum cost, else total cost
        if final.iloc[0]['total'] < final.iloc[0]['Minimumcost']:
            cost = final.iloc[0]['Minimumcost']
        else:
            cost = final.iloc[0]['total']
        partner = final.iloc[0]['PartnerID']
    return status, partner, cost


#: Used to feed the newly generating output csv file
csv_feeder = []

with open('input.csv', mode='r') as infile:
    reader = csv.reader(infile)
    df = pd.read_csv('partners.csv')
    #: Clean up the data by removing unwanted spaces from heading
    #: as well as from the elements in the required columns
    df['Size Slab (in GB)'] = df['Size Slab (in GB)'].str.strip()
    df['Theatre'] = df['Theatre'].str.strip()
    df.columns = df.columns.str.replace(' ', '')
    #: Split the SizeSlab -> Conversion from str to tuple
    df['sizeSplit'] = df['SizeSlab(inGB)'].apply(splitStr)
    #: Loop through the input file and clean by removing unwanted spaces
    for deck in reader:
        ref_deck = deck.copy()
        d_id = deck[0].strip()
        data_size = deck[1].strip()
        theatre = deck[2].strip()
        # Get the status, partner and cost for this specific theatre
        status, partner, cost = d_frame(df, int(data_size), theatre)
        item_dict = {
            "delivery_id": d_id,
            "status": status,
            "partner": partner,
            "cost": cost
        }
        csv_feeder.append(item_dict)

with open('sampleOut.csv', 'w', newline='') as file:
    writer = csv.writer(file)
    #: Write each delivery data to the csv file
    for item in csv_feeder:
        writer.writerow([str(item['delivery_id']), str(item['status']),
                        str(item['partner']), str(item['cost'])])
