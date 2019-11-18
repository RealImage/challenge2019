require 'spec_helper'

describe Theatre do
  let(:partners) do
    [
      { theatre: 'T1', size_slab_lower: 0, size_slab_higher: 100, min_cost: 1500, cost_per_gb: 20, partner_id: 'P1' },
      { theatre: 'T1', size_slab_lower: 100, size_slab_higher: 200, min_cost: 2000, cost_per_gb: 13, partner_id: 'P2' }
    ]
  end
  subject { described_class.new('T1') }

  context '#possible_partners_list' do
    it 'returns an array' do
      expect(subject.possible_partners_list(partners, 10_000).class).to eq(Array)
    end

    it 'gives the list of partners size slab is in the range' do
      expect(subject.possible_partners_list(partners, 100).map { |e| e[:partner_id] }).to eq %w[P1 P2]
    end

    it 'excludes the partners who is not in size slab' do
      expect(subject.possible_partners_list(partners, 1000).map { |e| e[:partner_id] }).to eq []
    end
  end

  context '#cost_of_delivery' do
    it 'calculates cost of delivery as min cost when actual cost is too low' do
      expect(subject.cost_of_delivery(partners, 10).map { |e| e[:cost_of_delivery] }).to eq([1500, 2000])
    end

    it 'calculates actual cost of delivery' do
      expect(subject.cost_of_delivery(partners, 200).map { |e| e[:cost_of_delivery] }).to eq([4000, 2600])
    end
  end

  context '#minimum_cost_of_delivery' do
    it 'throws exception if cost_of_delivery is nil' do
      expect { subject.minimum_cost_of_delivery(partners)[:partner_id] }.to raise_error(RuntimeError)
    end

    it 'find the partner_id for cost is min #1' do
      calculated_partners = subject.cost_of_delivery(partners, 10)
      expect(subject.minimum_cost_of_delivery(calculated_partners)[:partner_id]).to eq('P1')
    end

    it 'find the partner_id for cost is min #2' do
      calculated_partners = subject.cost_of_delivery(partners, 200)
      expect(subject.minimum_cost_of_delivery(calculated_partners)[:partner_id]).to eq('P2')
    end
  end
end
