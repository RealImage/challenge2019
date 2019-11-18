require 'spec_helper'

describe 'Test the script' do
  context 'scenario 1' do
    it 'example 1' do
      find_the_minimum(
        'spec/data/input1.csv',
        'spec/data/partners1.csv',
        'spec/data/capacities1.csv'
      )

      compare_output_csv('spec/data/output1.csv')
    end

    it 'example 2' do
      find_the_minimum(
        'spec/data/input2.csv',
        'spec/data/partners1.csv',
        'spec/data/capacities1.csv'
      )

      compare_output_csv('spec/data/output2.csv')
    end

    it 'example 3' do
      find_the_minimum(
        'spec/data/input3.csv',
        'spec/data/partners1.csv',
        'spec/data/capacities1.csv'
      )

      compare_output_csv('spec/data/output3.csv')
    end
  end

  context 'scenario 2' do
    it 'example 1' do
      find_the_minimum_with_capacity(
        'spec/data/input4.csv',
        'spec/data/partners1.csv',
        'spec/data/capacities1.csv'
      )

      compare_output_csv('spec/data/output4.csv')
    end
  end

  def compare_output_csv(expected_output)
    actual_output = CSV.parse(File.read('tmp/output.csv'))
    expected_output = CSV.parse(File.read(expected_output))

    expect(actual_output.count).to eq(expected_output.count)
    actual_output.count.times do |i|
      actual_row = actual_output[i]
      expected_row = expected_output[i]

      expect(actual_row.count).to eq(expected_row.count)

      actual_row.count.times do |j|
        expect(actual_row[j].to_s.strip).to eq(expected_row[j].to_s.strip)
      end
    end
  end
end
