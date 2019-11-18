require 'spec_helper'

describe CSVHelper do
  include CSVHelper

  context '#parse_capacities' do
    it 'returns a hash' do
      expect(parse_capacities('capacities.csv').class).to eq(Hash)
    end

    it 'returns a appropriate value' do
      expect(parse_capacities('capacities.csv')).to eq('P1' => 350, 'P2' => 500, 'P3' => 1500)
    end
  end

  context '#convert_to_csv' do
    it 'creates file in tmp directory' do
      convert_to_csv([[1]])
      expect(File).to exist('tmp/output.csv')
    end

    it 'creates file with appropriate value' do
      convert_to_csv([[1, 2], [3, 4]])
      data = File.read('tmp/output.csv')

      expect(data).to eq("1,2\n3,4\n")
    end
  end

  context '#parse_input' do
    it 'returns an array' do
      expect(parse_input('spec/data/input.csv').class).to eq(Array)
    end

    it 'returns a appropriate value' do
      expect(parse_input('spec/data/input.csv')).to eq([
                                                         { distribution_id: 'D1',
                                                           required_size: 150,
                                                           theatre: 'T1' }
                                                       ])
    end
  end
end
