import pandas as pd

def get_status(inpt,rang):
    """
    Method to find the status of delivery by partner
    arg: size of delivery from input file, Range of size of delivery by partner
    retuns : status as True if size is in range else false
    """
    values = rang.split('-')
    min = int(values[0].strip())
    max = int(values[1].strip())
    if int(inpt) in range(min,max):
        return True
    else:
        return False


def get_status_cost(row, partners,final_list):
    """
    Method to determine cost of delivery and status of delivery
    arg: each row from input csv , partners DF Final_list of deliveries
    retuns : Updated final list with dicts of partners delivers
    """
    for index, partner in partners.iterrows():
        if row['theatre ID'].strip() == partner['Theatre'].strip():
            status = get_status(row["size of delivery"], partner['Size Slab (in GB)'].strip())
            if status:
                cost = int(partner["Minimum cost"]) * int(row["size of delivery"])
                item = {
                        'delivery ID' : row['delivery ID'],
                        'Status' : status,
                        'cost'   : cost,
                        'partner' : partner['Partner ID']
                       }
                final_list.append(item)


def get_list_deliveris(partners,input_data):
    """
    Method to create the output file with partners delivers to theatres
    arg: Input df, partner df
    Output : Creates a ouptu csv with partners delivers to theatres
    """
    final_list = []
    for index, row in input_data.iterrows():
        get_status_cost(row,partners, final_list)
    if len(final_list) > 0:
        df = pd.DataFrame(final_list)
        # fetch the partner who delivers at minimum cost from the 
        # retrieved partners delivery
        out = df.groupby("delivery ID")['cost'].agg('min')
        df2 = df[df.cost.isin(out.tolist())].reset_index(drop=True)
        for index, row in input_data.iterrows():
            if row["delivery ID"] not in df2['delivery ID'].tolist():
                df2.loc[len(df2.index)] = [row["delivery ID"], "False", " ", " "]           
        df2.to_csv('sample_output.csv')

if __name__ == "__main__":
    colu = ["delivery ID", 'size of delivery', 'theatre ID']
    partners = pd.read_csv('partners.csv')
    input_data = pd.read_csv('input.csv' , header = None, names= colu)
    get_list_deliveris(partners,input_data)