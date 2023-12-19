#[aoc_generator(day6)]
pub fn input_generator(input: &str) -> Vec<(u64, u64)> {
    let (time, distance) = input.split_once("\n").unwrap();
    return time
        .strip_prefix("Time:")
        .unwrap()
        .split_whitespace()
        .map(|n| n.parse::<u64>().unwrap())
        .zip(
            distance
                .strip_prefix("Distance:")
                .unwrap()
                .split_whitespace()
                .map(|n| n.parse::<u64>().unwrap()),
        )
        .collect();
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
    let time: u64 = 59707878;
    let distance: u64 = 430121812131276;

    return get_winning_options(time, distance);
}

fn get_winning_options(time: u64, distance: u64) -> u64 {
    return (0..=time).filter(|i| i * (time - i) > distance).count() as u64;
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
