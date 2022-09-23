# frozen_string_literal: true

# Database with methods to read and write
class Database
  attr_reader :data

  @@row_id = 0
  def initialize(hash_collection = [])
    @data = hash_collection
  end

  private

  def insert(hash)
    @@row_id += 1
    @data.push({ 'id' => @@row_id }.merge(hash))
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

  def find_by_value(value:, keys: [])
    filterd_data = @data.select { |hash| hash.value?(value) }
    return filterd_data if keys.empty?

    filterd_data.map { |hash| hash.slice(*keys) }
  end

  def remove(key_value_pair)
    @data.delete_if { |hash| pair_present?(hash, key_value_pair) }
  end

  def pair_present?(hash, key_value_pair)
    Enumerable.instance_method(:include?).bind(hash).call(key_value_pair.to_a.flatten)
  end
end
