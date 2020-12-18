using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

namespace AdventOfCode.Solutions.Year2020
{

    class Day18 : ASolution
    {
        List<string> lines;
        public Day18() : base(18, 2020, "")
        {
            lines = Input.SplitByNewline();
        }

        protected override string SolvePartOne()
        {
            return null;
            return lines.Select(line => Item.Eval(Item.Parse(line)).Value).Sum().ToString();
        }


        protected override string SolvePartTwo()
        {
            return lines.Select(line => Item.Eval(Item.Parse(line, true)).Value).Sum().ToString();
        }
    }

    public class Item
    {
        public long Value;

        public static Item Parse(string line, bool instantTimes = false)
        {
            line = RemoveBrackets(line, Parse, instantTimes);
            Operation product = null;
            var values = line.Split(' ');

            for (int i = 1; i < values.Length; i += 2)
            {
                if (product == null)
                {
                    product = new Operation(
                        new Val(values[i - 1]),
                        new Val(values[i + 1]),
                        values[i]);
                }
                else
                {
                    if(instantTimes && values[i] == "+")
                    {
                        Operation deepRight = product;

                        while(deepRight.Right is Operation op)
                        {
                            deepRight = op;
                        }
                        deepRight.Right = new Operation(
                            new Val(values[i - 1]),
                            new Val(values[i + 1]),
                            values[i]);

                    }
                    else
                        product = new Operation(
                            product,
                            new Val(values[i + 1]),
                            values[i]);
                }  
            }

            return product;
        }

        private static string RemoveBrackets(string line, Func<string, bool, Item> parse, bool instantTimes)
        {
            var withinBrackets = Regex.Match(line, @"\([^\(\)]*\)");
            while (withinBrackets.Success)
            {
                line = line.Replace(withinBrackets.Value, Eval(parse(withinBrackets.Value[1..^1], instantTimes)).Value.ToString());
                withinBrackets = Regex.Match(line, @"\([^\(\)]*\)");
            }
            return line;
        }

        public static Item Eval(Operation op)
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

        public static Item Eval(Item val)
        {
            if (val is Operation op)
                return Eval(op);
            else return val;
        }
    }

    public class Operation : Item
    {
        public Item Left;
        public Item Right;
        public Func<Item, Item, Item> Func;
        public string Funcc;

        public Operation(Item left, Item right, string func)
        {
            Left = left;
            Right = right;
            Func = func == "+" ? Plus : Times;
            Funcc = func;
        }

        public static Func<Item, Item, Item> Times = (left, right) => new Val(Eval(left).Value * Eval(right).Value);
        public static Func<Item, Item, Item> Plus = (left, right) => new Val(Eval(left).Value + Eval(right).Value);

    }

    public class Val : Item
    {
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
