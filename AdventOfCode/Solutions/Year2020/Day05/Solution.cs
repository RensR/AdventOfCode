using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{

    class Day05 : ASolution
    {
        readonly List<string> boardingPassLines = new List<string>();
        private List<long> seats;

        public Day05() : base(05, 2020, "Binary Boarding")
        {
            boardingPassLines = Input.SplitByNewline();
        }

        protected override string SolvePartOne()
        {
            seats = boardingPassLines.Select(CalculateNumber).ToList();
            return seats.Max().ToString();
        }

        private long CalculateNumber(string input)
        {
            long seat = 0;
            for(int i = 0; i < input.Length; i++)
            {
                if (input[^(i + 1)] is 'B' or 'R')
                {
                    seat += 1 << i;
                }
            }
            return seat;
        }

        protected override string SolvePartTwo()
        {
            seats.Sort();

            for(int i = 1; i < seats.Count; i ++)
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
