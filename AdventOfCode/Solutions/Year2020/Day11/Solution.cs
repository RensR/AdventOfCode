using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{

    class Day11 : ASolution
    {
        string[] seating;
        public Day11() : base(11, 2020, "Seating System")
        {
            seating = Input.SplitByNewline().ToArray();
        }

        protected override string SolvePartOne()
        {
            var newSeating = seating;
            var changed = true;
            while (changed)
            {
                (newSeating, changed) = UpdateSeating(newSeating);
            }
            return newSeating.Sum(s => s.Count(c => c == '#')).ToString();
        }

        private (string[], bool) UpdateSeating(string[] currentSeating)
        {
            var newSeating = new string[seating.Length];
            var changed = false;
            for(int i = 0; i < currentSeating.Length; i++)
            {
                var newLine = string.Empty;
                for (int j = 0; j < currentSeating[i].Length; j++)
                {
                    if (currentSeating[i][j] == '.')
                    {
                        newLine += '.';
                        continue;
                    }
                    var nextTo = 0;
                    for(var left = -1; left <= 1; left++)
                    {
                        if (i + left < 0 || i + left >= currentSeating.Length) continue;
                        for (var up = -1; up <= 1; up++)
                        {
                            if (left == 0 && up == 0) continue;
                            if (j + up < 0 || j + up >= currentSeating[i].Length) continue;

                            nextTo += currentSeating[i + left][j + up] == '#' ? 1 : 0;
                        }
                    }

                    if (currentSeating[i][j] == 'L')
                    {
                        newLine += nextTo == 0 ? '#' : 'L';
                    }
                    else if (currentSeating[i][j] == '#')
                    {
                        newLine += nextTo >= 4 ? 'L' : '#';
                    }
                }
                newSeating[i] = newLine;
                if (newSeating[i] != currentSeating[i])
                    changed = true;
            }

            return (newSeating, changed);
        }

        protected override string SolvePartTwo()
        {
            var newSeating = seating;
            var changed = true;
            while (changed)
            {
                (newSeating, changed) = UpdateSeatingB(newSeating);
            }
            return newSeating.Sum(s => s.Count(c => c == '#')).ToString();
        }

        private (string[], bool) UpdateSeatingB(string[] currentSeating)
        {
            var newSeating = new string[seating.Length];
            var changed = false;
            for (int i = 0; i < currentSeating.Length; i++)
            {
                var newLine = string.Empty;
                for (int j = 0; j < currentSeating[i].Length; j++)
                {
                    if (currentSeating[i][j] == '.')
                    {
                        newLine += '.';
                        continue;
                    }
                    var nextTo = 0;
                    for (var left = -1; left <= 1; left++)
                    {
                        if (i + left < 0 || i + left >= currentSeating.Length) continue;
                        for (var up = -1; up <= 1; up++)
                        {
                            if (left == 0 && up == 0) continue;
                            if (j + up < 0 || j + up >= currentSeating[i].Length) continue;

                            var leftleft = left;
                            var upup = up;
                            while(currentSeating[i + leftleft][j + upup] == '.')
                            {
                                leftleft += left;
                                upup += up;
                                if (i + leftleft < 0 || i + leftleft >= currentSeating.Length) break;
                                if (j + upup < 0 || j + upup >= currentSeating[i].Length) break;
                                
                            }
                            if (i + leftleft < 0 || i + leftleft >= currentSeating.Length) continue;
                            if (j + upup < 0 || j + upup >= currentSeating[i].Length) continue;
                            nextTo += currentSeating[i + leftleft][j + upup] == '#' ? 1 : 0;
                        }
                    }

                    if (currentSeating[i][j] == 'L')
                    {
                        newLine += nextTo == 0 ? '#' : 'L';
                    }
                    else if (currentSeating[i][j] == '#')
                    {
                        newLine += nextTo >= 5 ? 'L' : '#';
                    }
                }
                newSeating[i] = newLine;
                if (newSeating[i] != currentSeating[i])
                    changed = true;
            }

            return (newSeating, changed);
        }
    }
}
