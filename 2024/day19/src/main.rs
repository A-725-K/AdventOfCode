use std::{
    collections::{HashMap, HashSet, VecDeque},
    env,
    fs::read_to_string,
};

fn parse_input(filename: &str) -> (Vec<String>, Vec<String>) {
    let file = read_to_string(filename).unwrap();
    let lines = file.lines().map(String::from).collect::<Vec<String>>();
    let patterns: Vec<String> = (&lines)[0]
        .split(",")
        .map(|s| s.trim().to_string())
        .collect();
    let towels = lines.clone().split_off(2);
    (patterns, towels)
}

// Yay! I got it! :)
fn can_be_generated(towel: &String, patterns: &Vec<String>) -> bool {
    let mut q = VecDeque::new();
    q.push_back(String::from(towel));

    let mut memo = HashSet::new();
    while let Some(s) = q.pop_front() {
        for p in patterns {
            if s.starts_with(p) {
                let scp = s.clone();
                let (_, t_str) = scp.split_at(p.len());
                if t_str.is_empty() {
                    return true;
                }
                if !memo.contains(t_str) {
                    q.push_back(String::from(t_str));
                    memo.insert(String::from(t_str));
                }
            }
        }
    }
    false
}

fn part1(patterns: &Vec<String>, towels: &Vec<String>) {
    let mut possible_designs = 0;
    for towel in towels {
        if can_be_generated(towel, patterns) {
            possible_designs += 1;
        }
    }
    println!("There are {possible_designs} possible designs");
}

// Unfortunately the following solution is not entirely mine :(
//
// This approach is different: look for a pattern from the end of the string
// instead of from the beginning. Keep memoized the count of combinations for
// each substring you are considering.
fn count_combinations(
    towel: &String,
    patterns: &Vec<String>,
    max_pattern_len: usize,
    cache: &mut HashMap<String, usize>,
) -> usize {
    if towel.is_empty() {
        return 1;
    }
    if let Some(cached) = cache.get(towel) {
        return *cached;
    }
    // OBS: +1 to consider also the whole string
    // To optimize even more, consider only the strings that are shorter or the
    // same length as the longest pattern. After that we MUST have some
    // repetition, or in another way, longer strings should've been generated
    // with a combination of some other patters.
    let end = usize::min(towel.len(), max_pattern_len) + 1;
    let mut combinations = 0;
    for i in 0..end {
        let (radix, suffix) = towel.split_at(i);
        if patterns.contains(&String::from(radix)) {
            combinations +=
                count_combinations(&suffix.to_string(), patterns, max_pattern_len, cache);
        }
    }
    cache.insert(towel.to_string(), combinations);
    combinations
}

fn part2(patterns: &Vec<String>, towels: &Vec<String>) {
    let mut all_possible_designs = 0;
    let max_pattern_len = patterns.into_iter().map(|pt| pt.len()).max().unwrap();
    let mut cache = HashMap::new();
    for towel in towels {
        let combinations = count_combinations(towel, patterns, max_pattern_len, &mut cache);
        all_possible_designs += combinations;
    }
    println!("There are {all_possible_designs} possible designs combinations");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let (patterns, towels) = parse_input(&args[1]);
    part1(&patterns, &towels);
    part2(&patterns, &towels);
}
