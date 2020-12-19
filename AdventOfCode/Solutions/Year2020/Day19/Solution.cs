using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day19 : ASolution
    {
        readonly List<string> toTest;
        readonly Dictionary<int, Match> rules = new Dictionary<int, Match>();

        public Day19() : base(19, 2020, "Monster Messages")
        {
            var inputParts = Input.Split("\n\n");

            foreach(var stringRule in inputParts[0].SplitByNewline())
            {
                var idAndRule = stringRule.Split(':');
                rules[int.Parse(idAndRule[0])] = idAndRule[1].Contains('"') ?
                    new CharMatch(idAndRule[1]) : 
                    new DeepMatch(idAndRule[1]);
            }

            foreach(int key in rules.Keys)
            {
                rules[key].UpdateMatches(rules);
            }

            toTest = inputParts[1].SplitByNewline();
        }

        protected override string SolvePartOne()
        {
            return null;
            return toTest.Where(MatchString).Count().ToString();
        }

        private bool MatchString(string input)
        {
            return rules[0].MatchString(input).Contains(string.Empty);
        }

        protected override string SolvePartTwo()
        {
            //return null;

            rules[8] = new CyclicMatch(" 42 | 42 8");
            rules[11] = new CyclicMatch(" 42 31 | 42 11 31");

            foreach (int key in rules.Keys)
            {
                rules[key].UpdateMatches(rules);
            }

            return toTest.Where(MatchString).Count().ToString();
        }
    }

    abstract class Match
    {
        public abstract List<string> MatchString(string input);
        public abstract void UpdateMatches(Dictionary<int, Match> matches);
    }

    class CharMatch : Match
    {
        public string matchingString;

        public CharMatch(string s)
        {
            matchingString = s[2..^1];
        }
        public override List<string> MatchString(string input)
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
        public List<Match> Either = new List<Match>();
        public List<Match> Or = new List<Match>();
        readonly List<int> EitherInt;
        readonly List<int> OrInt = new List<int>();

        public DeepMatch(string line)
        {
            var eitherOr = line.Split('|');
            EitherInt = eitherOr[0].Trim().Split(' ').Select(int.Parse).ToList();
            if(eitherOr.Length > 1)
            {
                OrInt = eitherOr[1].Trim().Split(' ').Select(int.Parse).ToList();
            }
        }

        public override List<string> MatchString(string input)
        {
            var resultString = MatchSingleString(input, Either);
            resultString.AddRange(MatchSingleString(input, Or));
            return resultString;
        }

        public static List<string> MatchSingleString(string input, List<Match> matching)
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
            Either = new List<Match>();
            Or = new List<Match>();
            foreach (var match in EitherInt)
            {
                Either.Add(matches[match]);
            }

            foreach (var match in OrInt)
            {
                Or.Add(matches[match]);
            }
        }
    }

    class CyclicMatch : DeepMatch
    {
        public CyclicMatch(string line) : base(line)
        {
            
        }

        public override void UpdateMatches(Dictionary<int, Match> matches)
        {
            base.UpdateMatches(matches);
        }
    }
}
