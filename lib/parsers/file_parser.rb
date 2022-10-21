module Parsers
  class FileParser
    attr_reader :file_name, :opts
    attr_accessor :content

    def initialize(file_name:, **opts)
      @file_name = file_name
      @opts = opts
      @content = []
    end

    def parse
      raise NotImplementedError
    end
  end
end
