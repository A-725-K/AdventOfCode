use std::collections::HashMap;
use std::time::Instant;
use std::{env, fs::read_to_string, process};

mod aoc;

type MethodList = HashMap<usize, Vec<fn(&Vec<String>, usize)>>;

fn add_methods(methods: &mut MethodList) {
    methods.insert(11, vec![aoc::day11::part1, aoc::day11::part2]);
    methods.insert(10, vec![aoc::day10::part1, aoc::day10::part2]);
    methods.insert(09, vec![aoc::day09::part1, aoc::day09::part2]);
    methods.insert(08, vec![aoc::day08::part1, aoc::day08::part2]);
    methods.insert(07, vec![aoc::day07::part1, aoc::day07::part2]);
    methods.insert(06, vec![aoc::day06::part1, aoc::day06::part2]);
    methods.insert(05, vec![aoc::day05::part1, aoc::day05::part2]);
    methods.insert(04, vec![aoc::day04::part1, aoc::day04::part2]);
    methods.insert(03, vec![aoc::day03::part1, aoc::day03::part2]);
    methods.insert(02, vec![aoc::day02::part1, aoc::day02::part2]);
    methods.insert(01, vec![aoc::day01::part1, aoc::day01::part2]);
}

fn parse_input(filename: &str) -> Vec<String> {
    match read_to_string(filename) {
        Ok(content) => content.lines().map(String::from).collect(),
        Err(e) => {
            println!("Cannot open file {filename}: {}", e);
            process::exit(1);
        }
    }
}

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() == 2 && (args[1] == "-h" || args[1] == "--help" || args[1] == "-?") {
        println!("Usage: cargo run -- <DAY: 1|..|25> <PART: 1|2> [<small|big>]");
        process::exit(0);
    }

    if args.len() < 2 || args.len() > 4 {
        println!("Expecting 2 or 3 arguments: <DAY: 1|..|25> <PART: 1|2> [<FILENAME>]");
        process::exit(1);
    }

    let day;
    let part;
    match args[1].trim().parse::<usize>() {
        Ok(n) => {
            if n < 1 || n > 25 {
                println!("Select the correct day between 1 and 25: {n}");
                process::exit(1);
            }
            day = n;
        }
        Err(e) => {
            println!("Expecting a number, got: {}. Error: {}", args[1], e);
            process::exit(1);
        }
    }
    match args[2].trim().parse::<usize>() {
        Ok(n) => {
            if n < 1 || n > 2 {
                println!("Decide if you want to run part 1 or 2: {n}");
                process::exit(1);
            }
            part = n;
        }
        Err(e) => {
            println!("Expecting a number, got: {}. Error: {}", args[1], e);
            process::exit(1);
        }
    }

    // Read input
    let day_str = if args[1].parse::<usize>().unwrap() < 10 {
        format!("0{}", args[1])
    } else {
        format!("{}", args[1])
    };
    let base_dir = format!("src/inputs/day{}", day_str);
    let input_filename = if args.len() == 4 {
        match args[3].as_str() {
            "small" => format!("{base_dir}/mini_input"),
            "big" => format!("{base_dir}/input"),
            s => format!("{base_dir}/{s}"),
        }
    } else {
        format!("{base_dir}/mini_input")
    };
    let input = parse_input(&input_filename);

    // Intialize all days until today
    let mut methods: MethodList = HashMap::new();
    add_methods(&mut methods);

    if let Some(day_funcs) = methods.get(&day) {
        let start = Instant::now();
        day_funcs[part - 1](&input, day);
        let duration = start.elapsed();
        println!("Time: {:?}", duration);
    } else {
        println!("Day {day} not yet implemented!");
        process::exit(1);
    }
}
