using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

namespace AdventOfCode.Solutions.Year2020
{
    class Day18 : ASolution
    {
        readonly List<string> lines;
        public Day18() : base(18, 2020, "Operation Order")
        {
            lines = Input.SplitByNewline();
        }

        protected override string SolvePartOne()
        {
            return lines.Select(line => Item.Eval(Item.Parse(line)).Value).Sum().ToString();
        }

        protected override string SolvePartTwo()
        {
            return lines.Select(line => Item.Eval(Item.Parse(line, true)).Value).Sum().ToString();
        }
    }

    public class Item
    {
        public static Item Parse(string line, bool instantPlus = false)
        {
            var values = RemoveBrackets(line, instantPlus).Split(' ');
            Operation product = null;

            for (int i = 1; i < values.Length; i += 2)
            {
                if (product == null)
                {
                    product = new Operation(new Val(values[i - 1]), new Val(values[i + 1]), values[i]);
                }
                else
                {
                    if(instantPlus && values[i] == "+")
                    {
                        Operation deepRight = product;

                        while(deepRight.Right is Operation op)
                            deepRight = op;

                        deepRight.Right = new Operation(
                            new Val(values[i - 1]),
                            new Val(values[i + 1]),
                            values[i]);
                    }
                    else
                    {
                        product = new Operation(product, new Val(values[i + 1]), values[i]);
                    }
                }  
            }

            return product;
        }

        private static string RemoveBrackets(string line, bool instantPlus)
        {
            var withinBrackets = Regex.Match(line, @"\([^\(\)]*\)");
            while (withinBrackets.Success)
            {
                line = line.Replace(withinBrackets.Value, Eval(Parse(withinBrackets.Value[1..^1], instantPlus)).Value.ToString());
                withinBrackets = Regex.Match(line, @"\([^\(\)]*\)");
            }
            return line;
        }

        public static Val Eval(Operation op)
        {
            while (op.Left is Operation opL)
            {
                op.Left = Eval(opL);
            }

            while (op.Right is Operation opR)
            {
                op.Right = Eval(opR);
            }

            return op.Func(op.Left, op.Right);
        }

        public static Val Eval(Item val)
        {
            if (val is Operation op)
                return Eval(op);
            else return val as Val;
        }
    }

    public class Operation : Item
    {
        public Item Left;
        public Item Right;
        public Func<Item, Item, Val> Func;
        public string Funcc;

        public static Func<Item, Item, Val> Times = (left, right) => new Val(Eval(left).Value * Eval(right).Value);
        public static Func<Item, Item, Val> Plus = (left, right) => new Val(Eval(left).Value + Eval(right).Value);

        public Operation(Item left, Item right, string func)
        {
            Left = left;
            Right = right;
            Func = func == "+" ? Plus : Times;
            Funcc = func;
        }
    }

    public class Val : Item
    {
        public long Value;

        public Val(long val)
        {
            Value = val;
        }

        public Val(string val)
        {
            Value = long.Parse(val);
        }
    }
}
