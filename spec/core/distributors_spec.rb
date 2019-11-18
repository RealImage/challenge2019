require 'spec_helper'

describe Distributors do
  subject { described_class.new('partners.csv', 'capacities.csv') }

  # context '#possible_partners_list' do
  #   it 'returns an array' do
  #     # expect(subject.class).to eq(Array)
  #     pending 'yet to implement'
  #   end
  # end

  context '#group_by_rank' do
    it 'splits by number of possible with count as key' do
      expect(
        subject.group_by_rank(
          [
            { 'D1' => { possible: true, cost_of_delivery: 2000 }, 'D2' => { possible: true, cost_of_delivery: 3600 }, 'D3' => { possible: true, cost_of_delivery: 3900 } },
            { 'D1' => { possible: true, cost_of_delivery: 2000 }, 'D2' => { possible: true, cost_of_delivery: 3600 }, 'D3' => { possible: true, cost_of_delivery: 3900 } },
            { 'D1' => { possible: false, cost_of_delivery: 2000 }, 'D2' => { possible: true, cost_of_delivery: 3600 }, 'D3' => { possible: true, cost_of_delivery: 3900 } },
            { 'D1' => { possible: true, cost_of_delivery: 2000 }, 'D2' => { possible: false, cost_of_delivery: 3600 }, 'D3' => { possible: false, cost_of_delivery: 3900 } }
          ]
        )
      ).to eq(
        1 => [
          { 'D1' => { possible: true, cost_of_delivery: 2000 }, 'D2' => { possible: false, cost_of_delivery: 3600 }, 'D3' => { possible: false, cost_of_delivery: 3900 } }
        ],
        2 => [
          { 'D1' => { possible: false, cost_of_delivery: 2000 }, 'D2' => { possible: true, cost_of_delivery: 3600 }, 'D3' => { possible: true, cost_of_delivery: 3900 } }

        ],
        3 => [
          { 'D1' => { possible: true, cost_of_delivery: 2000 }, 'D2' => { possible: true, cost_of_delivery: 3600 }, 'D3' => { possible: true, cost_of_delivery: 3900 } },
          { 'D1' => { possible: true, cost_of_delivery: 2000 }, 'D2' => { possible: true, cost_of_delivery: 3600 }, 'D3' => { possible: true, cost_of_delivery: 3900 } }

        ]
      )
    end
  end
end
