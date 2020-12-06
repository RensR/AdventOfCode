using System;
using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day06 : ASolution
    {
        readonly List<string> lines;

        public Day06() : base(06, 2020, "Custom Customs")
        {
            lines = Input.Split("\n\n").ToList();
        }

        protected override string SolvePartOne()
        {
            return lines.Select(line => line.Replace("\n", "").Distinct().Count()).Sum().ToString();
        }

        protected override string SolvePartTwo()
        {
            return lines.Select(line =>
            {
                return line
                .Split('\n')
                .Select(q => q.Distinct())
                .Aggregate<IEnumerable<char>, IEnumerable<char>>( new List<char> { 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'z', 'x', 'c', 'v', 'b', 'n', 'm' },
                    (a, b) => a.Intersect(b))
                .Count();
            }).Sum().ToString();
        }

        // Made own intersect, replaced by built in Intersect
        private static IEnumerable<char> GetMatching(IEnumerable<char> first, IEnumerable<char> second)
        {
            foreach(char f in first)
            {
                if (second.Contains(f))
                    yield return f;
            }
        }
    }
}
