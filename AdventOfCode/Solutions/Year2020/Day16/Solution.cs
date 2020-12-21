using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day16 : ASolution
    {
        List<Rule> Rules;
        readonly List<int> MyTicket;
        readonly List<List<int>> OtherTickets;

        public Day16() : base(16, 2020, "Ticket Translation")
        {
            var segments = Input.Split("\n\n");

            Rules = segments[0].SplitByNewline().Select(line => new Rule(line)).ToList();
            MyTicket = segments[1].SplitByNewline()[1].Split(',').Select(int.Parse).ToList();
            OtherTickets = segments[2].SplitByNewline().ToArray()[1..].Select(line => line.Split(',').Select(int.Parse).ToList()).ToList();
        }

        protected override string SolvePartOne()
        {
            return (from ticket in OtherTickets
                    from val in ticket
                    where !Validate(val)
                    select val).Sum().ToString();
        }

        bool Validate(int value)
        {
            foreach(var rule in Rules)
            {
                if ((value >= rule.LowerA && value <= rule.UpperA)  || (value >= rule.LowerB && value <= rule.UpperB))
                    return true;
            }
            return false;
        }

        protected override string SolvePartTwo()
        {
            var goodTickets = new List<List<int>>();
            foreach (var ticket in OtherTickets)
            {
                var valid = true;
                foreach (var val in ticket)
                {
                    if (!Validate(val))
                        valid = false;
                }
                if(valid)
                    goodTickets.Add(ticket);
            }

            foreach (Rule rule in Rules)
            {
                for (int i = 0; i < MyTicket.Count; i++)
                {
                    if (goodTickets.All(ticket => (ticket[i] >= rule.LowerA && ticket[i] <= rule.UpperA) || (ticket[i] >= rule.LowerB && ticket[i] <= rule.UpperB)))
                    {
                        rule.Index.Add(i);
                    }
                }
            }

            Rules = RecudeLists(Rules);

            var currentRules = Rules.Where(rule => rule.Name.StartsWith("departure"));

            long returnI = 1;
            foreach (Rule rule in currentRules)
            {
                returnI *= MyTicket[rule.FinalIndex]; 
            }

            return returnI.ToString();
        }

        public List<Rule> RecudeLists(List<Rule> rules)
        {
            var done = new List<Rule>();

            while (done.Count < MyTicket.Count - 1)
            {
                foreach (Rule rule in rules)
                {
                    if (rule.Index.Count == 1)
                    {
                        rule.FinalIndex = rule.Index.First();
                        done.Add(rule);

                        foreach (Rule otherRule in rules)
                        {
                            if (otherRule != rule)
                                otherRule.Index.Remove(rule.Index.First());
                        }
                    }
                }
                done.ForEach(d => rules.Remove(d));
            }
            return done;
        }
    }

    class Rule
    {
        public string Name;

        public int LowerA;
        public int LowerB;
        public int UpperA;
        public int UpperB;

        public List<int> Index = new();
        public int FinalIndex;

        public Rule(string line)
        {
            var parts = line.Split(':');

            Name = parts[0];
            var rules = parts[1].Split(" or ");
            var ruleA = rules[0].Split('-');
            var ruleB = rules[1].Split('-');

            LowerA = int.Parse(ruleA[0]);
            LowerB = int.Parse(ruleB[0]);
            UpperA = int.Parse(ruleA[1]);
            UpperB = int.Parse(ruleB[1]);
        }
    }
}
