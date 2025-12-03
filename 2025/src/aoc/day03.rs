use std::collections::HashSet;

fn parse_input(lines: &Vec<String>) -> Vec<Vec<usize>> {
    let mut batteries = vec![];
    lines.iter().for_each(|line| {
        let digits = line
            .chars()
            .map(|c| c.to_digit(10).unwrap() as usize)
            .collect::<Vec<usize>>();
        batteries.push(digits);
    });
    batteries
}

fn find_max_idx(battery: &Vec<usize>, start: usize, end: usize) -> (usize, usize) {
    battery[start..end]
        .iter()
        .enumerate()
        .max_by(|(_, a), (_, b)| a.cmp(b))
        .map(|(idx, &max)| (idx, max))
        .unwrap()
}

pub fn part1(lines: &Vec<String>, _day: usize) {
    let batteries = parse_input(lines);
    let mut joltage = 0;

    for battery in batteries {
        let n = battery.len();
        let (idx, fst_digit) = find_max_idx(&battery, 0, n);
        if idx == n - 1 {
            let (_, snd_digit) = find_max_idx(&battery, 0, n - 1);
            joltage += snd_digit * 10 + fst_digit;
        } else {
            let (_, snd_digit) = find_max_idx(&battery, idx + 1, n);
            joltage += fst_digit * 10 + snd_digit;
        }
    }

    println!("The total output joltage is: {joltage}");
}

fn find_max_idx_with_memory(
    battery: &Vec<usize>,
    seen: &HashSet<usize>,
    start: usize,
    end: usize,
) -> (usize, usize) {
    let (mut max_idx, mut max) = (0, 0);
    for i in start..end {
        if battery[i] > max && !seen.contains(&i) {
            max = battery[i];
            max_idx = i;
        }
    }
    (max_idx, max)
}

pub fn part2(lines: &Vec<String>, _day: usize) {
    let batteries = parse_input(lines);
    let mut joltage = 0;

    for battery in batteries {
        let mut digits = vec![];
        let mut padding = 11i32;
        let mut seen = HashSet::new();
        let mut start = 0;
        while padding >= 0 {
            let end = battery.len() as i32 - padding;
            let (idx, digit) = find_max_idx_with_memory(&battery, &seen, start, end as usize);
            if idx < start {
                start = idx + 1;
                continue;
            }
            digits.push(digit);
            seen.insert(idx);
            start = idx;
            padding -= 1;
        }
        let battery_joltage = digits.iter().fold(0, |acc, el| acc * 10 + el);
        joltage += battery_joltage;
    }
    println!("The new total output joltage is {joltage}");
}
