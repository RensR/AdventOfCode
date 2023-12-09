#[aoc_generator(day9)]
pub fn input_generator(input: &str) -> Vec<Vec<i32>> {
    return input
        .lines()
        .map(|l| {
            l.split(" ")
                .map(|val| val.parse::<i32>().unwrap())
                .collect::<Vec<_>>()
        })
        .collect();
}

#[aoc(day9, part1)]
pub fn part1(game: &Vec<Vec<i32>>) -> i32 {
    return game.iter().map(|range| reduce_range(range)).sum();
}

fn reduce_range(range: &Vec<i32>) -> i32 {
    if range.iter().all(|val| *val == 0) {
        return 0;
    }
    let new_range = range.windows(2).map(|w| w[1] - w[0]).collect();
    return reduce_range(&new_range) + range.last().unwrap();
}

#[aoc(day9, part2)]
pub fn part2(game: &Vec<Vec<i32>>) -> i32 {
    return game.iter().map(|range| reduce_range_left(range)).sum();
}

fn reduce_range_left(range: &Vec<i32>) -> i32 {
    if range.iter().all(|val| *val == 0) {
        return 0;
    }
    let new_range = range.windows(2).map(|w| w[1] - w[0]).collect();
    return range.first().unwrap() - reduce_range_left(&new_range);
}

#[cfg(test)]
mod tests {
    use super::{input_generator, part1, part2};

    #[test]
    fn basics() {
        let input = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45";
        assert_eq!(part1(&input_generator(input)), 114);
        assert_eq!(part2(&input_generator(input)), 2);
    }

    #[test]
    fn idk() {
        let input = "4 -1 -6 -11";
        assert_eq!(part1(&input_generator(input)), -16);
    }
}
