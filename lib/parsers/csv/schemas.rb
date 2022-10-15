module Parsers
  module Csv
    module Schemas
      INPUT_SCHEMA = {
        delivery_id: {
          index: 0,
          type: ::String
        },
        delivery_size: {
          index: 1,
          type: ::Integer
        },
        theatre_id: {
          index: 2,
          type: ::String
        }
      }.freeze
      PARTNERS_SCHEMA = {
        theatre_id: {
          index: 0,
          type: ::String
        },
        slab_size: {
          index: 1,
          type: ::Range
        },
        min_cost: {
          index: 2,
          type: ::Integer
        },
        gb_cost: {
          index: 3,
          type: ::Integer
        },
        partner_id: {
          index: 4,
          type: ::String
        }
      }.freeze
      CAPACITY_SCHEMA = {
        partner_id: {
          index: 0,
          type: ::String
        },
        capacity: {
          index: 1,
          type: ::Integer
        }
      }
    end
  end
end
