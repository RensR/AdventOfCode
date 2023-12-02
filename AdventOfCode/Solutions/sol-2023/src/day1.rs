use std::collections::HashMap;

#[aoc(day1, part1, Improved)]
pub fn part1(input: &str) -> u32 {
    return input
        .lines()
        // Parse input into a vector of vectors of digits
        .map(|c| c.chars().filter_map(|a| a.to_digit(10)).collect::<Vec<_>>())
        // Reduce each vector of digits into a single number
        .map(|vec| 10 * vec.first().unwrap() + vec.last().unwrap())
        .sum();
}

#[aoc(day1, part2)]
pub fn part2(input: &str) -> u32 {
    return input.lines().map(|c| parse_string_digits(c)).sum();
}

fn parse_string_digits(t_num: &str) -> u32 {
    let mut number_lookup = HashMap::new();
    number_lookup.insert("one".to_string(), 1);
    number_lookup.insert("two".to_string(), 2);
    number_lookup.insert("three".to_string(), 3);
    number_lookup.insert("four".to_string(), 4);
    number_lookup.insert("five".to_string(), 5);
    number_lookup.insert("six".to_string(), 6);
    number_lookup.insert("seven".to_string(), 7);
    number_lookup.insert("eight".to_string(), 8);
    number_lookup.insert("nine".to_string(), 9);

    let chars = t_num.chars().collect::<Vec<_>>();
    let mut result: Vec<u32> = Vec::new();

    let mut i = 0;
    while i < chars.len() {
        if chars[i].to_digit(10) != None {
            result.push(chars[i].to_digit(10).unwrap());
        } else {
            for (key, value) in number_lookup.iter() {
                if i + key.len() > chars.len() {
                    continue;
                }
                if chars[i..i + key.len()].iter().collect::<String>() == *key {
                    result.push(*value);
                    break;
                }
            }
        }
        i = i + 1;
    }
    return 10 * result.first().unwrap() + result.last().unwrap();
}

#[cfg(test)]
mod tests {
    use super::{part1, part2};

    #[test]
    fn basics() {
        assert_eq!(part1("ninefourone1"), 11);
        assert_eq!(part2("ninefourone1"), 91);
    }

    #[test]
    fn annoying_strings() {
        assert_eq!(part2("sevenine"), 79);
    }
}
