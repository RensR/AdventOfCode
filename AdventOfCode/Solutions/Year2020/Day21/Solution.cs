using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

namespace AdventOfCode.Solutions.Year2020
{
    class Day21 : ASolution
    {
        readonly Dictionary<string, Allergen> Allergens = new();
        readonly List<string> FoodList = new();

        public Day21() : base(21, 2020, "Allergen Assessment")
        {
            foreach (var line in Input.SplitByNewline())
            {
                var allergenList = Regex.Match(line, @"\(contains (.+)\)")
                    .Groups[1].Value.Trim()
                    .Split(',')
                    .Select(al => al.Trim());

                var foodList = new string(line.TakeWhile(c => c != '(').ToArray()).Trim().Split(' ')
                    .Select(al => al.Trim()).ToList();
                FoodList.AddRange(foodList);
                foreach (var allergenInformation in allergenList)
                {
                    if (!Allergens.ContainsKey(allergenInformation))
                    {
                        Allergens[allergenInformation] = new Allergen(allergenInformation);
                    }

                    Allergens[allergenInformation].FoundIn.Add(foodList);
                }
            }

            foreach (var allergen in Allergens.Values)
            {
                allergen.ShortList.AddRange(from ingredient in allergen.FoundIn[0]
                    where allergen.FoundIn.All(otherP => otherP.Contains(ingredient))
                    select ingredient);
            }

            var done = new List<string>();
            var singleOne = Allergens.Values.FirstOrDefault(allergen => allergen.ShortList.Count == 1);
            while (singleOne != null)
            {
                foreach (var others in Allergens.Values)
                {
                    if (others.Name == singleOne.Name) continue;
                    others.ShortList.Remove(singleOne.ShortList[0]);
                }

                done.Add(singleOne.Name);
                singleOne = Allergens.Values.FirstOrDefault(allergen =>
                    allergen.ShortList.Count == 1 && !done.Contains(allergen.Name));
            }
        }

        protected override string SolvePartOne()
        {
            var realAllergens = Allergens.Values.SelectMany(allergen => allergen.ShortList);
            var nonAllergen = (from food in FoodList
                where !realAllergens.Contains(food)
                select food).Count();
            return nonAllergen.ToString();
        }

        protected override string SolvePartTwo()
        {
            var sorted = Allergens.Values.OrderBy(all => all.Name);

            return string.Join(',', sorted.Select(all => all.ShortList[0]));
        }
    }

    class Allergen
    {
        public List<List<string>> FoundIn = new List<List<string>>();
        public string Name;
        public List<string> ShortList = new List<string>();

        public Allergen(string name)
        {
            Name = name;
        }
    }
}
