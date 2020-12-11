using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day09 : ASolution
    {
        private readonly long[] numbers;

        public Day09() : base(09, 2020, "Encoding Error")
        {
            numbers = Input.SplitByNewline().Select(long.Parse).ToArray();
        }

        protected override string SolvePartOne()
        {
            int maxLookBack = 25;
            return numbers[Enumerable.Range(maxLookBack, numbers.Length).First(num => !CheckPrev(maxLookBack, num))].ToString();
        }

        private bool CheckPrev(int maxLookBack, int index)
        {
            var current = numbers[index];
            var sums = numbers[(index - maxLookBack)..index];

            for (int i = 0; i < sums.Length; i++)
            {
                for (int j = 0; j < sums.Length; j++)
                {
                    if (j == i) continue;

                    if (sums[i] + sums[j] == current)
                        return true;
                }
            }

            return false;
        }

        protected override string SolvePartTwo()
        {
            var toFind = 144381670;

            for(int i = 0; i < numbers.Length; i++)
            {
                long sum = 0;
                for(var j = i; j < numbers.Length; j++)
                {
                    sum += numbers[j];
                    if(sum == toFind)
                    {
                        var range = numbers[i..j];
                        return (range.Min() + range.Max()).ToString();
                    }
                    if (sum > toFind) break;
                }
            }

            return "ERROR";
        }
    }
}
