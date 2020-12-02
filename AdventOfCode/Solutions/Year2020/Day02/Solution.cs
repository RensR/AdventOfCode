using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{

    class Day02 : ASolution
    {
        private readonly string[] _lines;
        
        public Day02() : base(02, 2020, "")
        {
            _lines = Input.Split('\n');
        }

        protected override string SolvePartOne()
        {
            return _lines.Select(IsValidPassword).Sum().ToString();
        }

        private int IsValidPassword(string rulesAndPass)
        {
            var parts = rulesAndPass.Split(' ');
            var minMax = parts[0].Split('-').Select(int.Parse).ToArray();
            var chara = parts[1].First();
            var count = parts[2].Count(c => c == chara);

            return count >= minMax[0] && count <= minMax[1] ? 1 : 0;
        }

        protected override string SolvePartTwo()
        {
            return _lines.Select(IsValidPassword2).Sum().ToString();
        }

        private static int IsValidPassword2(string rulesAndPass)
        {
            var parts = rulesAndPass.Split(' ');
            var minMax = parts[0].Split('-').Select(int.Parse).ToArray();
            var chara = parts[1].First();


            if (minMax[0] <= parts[2].Length && parts[2][minMax[0] - 1] == chara)
            {
                if (minMax[1] > parts[2].Length || parts[2][minMax[1] - 1] != chara)
                {
                    return 1;
                }

                return 0;
            }

            return minMax[1] <= parts[2].Length && parts[2][minMax[1] - 1] == chara ? 1 : 0;
        }
    }
}
