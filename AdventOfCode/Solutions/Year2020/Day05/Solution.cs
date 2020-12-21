using System;
using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day05 : ASolution
    {
        private readonly List<long> seats;

        public Day05() : base(05, 2020, "Binary Boarding")
        {
            // new solution
            seats = Input.SplitByNewline().Select(line =>
                    Convert.ToInt64(line
                        .Replace('F', '0')
                        .Replace('B', '1')
                        .Replace('L', '0')
                        .Replace('R', '1'), 2))
                .ToList();
            seats.Sort();
        }

        protected override string SolvePartOne()
        {
            return seats[^1].ToString();
        }

        protected override string SolvePartTwo()
        {
            for (var i = 1; i < seats.Count; i++)
            {
                if (seats[i - 1] + 1 != seats[i])
                {
                    return (seats[i - 1] + 1).ToString();
                }
            }

            return "error";
        }
    }
}
