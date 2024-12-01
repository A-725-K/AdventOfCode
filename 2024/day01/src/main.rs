use std::{collections::HashMap, env, fs::read_to_string};

fn parse_input(filename: &str) -> (Vec<i32>, Vec<i32>) {
    let mut locs: (Vec<i32>, Vec<i32>) = (Vec::new(), Vec::new());
    for line in read_to_string(filename).unwrap().lines() {
        let l = line.replace("   ", ";");
        let fields: Vec<&str> = l.split(';').collect();
        locs.0.push(fields[0].parse::<i32>().unwrap());
        locs.1.push(fields[1].parse::<i32>().unwrap());
    }
    return locs;
}

fn part1(locs: &mut (Vec<i32>, Vec<i32>)) {
    locs.0.sort();
    locs.1.sort();

    let iter = locs.0.iter().zip(locs.1.iter());
    let mut diff = 0;
    for el in iter {
        diff += (el.0 - el.1).abs();
    }
    println!("The diff is {}", diff);
}

fn part2(locs: &(Vec<i32>, Vec<i32>)) {
    let mut freq = HashMap::new();

    for loc_id in locs.1.iter() {
        freq.entry(loc_id).and_modify(|c| *c += 1).or_insert(1);
    }

    let mut diff = 0;
    for loc_id in locs.0.iter() {
        if let Some(n) = freq.get(loc_id) {
            diff += loc_id * n;
        }
    }
    println!("The real diff is {}", diff)
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let mut locs = parse_input(&args[1]);
    part1(&mut locs);
    part2(&locs);
}
