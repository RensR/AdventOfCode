use queues::*;
use std::collections::HashSet;

#[aoc_generator(day10)]
pub fn input_generator(input: &str) -> (Vec<Vec<char>>, (usize, usize)) {
    let map = input
        .lines()
        .map(|l| l.chars().collect::<Vec<_>>())
        .collect::<Vec<Vec<char>>>();

    let mut start = (0, 0);
    map.iter().enumerate().for_each(|(y, row)| {
        row.iter().enumerate().for_each(|(x, val)| {
            if *val == 'S' {
                start = (x, y);
            }
        });
    });

    return (map, start);
}

#[aoc(day10, part1)]
pub fn part1((map, start): &(Vec<Vec<char>>, (usize, usize))) -> i32 {
    let (mut cur_x, mut cur_y) = start;
    let mut pipe_length = 1;
    let mut last_move = Direction::Down;
    cur_y += 1;
    let mut cur_char = map[cur_y][cur_x];
    while cur_char != 'S' {
        match last_move {
            Direction::Up => match cur_char {
                '|' => {
                    cur_y -= 1;
                    last_move = Direction::Up;
                }
                'F' => {
                    cur_x += 1;
                    last_move = Direction::Right;
                }
                '7' => {
                    cur_x -= 1;
                    last_move = Direction::Left;
                }
                _ => {
                    panic!("Invalid char: {}", cur_char);
                }
            },
            Direction::Down => match cur_char {
                '|' => {
                    cur_y += 1;
                    last_move = Direction::Down;
                }
                'J' => {
                    cur_x -= 1;
                    last_move = Direction::Left;
                }
                'L' => {
                    cur_x += 1;
                    last_move = Direction::Right;
                }
                _ => {
                    panic!("Invalid char: {}", cur_char);
                }
            },
            Direction::Left => match cur_char {
                '-' => {
                    cur_x -= 1;
                    last_move = Direction::Left;
                }
                'F' => {
                    cur_y += 1;
                    last_move = Direction::Down;
                }
                'L' => {
                    cur_y -= 1;
                    last_move = Direction::Up;
                }
                _ => {
                    panic!("Invalid char: {}", cur_char);
                }
            },
            Direction::Right => match cur_char {
                '-' => {
                    cur_x += 1;
                    last_move = Direction::Right;
                }
                'J' => {
                    cur_y -= 1;
                    last_move = Direction::Up;
                }
                '7' => {
                    cur_y += 1;
                    last_move = Direction::Down;
                }
                _ => {
                    panic!("Invalid char: {}", cur_char);
                }
            },
        }

        cur_char = map[cur_y][cur_x];
        pipe_length += 1;
    }

    return pipe_length / 2;
}

#[derive(Debug)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

#[aoc(day10, part2)]
pub fn part2((map, start): &(Vec<Vec<char>>, (usize, usize))) -> i32 {
    let (mut cur_x, mut cur_y) = start;
    let mut last_move = Direction::Down;
    let mut wall = HashSet::new();
    wall.insert((cur_x, cur_y));

    cur_y += 1;
    wall.insert((cur_x, cur_y));

    let mut inner = HashSet::new();
    let mut outer = HashSet::new();

    let mut cur_char = map[cur_y][cur_x];
    while cur_char != 'S' {
        match last_move {
            Direction::Up => match cur_char {
                '|' => {
                    if cur_x != 0 {
                        inner.insert((cur_x - 1, cur_y));
                    }

                    outer.insert((cur_x + 1, cur_y));

                    cur_y -= 1;
                    last_move = Direction::Up;
                }
                'F' => {
                    if cur_y != 0 && cur_x != 0 {
                        inner.insert((cur_x - 1, cur_y - 1));
                    }
                    outer.insert((cur_x + 1, cur_y + 1));

                    cur_x += 1;
                    last_move = Direction::Right;
                }
                '7' => {
                    if cur_x != 0 {
                        inner.insert((cur_x - 1, cur_y + 1));
                    }
                    if cur_y != 0 {
                        outer.insert((cur_x + 1, cur_y - 1));
                        outer.insert((cur_x, cur_y - 1));
                    }
                    outer.insert((cur_x + 1, cur_y));

                    cur_x -= 1;
                    last_move = Direction::Left;
                }
                _ => {
                    panic!("Invalid char: {}", cur_char);
                }
            },
            Direction::Down => match cur_char {
                '|' => {
                    inner.insert((cur_x + 1, cur_y));
                    if cur_x != 0 {
                        outer.insert((cur_x - 1, cur_y));
                    }

                    cur_y += 1;
                    last_move = Direction::Down;
                }
                'J' => {
                    if cur_y != 0 && cur_x != 0 {
                        outer.insert((cur_x - 1, cur_y - 1));
                    }
                    inner.insert((cur_x + 1, cur_y + 1));

                    cur_x -= 1;
                    last_move = Direction::Left;
                }
                'L' => {
                    if cur_x != 0 {
                        outer.insert((cur_x - 1, cur_y + 1));
                    }
                    if cur_y != 0 {
                        inner.insert((cur_x + 1, cur_y - 1));
                    }

                    cur_x += 1;
                    last_move = Direction::Right;
                }
                _ => {
                    panic!("Invalid char: {}", cur_char);
                }
            },
            Direction::Left => match cur_char {
                '-' => {
                    inner.insert((cur_x, cur_y + 1));
                    if cur_y != 0 {
                        outer.insert((cur_x, cur_y - 1));
                    }

                    cur_x -= 1;
                    last_move = Direction::Left;
                }
                'F' => {
                    if cur_x != 0 {
                        outer.insert((cur_x - 1, cur_y));
                        if cur_y != 0 {
                            outer.insert((cur_x - 1, cur_y - 1));
                        }
                    }
                    if cur_y != 0 {
                        outer.insert((cur_x , cur_y - 1));
                    }

                    inner.insert((cur_x + 1, cur_y + 1));

                    cur_y += 1;
                    last_move = Direction::Down;
                }
                'L' => {
                    if cur_x != 0 {
                        inner.insert((cur_x - 1, cur_y + 1));
                    }
                    if cur_y != 0 {
                        outer.insert((cur_x + 1, cur_y - 1));
                    }

                    cur_y -= 1;
                    last_move = Direction::Up;
                }
                _ => {
                    panic!("Invalid char: {}", cur_char);
                }
            },
            Direction::Right => match cur_char {
                '-' => {
                    if cur_y != 0 {
                        inner.insert((cur_x, cur_y - 1));
                    }
                    outer.insert((cur_x, cur_y + 1));

                    cur_x += 1;
                    last_move = Direction::Right;
                }
                'J' => {
                    if cur_y != 0 && cur_x != 0 {
                        inner.insert((cur_x - 1, cur_y - 1));
                    }
                    outer.insert((cur_x + 1, cur_y + 1));
                    outer.insert((cur_x + 1, cur_y ));

                    cur_y -= 1;
                    last_move = Direction::Up;
                }
                '7' => {
                    if cur_x != 0 {
                        outer.insert((cur_x - 1, cur_y + 1));
                    }
                    if cur_y != 0 {
                        inner.insert((cur_x + 1, cur_y - 1));
                    }

                    cur_y += 1;
                    last_move = Direction::Down;
                }
                _ => {
                    panic!("Invalid char: {}", cur_char);
                }
            },
        }
        wall.insert((cur_x, cur_y));
        cur_char = map[cur_y][cur_x];
    }

    // Remove the walls from the inner and outer sets
    inner.retain(|&loc| !wall.contains(&loc));
    outer.retain(|&loc| !wall.contains(&loc));

    print(map, outer.clone(), wall.clone());

    let result_maybe_in = calc(inner.clone(), wall.clone());
    let result_maybe_out = calc(outer.clone(), wall.clone());

    println!("{} {}", result_maybe_in, result_maybe_out);

    // Since we set the outer result to 10000000 we can return the smaller one
    // and that should be the correct answer.
    return if result_maybe_in < result_maybe_out { result_maybe_in } else { result_maybe_out };
}

fn calc(
    mut inner: HashSet<(usize, usize)>,
    wall: HashSet<(usize, usize)>,
) -> i32 {
    let mut inner_seen: Queue<(usize, usize)> = queue![];

    inner.iter().enumerate().for_each(|(_i, (x, y))| {
        let _ = inner_seen.add((*x, *y));
    });

    while inner_seen.size() != 0 {
        let (x, y) = inner_seen.remove().unwrap();

        let dir = vec![(-1i32, 0i32), (1, 0), (0, -1), (0, 1)];

        for (dx, dy) in dir {
            if x as i32 + dx < 0 || y as i32 + dy < 0 {
                return 10000000;
            }

            let new_loc = ((x as i32 + dx) as usize, (y as i32 + dy) as usize);
            if wall.contains(&new_loc) || inner.contains(&new_loc) {
                continue;
            }

            inner.insert(new_loc);
            let _ = inner_seen.add(new_loc);
        }
    }

    return inner.len() as i32;
}

fn print(map: &Vec<Vec<char>>, inner: HashSet<(usize, usize)>, wall: HashSet<(usize, usize)>) {
    println!("");
    for (y, row) in map.iter().enumerate() {
        let mut line = "".to_owned();
        for (x, _val) in row.iter().enumerate() {
            if inner.contains(&(x, y)) {
                line = line + "I";
            } else if wall.contains(&(x, y)) {
                line = line + "#";
            } else {
                line = line + map[y][x].to_string().as_str();
            }
        }
        println!("{}", line);
    }
    println!("");
}

#[cfg(test)]
mod tests {
    use super::{input_generator, part1, part2};

    #[test]
    fn basics() {
        let input = "..F7.
.FJ|.
SJ.L7
|F--J
LJ...";
        assert_eq!(part1(&input_generator(input)), 8);
        assert_eq!(part2(&input_generator(input)), 1);
    }

    #[test]
    fn part_2() {
        let input = "...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........";
        assert_eq!(part2(&input_generator(input)), 4);
    }

    #[test]
    fn part_inner_inner() {
        let input = "..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........";
        assert_eq!(part2(&input_generator(input)), 4);
    }

    #[test]
    fn part_idk() {
        let input = ".F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...";
        assert_eq!(part2(&input_generator(input)), 8);
    }

    #[test]
    fn part_agggg() {
        let input = "FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJIF7FJ-
L---JF-JLJIIIIFJLJJ7
|F|F-JF---7IIIL7L|7|
|FFJF7L7F-JF7IIL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L";
        assert_eq!(part2(&input_generator(input)), 10);
    }
}
