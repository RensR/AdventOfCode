namespace AdventOfCode.Solutions.Year2020
{

    class Day03 : ASolution
    {
        bool[,] input;
        public Day03() : base(03, 2020, "Toboggan Trajectory")
        {
            var lines = Input.SplitByNewline();
            input = new bool[lines.Length, lines[0].Length];
            for (int y = 0; y < lines.Length; y++)
            {
                string line = lines[y];
                for (int x = 0; x < line.Length; x++)
                {
                    input[y,x] = line[x] == '#';
                }
            }
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

        private int GetTreesOnSlope(int deltaX, int deltaY)
        {
            var width = input.GetLength(1);
            var treeCount = 0;
            var x = 0;
            for (var y = 0; y < input.GetLength(0); y+= deltaY, x += deltaX)
            {
                if (x >= width) x -= width;
                if (input[y, x]) treeCount++;
            }

            return treeCount;
        }
    }
}
