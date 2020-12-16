using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day15 : ASolution
    {
        readonly Dictionary<int, int> LastSpoken = new Dictionary<int, int>();
        int previous;
        int i = 0;

        public Day15() : base(15, 2020, "Rambunctious Recitation")
        {
            var numbers = Input.Split(',').ToList().Select(int.Parse).ToArray();

            for(i = 1; i <= numbers.Length; i++)
            {
                if(i == numbers.Length)
                    previous = numbers[i - 1];
                else
                    LastSpoken.Add(numbers[i - 1], i);
            }
        }

        protected override string SolvePartOne()
        {
           return SolveFor(2020).ToString();
        }

        private int SolveFor(int turn)
        {
            for (; i <= turn; i++)
            {
                if (!LastSpoken.ContainsKey(previous))
                {
                    LastSpoken[previous] = i - 1;
                    previous = 0;
                    continue;
                }
                var lastIndex = LastSpoken[previous];

                var diff = i - 1 - lastIndex;
                LastSpoken[previous] = i - 1;
                previous = diff;
            }


            return previous;
        }

        // Because 30000000 > 2020 and part two runs after part one we don't have to reset anything
        // and can start at 2020 for the solution vor part 2.
        protected override string SolvePartTwo()
        {
            return SolveFor(30000000).ToString();
        }
    }
}
