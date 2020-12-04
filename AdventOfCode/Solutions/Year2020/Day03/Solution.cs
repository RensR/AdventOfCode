using System.Collections.Generic;

namespace AdventOfCode.Solutions.Year2020
{

    class Day03 : ASolution
    {
        List<string> lines;
        public Day03() : base(03, 2020, "Toboggan Trajectory")
        {
            lines = Input.SplitByNewline();
        }

        protected override string SolvePartOne()
        {
            return GetTreesOnSlope(3, 1).ToString();
        }

        protected override string SolvePartTwo()
        {
            return 
                (GetTreesOnSlope(1, 1) *
                GetTreesOnSlope(3, 1) *
                GetTreesOnSlope(5, 1) *
                GetTreesOnSlope(7, 1) *
                GetTreesOnSlope(1, 2)).ToString();
        }

        private long GetTreesOnSlope(int deltaX, int deltaY)
        {
            var treeCount = 0;
            var x = 0;
            for (var y = 0; y < lines.Count; y += deltaY, x += deltaX)
            {
                var width = lines[y].Length;
                if (lines[y][x % width] == '#') treeCount++;
            }

            return treeCount;
        }
    }
}
