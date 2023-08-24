# frozen_string_literal: true

module LatexYearlyPlanner
  module TeX
    class TabularX < Tabular
      attr_accessor :width

      def initialize(**options)
        super

        @width = options.fetch(:width, '\\textwidth')
      end

      def to_s
        <<~LATEX
          {\\renewcommand{\\arraystretch}{#{vertical_stretch}}\\setlength{\\tabcolsep}{#{horizontal_spacing}}%
          \\begin{tabularx}{#{width}}{#{make_format}}
            #{build_rows}
          \\end{tabularx}}
        LATEX
          .strip
      end
    end
  end
end
