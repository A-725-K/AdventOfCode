use std::{env, fs::read_to_string};

fn parse_input(filename: &str) -> Vec<String> {
    read_to_string(filename)
        .unwrap()
        .lines()
        .map(String::from)
        .collect()
}

fn part1(lines: &Vec<String>) {}

fn part2(lines: &Vec<String>) {}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let lines = parse_input(&args[1]);
    part1(&lines);
    part2(&lines);
}
