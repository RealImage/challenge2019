require_relative 'theatre'

class Distributors
  def initialize(partners, capacities)
    @partners = partners
    @capacities = capacities
  end

  def calculate_minimum_delivery(input)
    output = []
    input.each do |row|
      theatre = Theatre.new(row[:theatre])

      possible_partners = theatre.possible_partners_list(@partners, row[:required_size])
      partners_with_cost_of_delivery = theatre.cost_of_delivery(possible_partners, row[:required_size])

      valid_partner = theatre.minimum_cost_of_delivery(partners_with_cost_of_delivery)
      output << [
        row[:distribution_id],
        !valid_partner.nil?,
        valid_partner && valid_partner[:partner_id],
        valid_partner && valid_partner[:cost_of_delivery]
      ]
    end

    output
  end

  def calculate_minimum_delivery_with_capacity(input)
    distributors_list = {}
    input.each do |row|
      theatre = Theatre.new(row[:theatre])

      possible_partners = theatre.possible_partners_list(@partners, row[:required_size])
      partners_with_cost_of_delivery = theatre.cost_of_delivery(possible_partners, row[:required_size])

      distributors_list[row[:distribution_id]] = partners_with_cost_of_delivery
    end

    all_partners = @capacities.keys
    distributors_list_keys = distributors_list.keys

    # Not sure what to do with this.
    # distributors_possible_partners = {}
    # distributors_list.each do |k, v|
    #   distributors_possible_partners[k] = v.map { |p| p[:partner_id] }
    # end

    possible_permutations = []

    all_partners.permutation(all_partners.count).to_a.each do |pl|
      distributors_list_keys.permutation(distributors_list_keys.count).to_a.each do |dl|
        tmp = {}

        capacities = @capacities.dup
        dl.each do |d|
          pl.each do |p|
            next if tmp[d] && tmp[d][:possible] # already possible

            relevant_dl = distributors_list[d].detect { |a| a[:partner_id] == p }

            next if relevant_dl[:size] > capacities[p]

            capacities[p] -= relevant_dl[:size]

            tmp[d] = {
              possible: true,
              cost_of_delivery: relevant_dl[:cost_of_delivery],
              partner_id: relevant_dl[:partner_id],
              theatre: relevant_dl[:theatre],
              size: relevant_dl[:size]
            }
          end
        end
        possible_permutations << tmp
      end
    end

    max_permutations = group_by_rank(possible_permutations).max_by { |a| a[0] }[1]

    max_permutations.min_by { |a| a.values.map { |x| x[:cost_of_delivery] }.inject(:+) }
  end

  def group_by_rank(values)
    rank = {}
    values.each do |val|
      r = val.values.inject(0) { |sum, n| sum + (n[:possible] ? 1 : 0) }

      rank[r] ||= []
      rank[r] << val
    end

    rank
  end
end
