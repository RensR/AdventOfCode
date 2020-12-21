using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day01 : ASolution
    {
        private readonly int[] _expenses;

        public Day01() : base(01, 2020, "Report Repair")
        {
            _expenses = Input.ToIntArray("\n").Where(e => e < 2020).ToArray();
        }

        protected override string SolvePartOne()
        {
            return (
                from a in _expenses 
                from b in _expenses 
                where a + b == 2020 
                select (a * b).ToString()).FirstOrDefault();
        }

        protected override string SolvePartTwo()
        {
            return (
                from a in _expenses 
                from b in _expenses 
                from c in _expenses 
                where a + b + c == 2020 
                select (a * b * c).ToString()).FirstOrDefault();
        }
    }
}
