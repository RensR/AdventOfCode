using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day19 : ASolution
    {
        List<string> toTest;
        Dictionary<int, Match> rules = new Dictionary<int, Match>();

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
            if (rules[0].MatchString(input, out var resultString))
                return resultString.Contains(string.Empty);
            return false;
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
        public abstract bool MatchString(string input, out List<string> resultString);
        public abstract void UpdateMatches(Dictionary<int, Match> matches);
    }

    class CharMatch : Match
    {
        public string matchingString;

        public CharMatch(string s)
        {
            matchingString = s[2..^1];
        }
        public override bool MatchString(string input, out List<string> resultString)
        {
            if (input.StartsWith(matchingString))
            {
                resultString = new List<string> { input[matchingString.Length..] };
                return true;
            }

            resultString = new List<string>();
            return false;
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

        List<int> EitherInt;
        List<int> OrInt = new List<int>();

        public DeepMatch(string line)
        {
            var eitherOr = line.Split('|');
            EitherInt = eitherOr[0].Trim().Split(' ').Select(int.Parse).ToList();
            if(eitherOr.Length > 1)
            {
                OrInt = eitherOr[1].Trim().Split(' ').Select(int.Parse).ToList();
            }
        }

        public override bool MatchString(string input, out List<string> resultString)
        {
            MatchSingleString(input, Either, out resultString);
            if(Or.Count > 0)
            {
                MatchSingleString(input, Or, out var moreResultStrings);
                resultString.AddRange(moreResultStrings);
            }
            if (resultString.Count > 0)
                return true;
            return false;
        }

        public bool MatchSingleString(string input, List<Match> matching, out List<string> resultString)
        {
            resultString = new List<string>();
            if (input == string.Empty)
                return false;

            var intermediateResults = new List<string> { input };
            foreach(var match in matching)
            {
                foreach(var inputString in intermediateResults.Where(iresult => iresult != string.Empty))
                {
                    match.MatchString(inputString, out var matchRemainder);
                    resultString.AddRange(matchRemainder);
                }
                intermediateResults = new List<string> (resultString );
            }

            return resultString.Count > 0;
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
