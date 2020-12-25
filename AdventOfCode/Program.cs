using System;
using System.Diagnostics;
using AdventOfCode.Solutions;

namespace AdventOfCode
{
    class Program
    {
        public static Config Config = Config.Get("config.json");
        private static readonly SolutionCollector Solutions = new(Config.Year, Config.Days);

        static void Main()
        {
            var total = new Stopwatch();
            total.Start();
            foreach (var solution in Solutions)
            {
                solution.Solve();
            }

            var ms = total.ElapsedMilliseconds;
            Console.WriteLine($"Total runtime: {ms} ms.");
        }
    }
}
