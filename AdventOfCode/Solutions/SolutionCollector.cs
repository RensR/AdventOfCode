using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions
{

    class SolutionCollector : IEnumerable<ASolution>
    {
        readonly IEnumerable<ASolution> _solutions;

        public SolutionCollector(int year, int[] days) => _solutions = LoadSolutions(year, days).ToArray();

        public ASolution GetSolution(int day)
        {
            try
            {
                return _solutions.Single(s => s.Day == day);
            }
            catch(InvalidOperationException)
            {
                return null;
            }
        }

        public IEnumerator<ASolution> GetEnumerator()
        {
            return _solutions.GetEnumerator();
        }

        IEnumerator IEnumerable.GetEnumerator()
        {
            return GetEnumerator();
        }

        static IEnumerable<ASolution> LoadSolutions(int year, int[] days)
        {
            if(days.Sum() == 0)
            {
                days = Enumerable.Range(1, 25).ToArray();
            }
            
            foreach(int day in days)
            {
                var solution = Type.GetType($"AdventOfCode.Solutions.Year{year}.Day{day:D2}");
                if(solution != null)
                {
                    yield return (ASolution)Activator.CreateInstance(solution);
                }
            }
        }
    }
}
