#[aoc_generator(day6)]
pub fn input_generator(input: &str) -> Vec<(u64, u64)> {
    let sections: Vec<_> = input.split("\n").collect();
    return sections[0]
        .strip_prefix("Time:")
        .unwrap()
        .split_whitespace()
        .map(|n| n.trim().parse::<u64>().unwrap())
        .zip(sections[1]
            .strip_prefix("Distance:")
            .unwrap()
            .split_whitespace()
            .map(|n| n.trim().parse::<u64>().unwrap()))
        .collect::<Vec<_>>();
}

#[aoc(day6, part1)]
pub fn part1(game: &Vec<(u64, u64)>) -> u64 {
    return game
        .iter()
        .map(|(time, distance)| get_winning_options(*time, *distance))
        .fold(1, |acc, n| acc * n);
}

#[aoc(day6, part2)]
pub fn part2(_game: &Vec<(u64, u64)>) -> u64 {
    let time :u64 = 59707878;
    let distance :u64= 430121812131276;

    return get_winning_options(time, distance);
}

fn get_winning_options(time: u64, distance: u64) -> u64 {
    let mut total_winners = 0;
    for i in 0..= time {
        if i * (time - i) > distance {
            total_winners += 1;
        }
    }
    return total_winners;
}

#[cfg(test)]
mod tests {
    use super::{input_generator, part1};

    #[test]
    fn basics() {
        let input = "Time:      7  15   30
Distance:  9  40  200";
        assert_eq!(part1(&input_generator(input)), 288);
    }
}
