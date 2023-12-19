use scanf::sscanf;
use std::collections::HashMap;

struct Instruction {
    id: String,
    ops: Vec<(char, char, u64, String)>,
}

struct Part {
    x: u64,
    m: u64,
    a: u64,
    s: u64,
}

impl Part {
    fn new(x: u64, m: u64, a: u64, s: u64) -> Self {
        Self { x, m, a, s }
    }

    fn value(&self) -> u64 {
        return self.x + self.m + self.a + self.s;
    }
}

#[aoc(day19, part1)]
pub fn part1(input: &str) -> u64 {
    let (all_rules, parts) = input.split_once("\n\n").unwrap();
    let parsed_rules = all_rules
        .lines()
        .map(|rule_line| {
            let instruction = parse_rule_line(rule_line);
            return (instruction.id.to_string(), instruction);
        })
        .collect::<HashMap<_, _>>();

    return parts
        .lines()
        .map(|part_line| {
            let (mut x, mut m, mut a, mut s) = (0, 0, 0, 0);
            sscanf!(
                &part_line[1..part_line.len() - 1],
                "x={},m={},a={},s={}",
                x,
                m,
                a,
                s
            );
            return apply_instruction(&parsed_rules, &Part::new(x, m, a, s));
        })
        .sum();
}

fn apply_instruction(instructions: &HashMap<String, Instruction>, part: &Part) -> u64 {
    let mut cur = "in";
    loop {
        for (item, op, value, next) in instructions.get(cur).unwrap().ops.iter() {
            let value_to_check = get_attribute(*item, part);
            if checker(*op, value_to_check, *value) {
                if next == &"A" {
                    return part.value();
                }
                if next == &"R" {
                    return 0;
                }
                cur = next;
                break;
            }
        }
    }
}

fn checker(f: char, left: u64, right: u64) -> bool {
    return match f {
        '<' => left < right,
        '>' => left > right,
        ' ' => true,
        _ => panic!("Unknown operator"),
    };
}

fn get_attribute(f: char, part: &Part) -> u64 {
    return match f {
        'x' => part.x,
        'm' => part.m,
        'a' => part.a,
        's' => part.s,
        ' ' => 9999992222,
        _ => panic!("Unknown attribute {:?}", f),
    };
}

fn parse_rule_line(rule_line: &str) -> Instruction {
    let mut id: String = String::new();
    let mut rules: String = String::new();
    sscanf!(&rule_line, "{}{{{}}}", id, rules);

    let ops = rules
        .split(",")
        .map(|r| {
            if r.contains(":") {
                let (measurement, result) = r.split_once(":").unwrap();
                let item = measurement.chars().next().unwrap();
                let op = measurement.chars().nth(1).unwrap();
                let value = measurement[2..].parse::<u64>().unwrap();

                return (item, op, value, result.to_string());
            }

            return (' ', ' ', 0, r.to_string());
        })
        .collect::<Vec<_>>();

    return Instruction { id, ops };
}

#[aoc(day19, part2)]
pub fn part2(input: &str) -> u64 {
    let (all_rules, _) = input.split_once("\n\n").unwrap();
    let parsed_rules = all_rules
        .lines()
        .map(|rule_line| {
            let instruction = parse_rule_line(rule_line);
            return (instruction.id.to_string(), instruction);
        })
        .collect::<HashMap<_, _>>();

    let mut total = 0;
    let size = 40;
    for x in 0..size {
        for m in 0..size {
            for a in 0..size {
                for s in 0..size {
                    total += apply_instruction(&parsed_rules, &Part::new(x, m, a, s));
                }
            }
        }
    }
    return total;
}

#[cfg(test)]
mod tests {
    use super::part1;

    #[test]
    fn basics() {
        let input = "px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}";
        assert_eq!(part1(&input), 19114);
    }
}
