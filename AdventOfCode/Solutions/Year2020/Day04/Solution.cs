using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{

    class Day04 : ASolution
    {
        readonly List<string> passportLines;

        public Day04() : base(04, 2020, "Passport Processing")
        {
            passportLines = Input.Split("\n\n").Select(x => x.Replace("\n", " ")).ToList();
        }

        protected override string SolvePartOne()
        {
            var passports = passportLines.Select(line => new Passport(line));
            return passports.Count(passport => passport.IsFakeValid()).ToString();
        }

        protected override string SolvePartTwo()
        {
            var passports = passportLines.Select(line => new Passport(line, true));
            return passports.Count(passport => passport.IsFakeValid()).ToString();
        }
    }
}
