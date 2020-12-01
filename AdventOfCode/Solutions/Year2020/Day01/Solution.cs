using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{

    class Day01 : ASolution
    {
        readonly int[] _expenses;

        public Day01() : base(01, 2020, "Report Repair")
        {
            _expenses = Input.ToIntArray("\n").Where(e => e < 2020).ToArray();
        }

        protected override string SolvePartOne()
        {
            foreach (var a in _expenses)
            {
                foreach (var b in _expenses)
                {
                    if (a + b == 2020) return (a * b).ToString();
                }
            }
            return null;
        }

        protected override string SolvePartTwo()
        {
            foreach (var a in _expenses)
            {
                foreach (var b in _expenses)
                {
                    foreach (var c in _expenses)
                    {
                        if (a + b + c == 2020) return (a * b * c).ToString();
                    }
                }
            }
            return null;
        }
    }
}
