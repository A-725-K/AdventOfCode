use regex::Regex;
use std::{env, fs::read_to_string};

#[derive(Debug)]
struct ClawMachine {
    a: (usize, usize),
    b: (usize, usize),
    x: (usize, usize),
}

const MAX_TRIES: usize = 100;
const COST_A: usize = 3;
const COST_B: usize = 1;
const CONVERSION_ERROR: usize = 10000000000000;

fn parse_input(filename: &str) -> Vec<ClawMachine> {
    let mut claw_machines: Vec<ClawMachine> = vec![];
    let regex = Regex::new(r".*: X.([0-9]+), Y.([0-9]+)").unwrap();
    let s = read_to_string(filename).unwrap();
    let lines: Vec<&str> = s.lines().collect();
    for i in (0..lines.len()).step_by(4) {
        let matches_a = regex.captures(lines[i]).unwrap();
        let matches_b = regex.captures(lines[i + 1]).unwrap();
        let matches_x = regex.captures(lines[i + 2]).unwrap();
        let cm = ClawMachine {
            a: (
                matches_a.get(1).unwrap().as_str().parse::<usize>().unwrap(),
                matches_a.get(2).unwrap().as_str().parse::<usize>().unwrap(),
            ),
            b: (
                matches_b.get(1).unwrap().as_str().parse::<usize>().unwrap(),
                matches_b.get(2).unwrap().as_str().parse::<usize>().unwrap(),
            ),
            x: (
                matches_x.get(1).unwrap().as_str().parse::<usize>().unwrap(),
                matches_x.get(2).unwrap().as_str().parse::<usize>().unwrap(),
            ),
        };
        claw_machines.push(cm);
    }
    claw_machines
}

fn win_prize(claw_machine: &ClawMachine) -> usize {
    let mut candidates = vec![];
    for i in 0..MAX_TRIES {
        for j in 0..MAX_TRIES {
            if claw_machine.a.0 * i + claw_machine.b.0 * j == claw_machine.x.0 {
                candidates.push((i, j));
            }
        }
    }
    let mut min_cost = usize::MAX;
    for (press_a, press_b) in candidates {
        if claw_machine.a.1 * press_a + claw_machine.b.1 * press_b == claw_machine.x.1 {
            let cost = press_a * COST_A + press_b * COST_B;
            if cost < min_cost {
                min_cost = cost;
            }
        }
    }
    if min_cost == usize::MAX {
        return 0;
    }
    min_cost
}

fn part1(claw_machines: &Vec<ClawMachine>) {
    let mut tot_tokens = 0;
    for claw_machine in claw_machines {
        tot_tokens += win_prize(claw_machine);
    }
    println!("To win all possible prizes it costs {tot_tokens} tokens");
}

fn part2(claw_machines: &mut Vec<ClawMachine>, debug: bool) {
    let mut tot_tokens = 0;
    for claw_machine in &mut *claw_machines {
        claw_machine.x.0 += CONVERSION_ERROR;
        claw_machine.x.1 += CONVERSION_ERROR;
    }

    for (i, claw_machine) in claw_machines.into_iter().enumerate() {
        let alpha = (claw_machine.a.1 * claw_machine.x.0) as i64;
        let beta = (claw_machine.a.0 * claw_machine.x.1) as i64;
        let n = (claw_machine.b.1 * claw_machine.a.0) as i64
            - (claw_machine.a.1 * claw_machine.b.0) as i64;
        if debug {
            println!("{:?}", claw_machine);
            println!("alpha={alpha} beta={beta} n={n}");
        }
        if (alpha - beta).abs() % n.abs() == 0 {
            let b_press_count = (alpha - beta).abs() / n.abs();
            let a_press_count = (claw_machine.x.0 as i64 - claw_machine.b.0 as i64 * b_press_count)
                / claw_machine.a.0 as i64;
            if debug {
                println!(
                    "a=({} - {} * {}) / {}={a_press_count}",
                    claw_machine.x.0, claw_machine.b.0, b_press_count, claw_machine.a.0
                );
                println!("b=({alpha} - {beta}) / {n}={b_press_count}");
                println!("({i}) WIN! a={a_press_count} b={b_press_count}\n");
            }
            tot_tokens += a_press_count as usize * COST_A + b_press_count as usize * COST_B;
        } else if debug {
            println!("no\n");
        }
    }
    println!(
        "To win all possible prizes after fixing the conversion error it costs {tot_tokens} tokens"
    );
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let mut claw_machines = parse_input(&args[1]);
    part1(&claw_machines);
    part2(&mut claw_machines, false);
}
