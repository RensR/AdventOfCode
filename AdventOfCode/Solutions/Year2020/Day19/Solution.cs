using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day19 : ASolution
    {
        private readonly List<string> toTest;
        private readonly Dictionary<int, Match> rules = new();

        public Day19() : base(19, 2020, "Monster Messages")
        {
            var inputParts = Input.Split("\n\n");

            foreach (var stringRule in inputParts[0].SplitByNewline())
            {
                var idAndRule = stringRule.Split(':');
                var id = int.Parse(idAndRule[0]);
                rules[id] = idAndRule[1].Contains('"') ?
                    new CharMatch(idAndRule[1]) :
                    new DeepMatch(idAndRule[1]);
            }

            foreach (var key in rules.Keys)
            {
                rules[key].UpdateMatches(rules);
            }

            toTest = inputParts[1].SplitByNewline();
        }

        protected override string SolvePartOne()
        {
            return toTest.Count(MatchString).ToString();
        }

        private bool MatchString(string input)
        {
            return rules[0].MatchString(input).Contains(string.Empty);
        }

        protected override string SolvePartTwo()
        {
            rules[8] = new CyclicMatch(8, "42 | 42 8");
            rules[11] = new CyclicMatch(11, "42 31 | 42 11 31");

            rules[8].UpdateMatches(rules);
            rules[11].UpdateMatches(rules);
            rules[0].UpdateMatches(rules);

            var empty = toTest
                .Select(MatchString)
                .Count(match => match);

            return empty.ToString();
        }
    }

    abstract class Match
    {
        public abstract IEnumerable<string> MatchString(string input);
        public abstract void UpdateMatches(Dictionary<int, Match> matches);
    }

    class CharMatch : Match
    {
        public string MatchingString;

        public CharMatch(string s)
        {
            MatchingString = s.Trim()[1..^1];
        }

        public override IEnumerable<string> MatchString(string input)
        {
            if (input.StartsWith(MatchingString))
            {
                return new List<string> { input[MatchingString.Length..] };
            }

            return new List<string>();
        }

        public override void UpdateMatches(Dictionary<int, Match> matches)
        {
        }
    }

    class DeepMatch : Match
    {
        public IEnumerable<List<Match>> Options = new List<List<Match>>();
        public List<List<int>> OptionsInt = new();

        public DeepMatch(string line)
        {
            foreach (var option in line.Split('|'))
            {
                OptionsInt.Add(option.Trim().Split(' ').Select(int.Parse).ToList());
            }
        }

        public override IEnumerable<string> MatchString(string input)
        {
            return Options.SelectMany(option => MatchSingleString(input, option)).Distinct();
        }

        public static IEnumerable<string> MatchSingleString(string input, List<Match> matching)
        {
            var resultString = new List<string>();
            if (input.Length < matching.Count || matching.Count == 0)
                return resultString;

            // The string to find
            var intermediateResults = new List<string> { input };
            foreach (var match in matching)
            {
                foreach (var inputString in intermediateResults.Where(iResult => iResult != string.Empty).Distinct())
                {
                    resultString.AddRange(match.MatchString(inputString));
                }
                intermediateResults = new List<string>(resultString);
                resultString = new List<string>();
            }

            return intermediateResults;
        }

        public override void UpdateMatches(Dictionary<int, Match> matches)
        {
            Options = (from intOption in OptionsInt
                       select new List<Match>(intOption.Select(iop => matches[iop])));
        }
    }

    class CyclicMatch : DeepMatch
    {
        public CyclicMatch(int id, string line) : base(line)
        {
            OptionsInt = new List<List<int>>();

            var split = line.Split('|');
            var first = split[0].Trim();
            OptionsInt.Add(new List<int>(first.Split(' ').Select(int.Parse)));

            var second = split[1].Trim();
            var next = second;
            for (var i = 0; i < 10; i++)
            {
                var toAdd = next.Replace(id.ToString(), first).Trim();
                OptionsInt.Add(new List<int>(toAdd.Split(' ').Select(int.Parse)));

                next = next.Replace(id.ToString(), second).Trim();
            }
        }
    }
}