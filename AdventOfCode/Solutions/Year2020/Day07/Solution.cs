using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

namespace AdventOfCode.Solutions.Year2020
{
    class Day07 : ASolution
    {
        private readonly Dictionary<string, List<string>> isContainedIn = new();
        private readonly Dictionary<string, List<string>> contains = new();
        readonly Dictionary<string, bool> seen = new();
        int count = -1;

        public Day07() : base(07, 2020, "Handy Haversacks")
        {
            var lines = Input.Replace("bags", "bag").Replace(" ", "").Replace(".", "").SplitByNewline();

            foreach(string line in lines)
            {
                var segments = line.Split("contain");
                var innerBags = segments[1].Split(",");
                var containingBag = segments[0];

                foreach (var innerBag in innerBags)
                {
                    // part 1
                    var cleanInnerBag = Regex.Replace(innerBag, @"\d*", "");

                    if (isContainedIn.ContainsKey(cleanInnerBag)) 
                    {
                        isContainedIn[cleanInnerBag].Add(containingBag);
                    }
                    else
                    {
                        isContainedIn.Add(cleanInnerBag, new List<string> { containingBag });
                    }

                    // part 2
                    if (innerBag == "nootherbag") continue;
                    var numberOfBags = int.Parse(innerBag[0].ToString());
                    var bagName = innerBag[1..];

                    if (contains.ContainsKey(containingBag))
                    {
                        contains[containingBag].AddRange(Enumerable.Repeat(cleanInnerBag, numberOfBags));
                    }
                    else
                    {
                        contains.Add(containingBag, Enumerable.Repeat(cleanInnerBag, numberOfBags).ToList());
                    }
                }
            }
        }

        protected override string SolvePartOne()
        {
            CheckNextBag("shinygoldbag");
            return count.ToString();
        }

        private void CheckNextBag(string current)
        {
            if (seen.ContainsKey(current))
                return;
            seen.Add(current, true);
            count++;
            if (!isContainedIn.ContainsKey(current)) 
                return;
            var canBeContainedIn = isContainedIn[current];
            foreach(var bag in canBeContainedIn)
            {
                CheckNextBag(bag);
            }
        }

        protected override string SolvePartTwo()
        {
            return (CheckBag("shinygoldbag") - 1).ToString();
        }

        private int CheckBag(string current)
        {
            if (!contains.ContainsKey(current))
                return 1;

            return 1 + contains[current].Sum(CheckBag);
        }
    }
}
