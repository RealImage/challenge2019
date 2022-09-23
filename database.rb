# frozen_string_literal: true

# Database with methods to read and write
class Database
  attr_reader :data

  @@row_id = 0
  def initialize(hash_collection = [])
    @data = hash_collection
  end

  def find(key_value_pair: nil, keys: [])
    return @data if key_value_pair.nil? && keys.empty?

    unless key_value_pair.nil?
      filterd_data = @data.select { |hash| pair_present?(hash, key_value_pair) }
      return filterd_data if keys.empty?

      return filterd_data.map { |hash| hash.slice(*keys) }
    end

    @data.map { |hash| hash.slice(*keys) }
  end

  private

  def pair_present?(hash, key_value_pair)
    Enumerable.instance_method(:include?).bind(hash).call(key_value_pair.to_a.flatten)
  end
end
