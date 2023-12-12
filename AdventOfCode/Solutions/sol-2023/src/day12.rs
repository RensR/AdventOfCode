use std::collections::HashMap;

#[aoc_generator(day12)]
pub fn input_generator(input: &str) -> Vec<(Vec<char>, Vec<usize>)> {
    return input
        .lines()
        .map(|l| {
            let (spring, report) = l.split_once(" ").unwrap();
            return (
                spring.chars().collect::<Vec<_>>(),
                report
                    .split(",")
                    .map(|n| n.parse::<usize>().unwrap())
                    .collect::<Vec<_>>(),
            );
        })
        .collect::<Vec<(Vec<char>, Vec<usize>)>>();
}

#[aoc(day12, part1)]
pub fn part1(input: &Vec<(Vec<char>, Vec<usize>)>) -> u64 {
    return input
        .iter()
        .map(|(springs, report)| solve(springs, report, 0, 0, 0, &mut HashMap::new()))
        .sum();
}

fn solve(
    springs: &Vec<char>,
    report: &Vec<usize>,
    start: usize,
    rp_index: usize,
    cur_broken: usize,
    hash_map: &mut HashMap<(usize, usize, usize), u64>,
) -> u64 {
    if hash_map.contains_key(&(start, rp_index, cur_broken)) {
        return *hash_map.get(&(start, rp_index, cur_broken)).unwrap();
    }

    let mut rep_index = rp_index;
    let mut cur_broken = cur_broken;
    for i in start..springs.len() {
        if springs[i] == '.' {
            if cur_broken == 0 {
                continue;
            }
            if rep_index < report.len() && cur_broken == report[rep_index] {
                rep_index += 1;
                cur_broken = 0;
                continue;
            } else {
                hash_map.insert((i, rep_index, cur_broken), 0);
                return 0;
            }
        }

        if springs[i] == '#' {
            cur_broken += 1;
            continue;
        }

        // Included
        let mut inc_result = 0;
        if rep_index < report.len() && cur_broken < report[rep_index] {
            inc_result = solve(&springs, report, i + 1, rep_index, cur_broken + 1, hash_map);
        }

        // Not included
        return if cur_broken == 0 {
            let result = inc_result + solve(&springs, report, i + 1, rep_index, 0, hash_map);
            hash_map.insert((i, rep_index, cur_broken), result);
            result
        } else if rep_index < report.len() && cur_broken == report[rep_index] {
            let result = inc_result + solve(&springs, report, i + 1, rep_index + 1, 0, hash_map);
            hash_map.insert((i, rep_index, cur_broken), result);
            result
        } else {
            hash_map.insert((i, rep_index, cur_broken), inc_result);
            inc_result
        };
    }

    if rep_index == report.len() {
        return if cur_broken == 0 { 1 } else { 0 };
    };
    if rep_index == report.len() - 1 && cur_broken == report[rep_index] {
        return 1;
    };

    return 0;
}

#[aoc(day12, part2)]
pub fn part2(input: &Vec<(Vec<char>, Vec<usize>)>) -> u64 {
    return input
        .iter()
        .map(|(springs, report)| {
            let mut new_vec_springs = springs.to_vec();
            let extra_q = "?".chars().collect::<Vec<_>>();
            new_vec_springs.append(&mut extra_q.to_vec());
            new_vec_springs.append(&mut springs.to_vec());
            new_vec_springs.append(&mut extra_q.to_vec());
            new_vec_springs.append(&mut springs.to_vec());
            new_vec_springs.append(&mut extra_q.to_vec());
            new_vec_springs.append(&mut springs.to_vec());
            new_vec_springs.append(&mut extra_q.to_vec());
            new_vec_springs.append(&mut springs.to_vec());

            let mut new_rep = report.to_vec();
            new_rep.append(&mut report.to_vec());
            new_rep.append(&mut report.to_vec());
            new_rep.append(&mut report.to_vec());
            new_rep.append(&mut report.to_vec());

            return solve(&new_vec_springs, &new_rep, 0, 0, 0, &mut HashMap::new());
        })
        .sum();
}

#[cfg(test)]
mod tests {
    use super::{input_generator, part1, part2};

    #[test]
    fn basics() {
        let input = "???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1";
        assert_eq!(part1(&input_generator(input)), 21);
        assert_eq!(part2(&input_generator(input)), 525152);
    }

    #[test]
    fn single_line() {
        assert_eq!(part1(&input_generator("???.### 1,1,3")), 1);
        assert_eq!(part2(&input_generator("???.### 1,1,3")), 1);
    }
}
