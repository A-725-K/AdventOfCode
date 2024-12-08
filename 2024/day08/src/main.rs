use itertools::Itertools;
use std::{
    collections::{HashMap, HashSet},
    env,
    fs::read_to_string,
};

fn parse_input(filename: &str) -> (HashMap<char, Vec<(i32, i32)>>, i32, i32) {
    let mut antennas: HashMap<char, Vec<(i32, i32)>> = HashMap::new();
    let (mut rows, mut cols): (i32, i32) = (0, 0);
    let mut read_cols = true;
    for (y, line) in read_to_string(filename).unwrap().lines().enumerate() {
        for (x, c) in line.chars().enumerate() {
            if read_cols {
                cols += 1;
            }
            if c != '.' {
                let antenna = antennas.entry(c).or_insert(vec![]);
                antenna.push((x as i32, y as i32));
            }
        }
        rows += 1;
        read_cols = false;
    }
    (antennas, rows, cols)
}

fn generate_antinodes(
    antennas_pair: &Vec<&(i32, i32)>,
    rows: i32,
    cols: i32,
    with_resonation: bool,
) -> Vec<(i32, i32)> {
    let (mut ant1, mut ant2) = (*antennas_pair[0], *antennas_pair[1]);

    let delta_x: i32 = ant1.0.abs_diff(ant2.0) as i32;
    let delta_y: i32 = ant1.1.abs_diff(ant2.1) as i32;

    let mut new_valid_antinodes = vec![];
    if with_resonation {
        new_valid_antinodes.push(ant1);
        new_valid_antinodes.push(ant2);
    }
    let (mut can_generate, mut can_generate_from_1, mut can_generate_from_2) = (true, true, true);
    let generate = |p1: (i32, i32), p2: (i32, i32)| {
        (
            if p1.0 < p2.0 {
                p1.0 - delta_x
            } else {
                p1.0 + delta_x
            },
            if p1.1 < p2.1 {
                p1.1 - delta_y
            } else {
                p1.1 + delta_y
            },
        )
    };
    while can_generate {
        let new_a1 = generate(ant1, ant2);
        let new_a2 = generate(ant2, ant1);
        if new_a1.0 >= 0 && new_a1.0 < cols && new_a1.1 >= 0 && new_a1.1 < rows {
            new_valid_antinodes.push(new_a1);
            ant1 = new_a1;
        } else {
            can_generate_from_1 = false;
        }
        if new_a2.0 >= 0 && new_a2.0 < cols && new_a2.1 >= 0 && new_a2.1 < rows {
            new_valid_antinodes.push(new_a2);
            ant2 = new_a2;
        } else {
            can_generate_from_2 = false;
        }
        if !with_resonation {
            break;
        }
        can_generate = can_generate_from_1 || can_generate_from_2;
    }
    new_valid_antinodes
}

fn part1(antennas: &HashMap<char, Vec<(i32, i32)>>, rows: i32, cols: i32) {
    let mut antinodes: HashSet<(i32, i32)> = HashSet::new();
    for (_, ants) in antennas {
        for pair in ants.into_iter().combinations(2) {
            for antinode in generate_antinodes(&pair, rows, cols, false).into_iter() {
                antinodes.insert(antinode);
            }
        }
    }
    println!(
        "The antennas generates {} antinodes without resonation",
        antinodes.len()
    );
}

fn part2(antennas: &HashMap<char, Vec<(i32, i32)>>, rows: i32, cols: i32) {
    let mut antinodes: HashSet<(i32, i32)> = HashSet::new();
    for (_, ants) in antennas {
        for pair in ants.into_iter().combinations(2) {
            for antinode in generate_antinodes(&pair, rows, cols, true).into_iter() {
                antinodes.insert(antinode);
            }
        }
    }
    println!(
        "The antennas generates {} antinodes with resonation",
        antinodes.len()
    );
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let (antennas, rows, cols) = parse_input(&args[1]);
    part1(&antennas, rows, cols);
    part2(&antennas, rows, cols);
}
