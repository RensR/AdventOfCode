use std::collections::HashMap;

#[derive(Debug)]
pub struct Card {
    id: u32,
    winning: Vec<u32>,
    got: Vec<u32>,
}

#[aoc_generator(day4)]
pub fn input_generator(input: &str) -> Vec<Card> {
    input
        .lines()
        .map(|card_line| {
            let (game_info, games_played) = card_line.trim().split_once(':').unwrap();
            let (winning, held) = games_played.split_once("|").unwrap();
            return Card {
                id: game_info
                    .strip_prefix("Card")
                    .unwrap()
                    .trim()
                    .parse()
                    .unwrap(),
                winning: winning
                    .split_whitespace()
                    .map(|n| n.trim().parse::<u32>().unwrap())
                    .collect::<Vec<_>>(),
                got: held
                    .split_whitespace()
                    .map(|n| n.trim().parse::<u32>().unwrap())
                    .collect::<Vec<_>>(),
            };
        })
        .collect()
}

#[aoc(day4, part1)]
pub fn part1(input: &Vec<Card>) -> u32 {
    return input
        .iter()
        .map(|card| {
            card.got
                .iter()
                .filter(|got| card.winning.contains(got))
                .count() as u32
        })
        .filter(|&winners| winners > 0)
        .map(|winners| {
            if winners == 1 {
                return 1;
            }
            return 2 << (winners - 2);
        })
        .sum();
}

#[aoc(day4, part2)]
pub fn part2(input: &Vec<Card>) -> u32 {
    let mut winning_card_amount: HashMap<u32, u32> = HashMap::new();

    return input
        .iter()
        .map(|card| {
            let mut next_card = card.id + 1;
            let current_cards = winning_card_amount.get(&card.id).unwrap_or(&0) + 1;
            for got in &card.got {
                if card.winning.contains(&got) {
                    winning_card_amount.insert(
                        next_card,
                        winning_card_amount.get(&next_card).unwrap_or(&0) + current_cards,
                    );
                    next_card += 1;
                }
            }

            return winning_card_amount.get(&card.id).unwrap_or(&0) + 1;
        })
        .sum();
}

#[cfg(test)]
mod tests {
    use super::{input_generator, part1, part2};

    #[test]
    fn basics() {
        let input = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11";
        assert_eq!(part1(&input_generator(input)), 13);
        assert_eq!(part2(&input_generator(input)), 30);
    }
}
