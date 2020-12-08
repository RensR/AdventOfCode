using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace AdventOfCode.Solutions.Year2020
{

    class Day08 : ASolution
    {
        Instruction[] baseInstructions;

        public Day08() : base(08, 2020, "")
        {
            baseInstructions = Input.SplitByNewline().ToList().Select(line => new Instruction(line)).ToArray();
        }

        protected override string SolvePartOne()
        {
            return RunProgram(baseInstructions).Item2.ToString();
        }

        protected override string SolvePartTwo()
        {
            for(int i = 0; i < baseInstructions.Length; i++)
            {
                Instruction[] newInstructions = Input.SplitByNewline().ToList().Select(line => new Instruction(line)).ToArray();
                if (baseInstructions[i].Id == "nop")
                {
                    newInstructions[i].Id = "jmp";
                }
                else if (baseInstructions[i].Id == "jmp")
                {
                    newInstructions[i].Id = "nop";
                }

                var results = RunProgram(newInstructions);
                if (results.Item1)
                {
                    return results.Item2.ToString();
                }
            }

            return "WRONG";
        }

        private (bool, int) RunProgram(Instruction[] instructions)
        {
            Dictionary<int, bool> SeenInstructons = new Dictionary<int, bool>();
            var accumulator = 0;
            var currentLine = 0;

            while (!SeenInstructons.ContainsKey(currentLine))
            {
                if (currentLine == instructions.Length)
                    return (true, accumulator);
                if (currentLine > instructions.Length)
                    return (false, accumulator);
                SeenInstructons[currentLine] = true;
                var currentInstruction = instructions[currentLine];
                switch (currentInstruction.Id)
                {
                    case "acc":
                        accumulator += currentInstruction.Value;
                        currentLine++;
                        break;
                    case "jmp":
                        currentLine += currentInstruction.Value;
                        break;
                    case "nop":
                        currentLine++;
                        break;
                }
            }

            return (false, accumulator);
        }
    }

    public class Instruction
    {
        public string Id;
        public int Value;

        public Instruction(string line)
        {
            Id = line[0..3];
            Value = int.Parse(line[4..]);
        }
    }
}
