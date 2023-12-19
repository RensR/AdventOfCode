#[aoc(day15, part1)]
pub fn part1(input: &str) -> u64 {
    return input
        .split(",")
        .map(|step| {
            step.chars()
                .fold(0, |acc, c| ((acc + (c as u64)) * 17) % 256)
        })
        .sum();
}

#[aoc(day15, part2)]
pub fn part2(_input: &str) -> u64 {
    return 1;
}

#[cfg(test)]
mod tests {
    use super::part1;

    #[test]
    fn basics() {
        let hash = "HASH";
        assert_eq!(part1(&hash), 52);
        let input = "rn=1";
        assert_eq!(part1(&input), 30);
        // let input2 = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7";
        //  assert_eq!(part1(&input2), 1320);
    }
}
