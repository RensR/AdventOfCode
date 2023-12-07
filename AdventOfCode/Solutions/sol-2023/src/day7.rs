use std::collections::HashMap;

#[aoc_generator(day7)]
pub fn input_generator(input: &str) -> Vec<(Vec<u64>, u64)> {
    return input
        .lines()
        .map(|l| {
            let (text, value) = l.trim().split_once(" ").unwrap();
            let value = value.parse::<u64>().unwrap();
            let cards = text
                .chars()
                .map(|c| {
                    if c.to_digit(10).is_some() {
                        return c.to_digit(10).unwrap() as u64;
                    } else {
                        match c {
                            'A' => return 14,
                            'K' => return 13,
                            'Q' => return 12,
                            'J' => return 11,
                            'T' => return 10,
                            _ => panic!("Invalid card"),
                        }
                    }
                })
                .collect::<Vec<_>>();
            return (cards, value);
        })
        .collect();
}

#[aoc(day7, part1)]
pub fn part1(game: &Vec<(Vec<u64>, u64)>) -> u64 {
    let mut scored_games = game
        .iter()
        .map(|(cards, value)| (hand_value(cards), value))
        .collect::<Vec<_>>();

    scored_games.sort_by(|(hand1, _value1), (hand2, _value2)| hand2.cmp(hand1));

    return (0..scored_games.len())
        .map(|i| scored_games[i].1 * (scored_games.len() - i) as u64)
        .sum();
}

fn hand_value(cards: &Vec<u64>) -> u64 {
    let mut number_lookup = HashMap::new();
    for card in cards {
        let count = number_lookup.entry(card).or_insert(0);
        *count += 1;
    }
    let pairs = number_lookup
        .iter()
        .filter(|(_key, value)| **value > 1)
        .map(|(_key, value)| value)
        .collect::<Vec<_>>();
    let raw_card_value = list_to_num(cards);
    return calc_hand_strength(&pairs) + raw_card_value;
}

fn hand_value_2(cards: &Vec<u64>) -> u64 {
    let mut number_lookup = HashMap::new();
    for card in cards {
        let count = number_lookup.entry(card).or_insert(0);
        *count += 1;
    }

    let jokers = *number_lookup.get(&1).unwrap_or(&0);
    number_lookup.remove(&1);

    let raw_card_value = list_to_num(cards);
    if jokers == 5 {
        let five_kind: u64 = 1e15 as u64;
        return five_kind + raw_card_value;
    }

    let mut pairs = number_lookup
        .iter()
        .map(|(key, value)| {
            if **key != 1 {
                return (key, *value);
            }
            return (key, *value);
        })
        .map(|(_key, value)| value)
        .collect::<Vec<_>>();
    pairs.sort_by(|a, b| b.cmp(a));

    pairs[0] += jokers;

    let filtered_cards: Vec<&u64> = pairs.iter().filter(|value| **value > 1).collect::<Vec<_>>();
    return calc_hand_strength(&filtered_cards) + raw_card_value;
}

fn calc_hand_strength(pairs: &Vec<&u64>) -> u64 {
    if pairs.len() == 0 {
        return 0;
    }

    // Multiples of 10 chosen for each easily printing the hand value
    let five_kind = 1e15 as u64;
    let four_kind = 1e14 as u64;
    let full_house = 1e13 as u64;
    let three_kind = 1e12 as u64;
    let two_pair = 1e11 as u64;
    let pair = 1e10 as u64;

    if pairs.len() == 1 {
        return match *pairs[0] {
            5 => five_kind,
            4 => four_kind,
            3 => three_kind,
            2 => pair,
            _ => panic!("Invalid number of pairs"),
        };
    }
    if *pairs[0] == 3 || *pairs[1] == 3 {
        return full_house;
    }

    return two_pair;
}

fn list_to_num(cards: &Vec<u64>) -> u64 {
    return cards.iter().fold(0, |acc, card| acc * 100 + card);
}

#[aoc(day7, part2)]
pub fn part2(game: &Vec<(Vec<u64>, u64)>) -> u64 {
    let mut scored_games = game
        .iter()
        .map(|(cards, value)| {
            (
                cards
                    .iter()
                    .map(|card| return if *card == 11 { 1 } else { *card })
                    .collect::<Vec<_>>(),
                value,
            )
        })
        .map(|(cards, value)| (hand_value_2(&cards), *value))
        .collect::<Vec<(u64, u64)>>();

    scored_games.sort_by(|(hand1, _value1), (hand2, _value2)| hand2.cmp(hand1));

    return (0..scored_games.len())
        .map(|i| scored_games[i].1 * (scored_games.len() - i) as u64)
        .sum();
}

#[cfg(test)]
mod tests {
    use super::{input_generator, part1, part2};

    #[test]
    fn basics() {
        let input = "32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483";
        assert_eq!(part1(&input_generator(input)), 6440);
        assert_eq!(part2(&input_generator(input)), 5905);
    }

    #[test]
    fn test_more() {
        let input = "    AAAAA 2
    22222 3
    AAAAK 5
    22223 7
    AAAKK 11
    22233 13
    AAAKQ 17
    22234 19
    AAKKQ 23
    22334 29
    AAKQJ 31
    22345 37
    AKQJT 41
    23456 43";
        assert_eq!(part1(&input_generator(input)), 1343);
    }
}
