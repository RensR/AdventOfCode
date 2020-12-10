using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day10 : ASolution
    {
        List<long> input;
        public Day10() : base(10, 2020, "")
        {
            input = Input.SplitByNewline().Select(long.Parse).ToList();
            input.Add(0);
            input.Add(input.Max() + 3);
            input.Sort();
        }

        protected override string SolvePartOne()
        {
            var one = 0;
            var three = 0;
            for(int i = 0; i < input.Count - 1; i++)
            {
                var diff = input[i + 1] - input[i];
                switch (diff)
                {
                    case 1:
                        one++;
                        break;
                    case 3:
                        buffer[input[i + 1]] = i + 1;
                        three++;
                        break;
                }
            }

            return (one * three).ToString();
        }

        protected override string SolvePartTwo()
        {
            return ConnectAll(input.ToArray()).ToString();
        }

        Dictionary<long, int> buffer = new Dictionary<long, int>();

        private long ConnectAll(long[] next)
        {
            long steps = 1;
            int fromIndex = 0;
            long prevItem = 0;
            foreach(var bufferItem in buffer.Keys)
            {
                var index = buffer[bufferItem];

                steps *= ConnectNext(prevItem, next[(fromIndex + 1)..(index + 1)]);
                fromIndex = index;
                prevItem = bufferItem;
            }
            return steps;
        }

        private long ConnectNext(long current, long[] next)
        {
            if (next.Length == 1)
            {
                if (next[0] - current <= 3)
                    return 1;
                else
                    return 0;
            }
            if (next[0] - current > 3)
                return 0;

            return ConnectNext(current, next[1..]) + ConnectNext(next[0], next[1..]);
        }
    }
}
