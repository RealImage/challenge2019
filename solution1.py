'''
To run:
1. create python virtual environment and activate it.
python -m venv qube-venv
cd qube-venv/Scripts
activate

2. install requirements.txt
<cd to the repository folder>
pip install -r requirements.txt

3. run the script: solutions1.py
python solutions.py

4. output could be seen on the console(Solved output 1)
or in the file(solved_output1.csv)
'''

import pandas as pd

class Solution:

    def __init__(
        self, input_file,
        partners_file,
        capacities_file=None
    ):
        self.input_df = pd.read_csv(input_file, header=None)
        self.input_df.columns = [
            'Delivery ID',
            'Required Size',
            'Theatre'
        ]
        print(f'Input DF:\n{self.input_df}')
        self.partners_df = pd.read_csv(partners_file)
        print(f'Partners DF:\n{self.partners_df}')
        if capacities_file:
            capacities_df = pd.read_csv(capacities_file)
            self.capacities_dict = \
                capacities_df.set_index('Partner ID').to_dict()
            print(f'Capacities Dict:\n{self.capacities_dict}')

    def get_possible_slab(
        self, slab_str,
        required_size
    ):
        '''
        Returns True if required size is under slab
        slab: a-b, True if a <= required_size <= b
        '''
        r1 =int(slab_str.split('-')[0])
        r2 = int(slab_str.split('-')[1]) + 1
        if required_size in range(r1, r2):
            return True
        return False

    def calculate_cost(
        self, minimum_cost,
        cost_per_gb, required_size
    ):
        '''
        cost = cost_per_gb * required_size
        returns minimum_cost if cost <= minimum_cost
        '''
        cost = cost_per_gb * required_size
        if cost > minimum_cost:
            return cost
        else:
            return minimum_cost

    def get_possible_partner_with_optimized_cost(
        self, required_size,
        theatre
    ):
        '''
        Filters record with min cost for
        a given required_size
        '''
        possible_partners_df = self.partners_df[self.partners_df[
            'Theatre'
        ].str.contains(theatre)]
        possible_partners_df = possible_partners_df.loc[
            possible_partners_df['Size Slab (in GB)'].apply(
                lambda x: self.get_possible_slab(
                x, required_size
            ))
        ]
        if possible_partners_df.empty:
            lst = [None] * 6
            lst.append('FALSE')
            return lst

        possible_partners_df['Cost'] = possible_partners_df[
            ['Minimum cost', 'Cost Per GB']
        ].apply(
                lambda x: self.calculate_cost(
                    x[0], x[1],
                    required_size
                ),
                axis=1
            )
        possible_partners_df = possible_partners_df[
            possible_partners_df['Cost'] == \
                possible_partners_df['Cost'].min()
        ].iloc[0]
        possible_partners_df['Is delivery possible'] = 'true'
        return possible_partners_df
        
    def create_output1(self):
        '''
        Creates solved_output1.csv with following attributes:
        ['Delivery ID', 'Is delivery possible', 'Partner ID', 'Cost']
        '''
        self.input_df[[
            'Theatre', 'Size Slab (in GB)',
            'Minimum cost', 'Cost Per GB',
            'Partner ID', 'Cost',
            'Is delivery possible'
        ]] = self.input_df[['Required Size', 'Theatre']].apply(
            lambda x: self.get_possible_partner_with_optimized_cost(
                x[0], x[1]
            ),
            axis=1,
            result_type='expand'
        )

        self.input_df = self.input_df[[
            'Delivery ID', 'Is delivery possible',
            'Partner ID', 'Cost'
        ]]
        # self.input_df['Is delivery possible'] = \
        #     self.input_df['Is delivery possible'].astype(str)
        print(f'Solved output 1:\n{self.input_df}')
        self.input_df.to_csv(
            'solved_output1.csv',
            header=None,
            index=False
        )


if __name__ == '__main__':
    solution_obj1 = Solution('input.csv', 'partners.csv',)
    solution_obj1.create_output1()
