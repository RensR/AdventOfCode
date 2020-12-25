namespace AdventOfCode.Solutions.Year2020
{
    class Day25 : ASolution
    {
        private readonly long _door;
        private readonly long _keyCard;
        
        public Day25() : base(25, 2020, "Combo Breaker")
        {
            var lines = Input.SplitByNewline();
            _door = long.Parse(lines[0]);
            _keyCard = long.Parse(lines[1]);
        }

        protected override string SolvePartOne()
        {
            var val = 1;
            for (var i = 1;; i++)
            {
                val = val * 7 % 20201227;

                if (val == _door)
                {
                    return Loop(_keyCard, i).ToString();
                }

                if (val == _keyCard)
                {
                    return Loop(_door, i).ToString();
                }
            }
        }

        private static long Loop(long subjectNumber, long loopSize)
        {
            long number = 1;
            for (long i = 1; i <= loopSize; i++)
            {
                number = number * subjectNumber % 20201227;
            }

            return number;
        }
        
        protected override string SolvePartTwo()
        {
            return null;
        }
    }
}
