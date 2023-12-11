use std::cmp::max;
use std::cmp::min;
#[aoc_generator(day11)]
pub fn input_generator(input: &str) -> (Vec<Vec<char>>, Vec<u64>, Vec<u64>) {
    let map = input
        .lines()
        .map(|l| l.chars().collect::<Vec<_>>())
        .collect::<Vec<Vec<char>>>();

    let mut x_expand = Vec::new();

    for i in 0..map.len() {
        if !map[i].contains(&'#') {
            x_expand.push(i as u64);
        }
    }

    let mut y_expand = Vec::new();

    for x in 0..map[0].len() {
        if map.iter().all(|line| line[x] != '#') {
            y_expand.push(x as u64);
        }
    }

    return (map, x_expand, y_expand);
}

#[aoc(day11, part1)]
pub fn part1(input: &(Vec<Vec<char>>, Vec<u64>, Vec<u64>)) -> u64 {
    return calculate_distance(input, 1);
}

#[aoc(day11, part2)]
pub fn part2(input: &(Vec<Vec<char>>, Vec<u64>, Vec<u64>)) -> u64 {
    return calculate_distance(input, 1000000 - 1);
}

fn calculate_distance(
    (map, x_expand, y_expand): &(Vec<Vec<char>>, Vec<u64>, Vec<u64>),
    expansion: u64,
) -> u64 {
    let mut galaxies = Vec::new();

    for i in 0..map.len() {
        for j in 0..map[0].len() {
            if map[i][j] == '#' {
                galaxies.push((i as i32, j as i32));
            }
        }
    }

    let mut sum = 0u64;
    for i in 0..galaxies.len() {
        for j in i..galaxies.len() {
            let left_x = min(galaxies[i].0, galaxies[j].0) as u64;
            let right_x = max(galaxies[i].0, galaxies[j].0) as u64;
            let up_y = min(galaxies[i].1, galaxies[j].1) as u64;
            let down_y = max(galaxies[i].1, galaxies[j].1) as u64;

            // add non-expanded base case
            sum += right_x - left_x + down_y - up_y;

            // add expansion
            for x_expansion in x_expand {
                if *x_expansion > left_x && *x_expansion < right_x {
                    sum += expansion;
                }
            }
            for y_expansion in y_expand {
                if *y_expansion > up_y && *y_expansion < down_y {
                    sum += expansion;
                }
            }
        }
    }

    return sum;
}

#[cfg(test)]
mod tests {
    use super::{input_generator, part1, part2};

    #[test]
    fn basics() {
        let input = "...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....";
        assert_eq!(part1(&input_generator(input)), 374);
        assert_eq!(part2(&input_generator(input)), 82000210);
    }
}
