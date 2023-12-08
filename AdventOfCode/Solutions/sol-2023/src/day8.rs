use num::integer::lcm;
use std::collections::HashMap;

pub struct Game {
    instructions: Vec<u8>,
    hash_map: HashMap<String, (String, String)>,
}

#[aoc_generator(day8)]
pub fn input_generator(input: &str) -> Game {
    return Game {
        instructions: input
            .lines()
            .nth(0)
            .unwrap()
            .to_string()
            .as_bytes()
            .to_vec(),
        hash_map: input
            .lines()
            .skip(2)
            .map(|l| {
                (
                    l[0..=2].to_string(),
                    (l[7..=9].to_string(), l[12..=14].to_string()),
                )
            })
            .collect::<HashMap<_, (_, _)>>(),
    };
}

#[aoc(day8, part1)]
pub fn part1(game: &Game) -> u64 {
    return run_maze(game, "AAA", true);
}

fn run_maze(game: &Game, start: &str, all_z: bool) -> u64 {
    let mut curr = start;
    let mut time = 0;

    while (all_z && curr != "ZZZ") || (!all_z && (curr.as_bytes()[2] as char) != 'Z') {
        let (left, right) = game.hash_map.get(curr).unwrap();
        if game.instructions[time % game.instructions.len()] as char == 'L' {
            curr = left;
        } else {
            curr = right;
        }
        time += 1;
    }
    return time as u64;
}

#[aoc(day8, part2)]
pub fn part2(game: &Game) -> u64 {
    return game
        .hash_map
        .keys()
        .filter_map(|start| {
            if start.as_bytes()[2] as char == 'A' {
                return Some(run_maze(game, start, false));
            }
            return None;
        })
        .into_iter()
        .reduce(|val, acc| lcm(val, acc))
        .unwrap();
}
