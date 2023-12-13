#[aoc_generator(day13)]
pub fn input_generator(input: &str) -> Vec<Vec<Vec<u32>>> {
    return input
        .split("\n\n")
        .map(|puzzle| {
            puzzle
                .lines()
                .map(|line| {
                    line.chars()
                        .map(|c| match c {
                            '#' => 1,
                            '.' => 0,
                            _ => panic!("invalid char"),
                        })
                        .collect()
                })
                .collect()
        })
        .collect();
}

#[aoc(day13, part1)]
pub fn part1(input: &Vec<Vec<Vec<u32>>>) -> u64 {
    return input
        .iter()
        .map(|game| calculate_fold(game, 0).unwrap())
        .sum();
}

fn calculate_fold(game: &Vec<Vec<u32>>, prev: u64) -> Option<u64> {
    let bit_val: Vec<_> = game
        .iter()
        .map(|line| {
            let mut cur = 0;
            for c in line {
                cur = cur << 1 | c;
            }
            cur
        })
        .collect();

    let check_hor = check_folds(bit_val, 100, prev);
    if check_hor != None {
        return check_hor;
    }

    let mut bit_val_vec = Vec::new();
    for i in 0..game[0].len() {
        let mut cur = 0;
        for line in game {
            cur = cur << 1 | line[i]
        }
        bit_val_vec.push(cur);
    }

    return check_folds(bit_val_vec, 1, prev);
}

fn check_folds(lines: Vec<u32>, multiplier: u64, prev: u64) -> Option<u64> {
    let found_vert_normal = find_fold(lines.clone());
    if found_vert_normal != None {
        let res = found_vert_normal.unwrap() * multiplier;
        if res != prev {
            return Some(res);
        }
    }

    let found_vert_flip = find_fold(lines.iter().rev().map(|v| *v).collect());
    if found_vert_flip != None {
        let res = (lines.len() as u64 - found_vert_flip.unwrap()) * multiplier;
        if res != prev {
            return Some(res);
        }
    }

    return None;
}

fn find_fold(lines: Vec<u32>) -> Option<u64> {
    for start_i in 0..(lines.len()) {
        let mut j = lines.len() - 1;
        let mut i = start_i;

        let mut found = true;
        while i < j {
            if lines[i] != lines[j] {
                found = false;
                break;
            }
            i += 1;
            j -= 1;
        }

        if found && i != j {
            return Some(i as u64);
        }
    }

    return None;
}

#[aoc(day13, part2)]
pub fn part2(input: &Vec<Vec<Vec<u32>>>) -> u64 {
    return input
        .iter()
        .map(|game| {
            for i in 0..game.len() {
                for j in 0..game[i].len() {
                    let og_score: u64 = calculate_fold(game, 0).unwrap();
                    let mut new_game = game.clone();
                    if new_game[i][j] == 1 {
                        new_game[i][j] = 0;
                    } else {
                        new_game[i][j] = 1;
                    }
                    let maybe_score = calculate_fold(&new_game, og_score);
                    if maybe_score != None {
                        let score = maybe_score.unwrap();
                        if score != og_score {
                            return score;
                        }
                    }
                }
            }
            panic!("no solution found")
        })
        .sum();
}

#[cfg(test)]
mod tests {
    use super::{input_generator, part1, part2};

    #[test]
    fn basics() {
        let input = "#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#";
        assert_eq!(part1(&input_generator(input)), 405);
        assert_eq!(part2(&input_generator(input)), 400);
    }
    #[test]
    fn rotation() {
        let input = "#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#

.#.##.#.#
.##..##..
.#.##.#..
#......##
#......##
.#.##.#..
.##..##.#

#..#....#
###..##..
.##.#####
.##.#####
###..##..
#..#....#
#..##...#";
        assert_eq!(part1(&input_generator(input)), 709);
        assert_eq!(part2(&input_generator(input)), 1400);
    }

    #[test]
    fn mini() {
        let input = "..#.#.#.#.#.#.#.#.
..#.#.#.###.#.####";
        assert_eq!(part1(&input_generator(input)), 1);
    }

    #[test]
    fn first_row() {
        let input = "....#....##..##..
##.#.....#....#..
....##...#.##.#..
..##..#####..####
###..##..######..
.....##.#.#..#.#.
.#.#####.##..##.#
###.###...####...
....####.#....#.#";
        assert_eq!(part1(&input_generator(input)), 12);
        assert_eq!(part2(&input_generator(input)), 1);
    }

    #[test]
    fn same_same() {
        let input = ".###.####
.##.##..#
.###.####
#....#..#
....#....
##.##....
##.###..#
#..#..##.
.#...#..#
....#....
##.######
...###.##
#.###....
##.#.####
##.#.####";
        assert_eq!(part2(&input_generator(input)), 7);
    }
}
