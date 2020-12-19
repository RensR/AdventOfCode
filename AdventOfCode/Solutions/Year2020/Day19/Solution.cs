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
            return toTest.Where(MatchString).Count().ToString();
        }

        private bool MatchString(string input)
        {
            return rules[0].MatchString(input).Contains(string.Empty);
        }

        protected override string SolvePartTwo()
        {
            //if (rules[8] is DeepMatch rule8)
            //    rule8.Or = new List<Match> { rules[42], rule8 };


            //if (rules[1] is DeepMatch rule11)
            //    rule11.Or = new List<Match> { rules[42], rule11, rules[31] };

            return null;
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
            if (input == string.Empty || matching.Count == 0)
                return resultString;

            var intermediateResults = new List<string> { input };
            foreach(var match in matching)
            {
                foreach(var inputString in intermediateResults.Where(iresult => iresult != string.Empty))
                {
                    resultString.AddRange(match.MatchString(inputString));
                }
                intermediateResults = new List<string> (resultString );
            }

            return resultString;
        }

        public override void UpdateMatches(Dictionary<int, Match> matches)
        {
            foreach(var match in EitherInt)
            {
                Either.Add(matches[match]);
            }

            foreach (var match in OrInt)
            {
                Or.Add(matches[match]);
            }
        }
    }

}
