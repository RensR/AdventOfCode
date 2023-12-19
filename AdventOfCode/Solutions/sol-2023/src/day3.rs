use std::collections::HashSet;

#[aoc(day3, part1)]
pub fn part1(input_s: &str) -> u32 {
    let input = input_s.lines().map(|l| l.as_bytes()).collect::<Vec<_>>();
    let height = input.len();
    let width = input[0].len();

    let mut seen: Vec<Vec<bool>> = vec![vec![false; width]; height];
    let mut total_num = 0;
    let directions: Vec<(i32, i32)> = get_direction();

    for y in 0..height {
        for x in 0..width {
            let cur = input[y][x] as char;
            if cur == '.' || seen[y][x] || cur.to_digit(10) == None {
                continue;
            }
            // get entire number
            let mut str_num = String::new();
            let mut k = x;
            let mut has_item = false;

            while k < width && (input[y][k] as char).to_digit(10) != None {
                str_num.push(input[y][k] as char);
                seen[y][k] = true;

                for (delta_x, delta_y) in &directions {
                    if delta_x + (k as i32) < 0
                        || delta_x + (k as i32) >= (height as i32)
                        || delta_y + (y as i32) < 0
                        || delta_y + (y as i32) >= (width as i32)
                    {
                        continue;
                    }
                    let found_char = input[(delta_y + (y as i32)) as usize]
                        [(delta_x + (k as i32)) as usize]
                        as char;
                    if found_char != '.' && found_char.to_digit(10) == None {
                        has_item = true;
                    }
                }
                k += 1;
            }

            if has_item {
                total_num += str_num.parse::<u32>().unwrap();
            }
        }
    }

    return total_num;
}

#[aoc(day3, part2)]
pub fn part2(input_s: &str) -> u32 {
    let input = input_s.lines().map(|l| l.as_bytes()).collect::<Vec<_>>();
    let height = input.len();
    let width = input[0].len();

    let mut total_num: u32 = 0;
    let directions: Vec<(i32, i32)> = get_direction();

    for y in 0..height {
        for x in 0..width {
            if (input[y][x] as char) != '*' {
                continue;
            }
            let mut numbers_found: Vec<u32> = Vec::new();
            let mut top_left: HashSet<(usize, usize)> = HashSet::new();

            for (delta_x, delta_y) in &directions {
                if delta_x + (x as i32) < 0
                    || delta_x + (x as i32) >= (width as i32)
                    || delta_y + (y as i32) < 0
                    || delta_y + (y as i32) >= (height as i32)
                {
                    continue;
                }
                let cur_x: usize = (delta_x + (x as i32)) as usize;
                let cur_y: usize = (delta_y + (y as i32)) as usize;

                // If it is a number, get the entire number
                if (input[cur_y][cur_x] as char).to_digit(10) != None {
                    let mut start_num = cur_x;
                    while start_num > 0
                        && (input[cur_y][start_num - 1] as char).to_digit(10) != None
                    {
                        start_num -= 1;
                    }

                    if top_left.contains(&(start_num, cur_y)) {
                        continue;
                    }

                    top_left.insert((start_num, cur_y));

                    let mut str_num = String::new();

                    while start_num < width
                        && (input[cur_y][start_num] as char).to_digit(10) != None
                    {
                        str_num.push(input[cur_y][start_num] as char);
                        start_num += 1;
                    }
                    numbers_found.push(str_num.parse::<u32>().unwrap());
                }
            }

            if numbers_found.len() == 2 {
                total_num += numbers_found[0] * numbers_found[1];
            }
        }
    }

    return total_num;
}

fn get_direction() -> Vec<(i32, i32)> {
    return vec![
        (-1, -1),
        (0, -1),
        (1, -1),
        (-1, 0),
        (1, 0),
        (-1, 1),
        (0, 1),
        (1, 1),
    ];
}

#[cfg(test)]
mod tests {
    use super::{part1, part2};

    #[test]
    fn basics() {
        let input = "467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..";
        assert_eq!(part1(input), 4361);
        assert_eq!(part2(input), 467835);
    }

    #[test]
    fn eol() {
        let input = "467.114
...*..7";
        assert_eq!(part2(input), 53238);
    }

    #[test]
    fn idk() {
        let input = "**.
.66";
        assert_eq!(part2(input), 0);
    }
}
