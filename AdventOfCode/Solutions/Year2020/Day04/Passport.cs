using System;
using System.Linq;
using System.Text.RegularExpressions;

namespace AdventOfCode.Solutions.Year2020
{
    class Passport
    {
        public string byr { get; set; } //(Birth Year)
        public string iyr { get; set; } //(Issue Year)
        public string eyr { get; set; } //(Expiration Year)
        public string hgt { get; set; } //(Height)
        public string hcl { get; set; } //(Hair Color)
        public string ecl { get; set; } //(Eye Color)
        public string pid { get; set; } //(Passport ID)
        public string cid { get; set; } //(Country ID)


        public Passport(string stringPassport)
        {
            var items = stringPassport.Split(' ');

            foreach (var item in items)
            {
                if (string.IsNullOrEmpty(item)) continue;
                var idAndValue = item.Split(':');
                var identifier = idAndValue[0];
                var value = idAndValue[1];
                switch (identifier)
                {
                    case "byr":
                        byr = value;
                        break;
                    case "iyr":
                        iyr = value;
                        break;
                    case "eyr":
                        eyr = value;
                        break;
                    case "hgt":
                        this.hgt = value;
                        break;
                    case "hcl":
                        this.hcl = value;
                        break;
                    case "ecl":
                        this.ecl = value;
                        break;
                    case "pid":
                        this.pid = value;
                        break;
                    case "cid":
                        this.cid = value;
                        break;
                }
            }
        }

        public Passport(string stringPassport, bool strict)
        {
            var items = stringPassport.Split(' ');

            foreach (var item in items)
            {
                if (string.IsNullOrEmpty(item)) continue;
                var idAndValue = item.Split(':');
                var identifier = idAndValue[0];
                var value = idAndValue[1];
                switch (identifier)
                {
                    case "byr":
                        if (value.All(char.IsDigit) && value.Length == 4)
                        {
                            var intValue = int.Parse(value);
                            if (intValue >= 1920 && intValue <= 2002)
                            {
                                byr = value;
                            }
                        }
                        break;
                    case "iyr":
                        if (value.All(char.IsDigit) && value.Length == 4)
                        {
                            var intValue = int.Parse(value);
                            if (intValue >= 2010 && intValue <= 2020)
                            {
                                iyr = value;
                            }
                        }
                        break;
                    case "eyr":
                        if (value.All(char.IsDigit) && value.Length == 4)
                        {
                            var intValue = int.Parse(value);
                            if (intValue >= 2020 && intValue <= 2030)
                            {
                                eyr = value;
                            }
                        }
                        break;
                    case "hgt":
                        if(value.Length > 2)
                        {
                            var unit = value.Substring(value.Length - 2, 2);
                            var meassurement = value[0..^2];
                            switch (unit)
                            {
                                case "cm":
                                    if (meassurement.All(char.IsDigit) && meassurement.Length == 3)
                                    {
                                        var intValue = int.Parse(meassurement);
                                        if(intValue >= 150 && intValue <= 193)
                                        {
                                            hgt = value;
                                        }
                                    }
                                    break;
                                case "in":
                                    if (meassurement.All(char.IsDigit) && meassurement.Length == 2)
                                    {
                                        var intValue = int.Parse(meassurement);
                                        if (intValue >= 59 && intValue <= 76)
                                        {
                                            hgt = value;
                                        }
                                    }
                                    break;
                            }
                        }
                        break;
                    case "hcl":
                        var rxhcl = new Regex("#[0-9a-f]{6}");
                        if (rxhcl.IsMatch(value))
                        {
                            hcl = value;
                        }
                        break;
                    case "ecl":
                        string[] rxecl = { "amb", "blu", "brn", "gry", "grn", "hzl", "oth" };
                        if (rxecl.Contains(value))
                        {
                            this.ecl = value;
                        }
                        break;
                    case "pid":
                        var rxpid = new Regex("[0-9]{9}");
                        if (rxpid.IsMatch(value) && value.Length == 9)
                        {
                            this.pid = value;
                        }
                        break;
                    case "cid":
                        this.cid = value;
                        break;
                }
            }
        }

        public bool IsValid()
        {
            return !string.IsNullOrEmpty(byr) &&
                !string.IsNullOrEmpty(iyr) &&
                !string.IsNullOrEmpty(eyr) &&
                !string.IsNullOrEmpty(hgt) &&
                !string.IsNullOrEmpty(hcl) &&
                !string.IsNullOrEmpty(ecl) &&
                !string.IsNullOrEmpty(pid) &&
                !string.IsNullOrEmpty(cid);

        }

        public bool IsFakeValid()
        {
            return !string.IsNullOrEmpty(byr) &&
              !string.IsNullOrEmpty(iyr) &&
              !string.IsNullOrEmpty(eyr) &&
              !string.IsNullOrEmpty(hgt) &&
              !string.IsNullOrEmpty(hcl) &&
              !string.IsNullOrEmpty(ecl) &&
              !string.IsNullOrEmpty(pid);
        }
    }
}
