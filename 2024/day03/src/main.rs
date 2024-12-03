use regex::Regex;
use std::{env, fs::read_to_string};

const ALL_REGEX: &str = r"(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don't\(\))";
const MUL_REGEX: &str = r"mul\((?<fst>[0-9]{1,3}),(?<snd>[0-9]{1,3})\)";

fn parse_input(filename: &str) -> String {
    read_to_string(filename).unwrap()
}

fn part1(memory: &String) {
    let regex = Regex::new(MUL_REGEX).unwrap();
    let mut val: i64 = 0;
    for (_, [fst, snd]) in regex.captures_iter(memory).map(|c| c.extract()) {
        let (n1, n2): (i64, i64) = (
            fst.parse().unwrap_or_default(),
            snd.parse().unwrap_or_default(),
        );
        val += n1 * n2;
    }
    println!("The final result is {val}");
}

fn part2(memory: &String) {
    let all_regex = Regex::new(ALL_REGEX).unwrap();
    let mul_regex = Regex::new(MUL_REGEX).unwrap();

    let mut val: i64 = 0;
    let mut enabled = true;
    for (_, [cmd]) in all_regex.captures_iter(memory).map(|c| c.extract()) {
        match cmd {
            "do()" => enabled = true,
            "don't()" => enabled = false,
            mul => {
                if enabled {
                    let Some(matches) = mul_regex.captures(mul) else {
                        panic!("No match!")
                    };
                    let (n1, n2): (i64, i64) = (
                        matches["fst"].parse().unwrap_or_default(),
                        matches["snd"].parse().unwrap_or_default(),
                    );
                    val += n1 * n2;
                }
            }
        };
    }
    println!("The refined final value is {val}");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let memory = parse_input(&args[1]);
    part1(&memory);
    part2(&memory);
}
