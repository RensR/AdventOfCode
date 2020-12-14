using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day13 : ASolution
    {
        readonly int depart;
        readonly List<string> busses;

        public Day13() : base(13, 2020, "Shuttle Search")
        {
            var result = Input.SplitByNewline();
            depart = int.Parse(result[0]);
            busses = result[1].Split(',').ToList();
        }

        protected override string SolvePartOne()
        {
            var bus = busses
                .Where(v => v != "x")
                .Select(v => int.Parse(v))
                .Select(bus => (bus - (depart % bus), bus))
                .Min();

            return (bus.Item1 * bus.bus).ToString();
        }

        protected override string SolvePartTwo()
        {
            var relevantBusses = new List<(long, long)>();

            for (int i = 0; i < busses.Count; i++)
            {
                if (busses[i] == "x") continue;
                relevantBusses.Add((long.Parse(busses[i]), i));
            }

            // The aggregate funtion would be nice to use but as it folds over the list
            // the last operations will merge bussed with a very large and a very small
            // loop time. This is incredibly inefficient, therefore we try to always 
            // merge busses with similar round trip times. This is also why we sort the 
            // list in the loop.
            // var bus = relevantBusses.Aggregate(MergeBusses); == slow

            while (relevantBusses.Count > 1)
            {
                var merged = new List<(long, long)>();

                for (int i = 0; i + 1 < relevantBusses.Count; i += 2)
                {
                    merged.Add(MergeBusses(relevantBusses[i], relevantBusses[i + 1]));
                }
                if (relevantBusses.Count % 2 != 0)
                    merged.Add(relevantBusses.Last());

                merged.Sort();
                relevantBusses = merged;
            }

          
            return (relevantBusses[0].Item1 - relevantBusses[0].Item2).ToString();
        }

        static (long, long) MergeBusses((long, long) a, (long, long) b)
        {
            (long timeA, long offsetA) = a;
            (long timeB, long offsetB) = b;
            for (long i = 1; ; i++)
            {
                if (timeA * i  % timeB  == 0)
                {
                    for (long la = 0, lb = 0; ;)
                    {
                        if(offsetA + la * timeA == offsetB + lb * timeB)
                        {
                            return (timeA * i, offsetA + la * timeA);
                        }

                        if (offsetA + la * timeA < offsetB + lb * timeB)
                            la++;
                        else
                            lb++;
                    }
                }
            }
        }
    }
}
