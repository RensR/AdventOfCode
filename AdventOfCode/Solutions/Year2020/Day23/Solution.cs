using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day23 : ASolution
    {
        public Day23() : base(23, 2020, "Crab Cups") { }
        
        protected override string SolvePartOne()
        {
            var dictionary = RunSim(9, 100);
            var theOne = dictionary[dictionary[1].Next];
            var result = string.Empty;
            while (theOne.Value is not 1)
            {
                result += theOne.Value.ToString();
                theOne = dictionary[theOne.Next];
            }

            return result;
        }
        
        protected override string SolvePartTwo()
        {
            var dictionary = RunSim(1_000_000, 10_000_000);
            var theOne = dictionary[1].Next;
            var theTwo = (long) dictionary[theOne].Next;
            return (theOne * theTwo).ToString();
        }

        private Dictionary<int, LinkedHashSetItem> RunSim(int maxInList, int runFor)
        {
            var dictionary = PrepareInput(maxInList);
            var current = dictionary[int.Parse(Input[..1])];

            for (var time = 0; time < runFor; time++)
            {
                var toRemove = current.Next;
                var instead = dictionary[dictionary[dictionary[toRemove].Next].Next].Next;
                current.Next = instead;

                for (var destination = current.Value - 1; ; destination--)
                {
                    if (destination == 0)
                        destination = maxInList;

                    if (destination == toRemove || destination == dictionary[toRemove].Next || destination == dictionary[dictionary[toRemove].Next].Next)
                        continue;

                    var destinationObject = dictionary[destination];
                    var currentNext = destinationObject.Next;
                    destinationObject.Next = toRemove;
                    dictionary[dictionary[dictionary[toRemove].Next].Next].Next = currentNext;

                    break;
                }

                current = dictionary[current.Next];
            }

            return dictionary;
        }

        private Dictionary<int, LinkedHashSetItem> PrepareInput(int items)
        {
            Dictionary<int, LinkedHashSetItem> dictionary = new();

            var input = Input.ToCharArray().Select(c => int.Parse(c.ToString())).ToList();

            for (var i = 0; i <= items; i++)
            {
                if (i < input.Count - 1)
                {
                    dictionary.Add(input[i], new LinkedHashSetItem( input[i + 1], input[i]));
                }
                else if (i == input.Count - 1)
                {
                    dictionary.Add(input[i],
                        items <= input.Count
                            ? new LinkedHashSetItem(input.First(), input[i])
                            : new LinkedHashSetItem(i + 2, input[i]));
                    i++;
                }
                else if (i < items)
                {
                    dictionary.Add(i, new LinkedHashSetItem(i + 1, i));
                }
                else if (i == items)
                {
                    dictionary.Add(i, new LinkedHashSetItem(input.First(), i));
                }
            }

            return dictionary;
        }

        private class LinkedHashSetItem
        {
            public int Next;
            public readonly int Value;

            public LinkedHashSetItem(int next, int value)
            {
                Next = next;
                Value = value;
            }

            public override string ToString()
            {
                return $"{Value} -> {Next}";
            }
        }
    }
}
