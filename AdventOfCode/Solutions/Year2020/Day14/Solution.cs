using System;
using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day14 : ASolution
    {
        List<string> lines;
        public Day14() : base(14, 2020, "Docking Data")
        {
            lines = Input.SplitByNewline();
        }

        protected override string SolvePartOne()
        {
            var valDict = new Dictionary<long, long>();

            var mask = string.Empty;
            foreach (var line in lines)
            {
                if(line.StartsWith("mask = "))
                {
                    mask = line[7..];
                    continue;
                }

                var instructionAndValue = line.Split(" = ");
                var bitValue = Convert.ToString(long.Parse(instructionAndValue[1]), 2).PadLeft(36).Replace(' ', '0').ToCharArray();

                for(var i = 0; i < mask.Length; i++)
                {
                    if (mask[i] == 'X') continue;
                    bitValue[i] = mask[i];
                }

                var location = long.Parse(instructionAndValue[0][4..^1]);
                var value = Convert.ToInt64(new string(bitValue), 2);

                valDict[location] = value;
            }


            return valDict.Values.Sum().ToString();
        }

        readonly Dictionary<long, long> valDict2 = new();

        protected override string SolvePartTwo()
        {
            var mask = string.Empty;
            foreach (var line in lines)
            {
                if (line.StartsWith("mask = "))
                {
                    mask = line[7..];
                    continue;
                }

                var instructionAndValue = line.Split(" = ");
                var bitLocations = Convert.ToString(long.Parse(instructionAndValue[0][4..^1]), 2).PadLeft(36).Replace(' ', '0').ToCharArray();

                for (var i = 0; i < mask.Length; i++)
                {
                    if (mask[i] == '0') continue;
                    bitLocations[i] = mask[i];
                }

                UpdateBit(bitLocations, long.Parse(instructionAndValue[1]));
            }

            return valDict2.Values.Sum().ToString();
        }

        private void UpdateBit(char[] bitLocations, long value)
        {
            var indexX = Array.FindIndex(bitLocations, c => c == 'X');
            if (indexX != -1)
            {
                bitLocations[indexX] = '1';
                UpdateBit(bitLocations, value);
                bitLocations[indexX] = '0';
                UpdateBit(bitLocations, value);
                bitLocations[indexX] = 'X';
            }
            else
            {
                var location = Convert.ToInt64(new string(bitLocations), 2);
                valDict2[location] = value;
            }
        }
    }
}
