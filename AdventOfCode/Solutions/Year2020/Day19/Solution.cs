using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace AdventOfCode.Solutions.Year2020
{
    class Day19 : ASolution
    {
        private readonly List<string> toTest;
        readonly Dictionary<int, Match> rules = new Dictionary<int, Match>();

        public Day19() : base(19, 2020, "Monster Messages")
        {
            var inputParts = Input.Split("\n\n");

            foreach(var stringRule in inputParts[0].SplitByNewline())
            {
                var idAndRule = stringRule.Split(':');
                rules[int.Parse(idAndRule[0])] = idAndRule[1].Contains('"') ?
                    new CharMatch(idAndRule[1]) : 
                    new DeepMatch(stringRule);
            }

            foreach(int key in rules.Keys)
            {
                rules[key].UpdateMatches(rules);
            }

            toTest = inputParts[1].SplitByNewline();
        }

        protected override string SolvePartOne()
        {
            //return null;
            return toTest.Where(s => MatchString(s, rules)).Count().ToString();
        }

        private static bool MatchString(string input, Dictionary<int, Match> rules)
        {
            return rules[0].MatchString(input).Contains(string.Empty);
        }

        private static IEnumerable<string> MatchString2(string input, Dictionary<int, Match> rules)
        {
            return rules[0].MatchString(input);
        }

        protected override string SolvePartTwo()
        {
            //return null;

            rules[8] = new CyclicMatch("8: 42 | 42 8");
            rules[11] = new CyclicMatch("11: 42 31 | 42 11 31");

            foreach (int key in rules.Keys)
            {
                rules[key].UpdateMatches(rules);
            }

            var result = new List<IEnumerable<string>>();

            toTest.ForEach(current =>
            {
                result.Add(MatchString2(current, rules));
            });

            //Parallel.ForEach(toTest, (current) =>
            //{
            //    result.Add(MatchString2(current, rules));
            //});

            var empty = result.Where(r => r.Contains(string.Empty)).ToList();

            return empty.Count().ToString();
        }
    }

    abstract class Match
    {
        public abstract IEnumerable<string> MatchString(string input);
        public abstract void UpdateMatches(Dictionary<int, Match> matches);
    }

    class CharMatch : Match
    {
        public string matchingString;

        public CharMatch(string s)
        {
            matchingString = s[2..^1];
        }

        public override IEnumerable<string> MatchString(string input)
        {
            if (input.StartsWith(matchingString))
            {
                return new List<string> { input[matchingString.Length..] };
            }

            return new List<string>();
        }

        public override void UpdateMatches(Dictionary<int, Match> matches)
        {
            return;
        }
    }

    class DeepMatch : Match
    {
        public int Id;
        public List<List<Match>> Options = new List<List<Match>>();
        public List<List<int>> OptionsInt = new List<List<int>>();

        public DeepMatch(string line)
        {
            var idAndRule = line.Split(':');
            Id = int.Parse(idAndRule[0]);
            foreach (var option in idAndRule[1].Split('|'))
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

            var intermediateResults = new List<string> { input };
            foreach(var match in matching)
            {
                foreach(var inputString in intermediateResults.Where(iresult => iresult != string.Empty).Distinct())
                {
                    resultString.AddRange(match.MatchString(inputString));
                }
                intermediateResults = new List<string> (resultString );
            }

            return resultString;
        }

        public override void UpdateMatches(Dictionary<int, Match> matches)
        {
            Options = (from intOption in OptionsInt
                       select new List<Match>(intOption.Select(iop => matches[iop]))).ToList();
        }
    }

    class CyclicMatch : DeepMatch
    {
        public CyclicMatch(string line) : base(line)
        {
            var idAndLine = line.Split(':');
            OptionsInt = new List<List<int>>();

            var split = idAndLine[1].Split('|');
            var first = split[0].Trim();
            OptionsInt.Add(new List<int>(first.Split(' ').Select(int.Parse)));

            var second = split[1].Trim();
            var next = second;
            for(int i = 0; i < 25; i++)
            {
                var toAdd = next.Replace(idAndLine[0], first).Trim();
                OptionsInt.Add(new List<int>(toAdd.Split(' ').Select(int.Parse)));

                next = next.Replace(idAndLine[0], second).Trim();
            }
        }
    }
}
