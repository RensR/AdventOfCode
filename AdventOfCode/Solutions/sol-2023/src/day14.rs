use grid::*;
use std::collections::HashSet;

#[aoc_generator(day14)]
pub fn input_generator(input: &str) -> (Grid<char>, HashSet<(usize, usize)>) {
    let length = input.find('\n').unwrap();
    let mut movable_stones = HashSet::new();
    let mut grd = Grid::from_vec(
        input
            .chars()
            .filter_map(|c| match c {
                '\n' => None,
                _ => Some(c),
            })
            .collect::<Vec<_>>(),
        length,
    );
    grd.rotate_left();
    grd.flip_rows();

    for i in 0..grd.cols() {
        for j in 0..grd.rows() {
            if grd[(i, j)] == 'O' {
                movable_stones.insert((i, j));
            }
        }
    }

    return (grd, movable_stones);
}

#[aoc(day14, part1)]
pub fn part1((grid, mv_stones): &(Grid<char>, HashSet<(usize, usize)>)) -> u64 {
    let mut mv_stone = mv_stones.clone();
    keep_moving(grid, (0, -1), &mut mv_stone);
    return count_stone(&mv_stone, grid.rows() as u64);
}

// 103356 high
#[aoc(day14, part2)]
pub fn part2((input, mv_stone): &(Grid<char>, HashSet<(usize, usize)>)) -> u64 {
    let mut mv_stone = mv_stone.clone();

    for i in 0..1000000000 {
        if i % 1000 == 0 {
            println!("{}", i);
        }
        keep_moving(&input, (0, -1), &mut mv_stone);
        keep_moving(&input, (-1, 0), &mut mv_stone);
        keep_moving(&input, (0, 1), &mut mv_stone);
        keep_moving(&input, (1, 0), &mut mv_stone);
    }

    return count_stone(&mv_stone, input.rows() as u64);
}

fn keep_moving(grid: &Grid<char>, direction: (i32, i32), mv_stone: &mut HashSet<(usize, usize)>) {
    let mut moved = true;
    while moved {
        // print_grid(&grid, &mv_stone);
        moved = sim_mv(&grid, direction, mv_stone);
    }
}

fn sim_mv(
    grid: &Grid<char>,
    direction: (i32, i32),
    mv_stone: &mut HashSet<(usize, usize)>,
) -> bool {
    let mut moved = false;
    for (x, y) in mv_stone.clone().into_iter() {
        let (nx, ny) = (x as i32 + direction.0, y as i32 + direction.1);
        if nx < 0 || ny < 0 || nx >= grid.cols() as i32 || ny >= grid.rows() as i32 {
            continue;
        }
        let (nx, ny) = (nx as usize, ny as usize);

        if grid[(nx, ny)] == '.' || grid[(nx, ny)] == 'O' {
            if mv_stone.contains(&(nx, ny)) {
                continue;
            } else {
                moved = true;
                mv_stone.remove(&(x, y));
                mv_stone.insert((nx, ny));
            }
        }
    }
    return moved;
}

fn count_stone(mv_stone: &HashSet<(usize, usize)>, height: u64) -> u64 {
    return mv_stone.iter().map(|(_x, y)| height - *y as u64).sum();
}

#[cfg(test)]
mod tests {
    use super::{input_generator, part1, part2};

    #[test]
    fn basics() {
        let input = "O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....";
        assert_eq!(part1(&input_generator(input)), 136);
        assert_eq!(part2(&input_generator(input)), 64);
    }
}

fn print_grid(grid: &Grid<char>, mv_stone: &HashSet<(usize, usize)>) {
    for y in 0..grid.rows() {
        for x in 0..grid.cols() {
            if mv_stone.contains(&(x, y)) {
                print!("O");
                continue;
            }
            print!("{}", grid[(x, y)]);
        }
        println!();
    }
    println!();
}
