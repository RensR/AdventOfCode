#[derive(Debug)]
pub struct Game {
    id: u32,
    games: Vec<Round>,
}

#[derive(Debug)]
pub struct Round {
    blue: u32,
    red: u32,
    green: u32
}

#[aoc_generator(day2)]
pub fn input_generator(input: &str) -> Vec<Game> {
    input
        .lines()
        .map(|l| {
            let mut game_info = l.trim().split(':');
            let mut game = Game {
                id: game_info.next().unwrap().strip_prefix("Game ").unwrap().parse().unwrap(),
                games: Vec::new(),
            };
            game.games = game_info.next().unwrap().split(';').map(|r| {
                let mut cur_round = Round {
                    blue: 0,
                    red: 0,
                    green: 0,
                };

                for rgb_part in r.split(',').map(|i| i.trim()) {
                    let part = rgb_part.split(' ').collect::<Vec<_>>();
                    match part[1] {
                        "blue" => {
                            cur_round.blue = part[0].parse().unwrap();
                        },
                        "red" => {
                            cur_round.red = part[0].parse().unwrap();
                        },
                        "green" => {
                            cur_round.green = part[0].parse().unwrap();
                        },
                        _ => {
                            panic!("Invalid color");
                        }
                    }
                }
                return cur_round;
            }).collect();
            return game;
        }).collect()
}

// Determine which games would have been possible if the bag had been loaded with only
// 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
#[aoc(day2, part1)]
pub fn part1(input: &Vec<Game>) -> u32 {
    let max_red = 12;
    let max_green = 13;
    let max_blue = 14;

    return input.iter()
        .filter(|gm| {
            for round in gm.games.iter() {
                if round.red > max_red || round.blue > max_blue || round.green > max_green {
                    return false;
                }
            }
            return true;
        })
        .map(|g| { g.id })
        .sum::<u32>();
}

#[aoc(day2, part2)]
pub fn part2(games: &Vec<Game>) -> u32 {
    return games.iter()
        .map(|round| {
            let (mut max_red, mut max_green, mut max_blue) = (0, 0,0);
            for round in round.games.iter() {
                if round.red > max_red {
                    max_red = round.red;
                }
                if round.blue > max_blue {
                    max_blue = round.blue;
                }
                if round.green > max_green {
                    max_green = round.green;
                }
            }
            return max_red * max_blue * max_green;
        }).sum::<u32>();
}

#[cfg(test)]
mod tests {
    use super::{ part1, input_generator, part2};

    #[test]
    fn basics() {
        let input = &input_generator(
            "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green");
        assert_eq!(part1(input), 8);
        assert_eq!(part2(input), 2286);
    }
}