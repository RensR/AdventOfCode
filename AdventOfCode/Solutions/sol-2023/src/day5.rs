use std::cmp::min;

#[derive(Debug)]
pub struct Almanac {
    seed: Vec<u32>,
    lookup: Vec<Vec<(u32, u32, u32)>>,
}

#[aoc_generator(day5)]
pub fn input_generator(input: &str) -> Almanac {
    let sections: Vec<_> = input.split("\n\n").collect();
    return Almanac {
        seed: sections
            .first()
            .unwrap()
            .strip_prefix("seeds: ")
            .unwrap()
            .split_whitespace()
            .map(|n| n.trim().parse::<u32>().unwrap())
            .collect::<Vec<_>>(),
        lookup: sections
            .iter()
            .skip(1)
            .map(|section| {
                return section
                    .lines()
                    .skip(1)
                    .map(|line| {
                        let mut parts = line.split_whitespace();
                        return (
                            parts.next().unwrap().parse::<u32>().unwrap(),
                            parts.next().unwrap().parse::<u32>().unwrap(),
                            parts.next().unwrap().parse::<u32>().unwrap(),
                        );
                    })
                    .collect::<Vec<_>>();
            })
            .collect::<Vec<Vec<_>>>(),
    };
}

#[aoc(day5, part1)]
pub fn part1(input: &Almanac) -> u32 {
    return input
        .seed
        .iter()
        .map(|seed: &u32| {
            input.lookup.iter().fold(*seed, |acc: u32, lookup| {
                for (dest, source, range_length) in lookup {
                    if acc < *source {
                        continue;
                    }
                    let diff = acc - *source;
                    if diff <= *range_length {
                        return dest + diff;
                    }
                }
                return acc;
            })
        })
        .min()
        .unwrap();
}

#[aoc(day5, part2)]
pub fn part2(input: &Almanac) -> u32 {
    let mut new_seed: Vec<(u32, u32)> = Vec::new();
    for (i, _) in input.seed.iter().enumerate() {
        if i % 2 == 0 {
            new_seed.push((input.seed[i], input.seed[i + 1]));
        }
    }

    return new_seed
        .iter()
        .map(|seed| return solve(seed, &input.lookup))
        .min()
        .unwrap();
}

fn solve((seed_start, range_length): &(u32, u32), lookup: &Vec<Vec<(u32, u32, u32)>>) -> u32 {
    let mut cur_seed = *seed_start;
    let end_seed = *seed_start + *range_length;
    let mut best_res = 4294967295;

    while cur_seed < end_seed {
        let new_res = lookup.iter().fold(cur_seed, |acc: u32, lookup| {
            for (dest, source, range_length) in lookup {
                if acc < *source {
                    continue;
                }
                let diff = acc - *source;
                if diff < *range_length {
                    return dest + diff;
                }
            }

            return acc;
        });
        best_res = min(best_res, new_res);
        cur_seed = cur_seed + 1;
    }

    return best_res;
}

#[cfg(test)]
mod tests {
    use super::{input_generator, part1, part2};

    #[test]
    fn basics() {
        let input = "seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4";
        assert_eq!(part1(&input_generator(input)), 35);
        assert_eq!(part2(&input_generator(input)), 46);
    }

    #[test]
    fn minimal() {
        let input = "seeds: 54 13

bla bla bla:
0 52 2";
        assert_eq!(part2(&input_generator(input)), 54);
    }
}
