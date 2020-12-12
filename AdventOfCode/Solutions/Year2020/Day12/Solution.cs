using System;
using System.Collections.Generic;

namespace AdventOfCode.Solutions.Year2020
{
    class Day12 : ASolution
    {
        readonly List<MoveInstruction> instructions = new List<MoveInstruction>();

        public Day12() : base(12, 2020, "Rain Risk")
        {
            Input.SplitByNewline().ForEach(line => instructions.Add(new MoveInstruction(line)));
        }

        protected override string SolvePartOne()
        {
            int degrees = 90;
            var x = 0;
            var y = 0;

            foreach (var instruction in instructions)
            {
                switch (instruction.Action)
                {
                    case 'N':
                        y += instruction.Value;
                        break;
                    case 'S':
                        y -= instruction.Value;
                        break;
                    case 'E':
                        x += instruction.Value;
                        break;
                    case 'W':
                        x -= instruction.Value;
                        break;
                    case 'L':
                        degrees -= instruction.Value;
                        break;
                    case 'R':
                        degrees += instruction.Value;
                        break;
                    case 'F':
                        int deltaY, deltaX;
                        (deltaX, deltaY) = GetDirection(degrees);
                        x += deltaX * instruction.Value;
                        y += deltaY * instruction.Value;
                        break;
                }
            }

            return (Math.Abs(x) + Math.Abs(y)).ToString();
        }

        public static (int, int) GetDirection(int degrees)
        {
            var degreesNorm = ((degrees % 360) + 360) % 360;
            return degreesNorm switch
            {
                0 => (0, 1),
                90 => (1, 0),
                180 => (0, -1),
                270 => (-1, 0),
                _ => throw new ArgumentException(),
            };
        }

        protected override string SolvePartTwo()
        {
            var x = 0;
            var y = 0;
            var wayX = 10;
            var wayY = 1;            

            foreach (var instruction in instructions)
            {
                switch (instruction.Action)
                {
                    case 'N':
                        wayY += instruction.Value;
                        break;
                    case 'S':
                        wayY -= instruction.Value;
                        break;
                    case 'E':
                        wayX += instruction.Value;
                        break;
                    case 'W':
                        wayX -= instruction.Value;
                        break;
                    case 'L':
                        for (int r = instruction.Value; r > 0; r -= 90)
                            (wayX, wayY) = (-wayY, wayX);
                        break;
                    case 'R':
                        for (int r = instruction.Value; r > 0; r -= 90)
                            (wayX, wayY) = (wayY, -wayX);
                        break;
                    case 'F':
                        x += wayX * instruction.Value;
                        y += wayY * instruction.Value;
                        break;
                }
            }

            return (Math.Abs(x) + Math.Abs(y)).ToString();
        }
    }

    class MoveInstruction
    {
        public char Action;
        public int Value;

        public MoveInstruction(string line)
        {
            Action = line[0];
            Value = int.Parse(line[1..]);
        }
    }
}
