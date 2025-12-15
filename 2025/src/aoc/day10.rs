use std::{
    collections::{HashSet, VecDeque},
    process::{Command, Stdio},
};

#[derive(Debug)]
struct Machine {
    indicator: usize,
    switches: Vec<usize>,
    _jswitches: Vec<Vec<usize>>,
    _joltages: Vec<usize>,
}

impl Machine {
    fn configure_indicator(&self) -> usize {
        let mut q = VecDeque::new();
        let mut seen = HashSet::new();
        let mut button_pressed = 0;

        q.push_back((0, 0));
        while !q.is_empty() {
            let (n, current_level) = q.pop_front().unwrap();
            if n == self.indicator {
                button_pressed = current_level;
                break;
            }
            if seen.contains(&n) {
                continue;
            }
            seen.insert(n);

            for switch in self.switches.clone() {
                q.push_back((n ^ switch, current_level + 1));
            }
        }
        button_pressed
    }
}

fn parse_input(lines: &Vec<String>) -> Vec<Machine> {
    lines
        .iter()
        .map(|line| {
            let fields: Vec<_> = line.split(" ").collect();

            // parse target
            let mut target = 0;
            let size = fields[0].len() - 2;
            for (i, c) in fields[0]
                .replace("[", "")
                .replace("]", "")
                .chars()
                .enumerate()
            {
                if c == '#' {
                    target += (2 as u32).pow((size - i - 1) as u32);
                }
            }

            // parse switches
            let mut switches = vec![];
            let mut jswitches = vec![];
            for i in 1..fields.len() - 1 {
                let sws: Vec<_> = fields[i]
                    .replace("(", "")
                    .replace(")", "")
                    .split(",")
                    .map(|n| n.parse::<usize>().unwrap())
                    .collect();
                let switch = sws
                    .iter()
                    .fold(0, |acc, el| acc + (2 as u32).pow((size - el - 1) as u32));
                switches.push(switch as usize);
                jswitches.push(sws);
            }

            // parse joltages
            let joltages = fields
                .last()
                .unwrap()
                .replace("{", "")
                .replace("}", "")
                .split(",")
                .map(|n| n.parse::<usize>().unwrap())
                .collect();

            Machine {
                indicator: target as usize,
                switches: switches,
                _jswitches: jswitches,
                _joltages: joltages,
            }
        })
        .collect()
}

pub fn part1(lines: &Vec<String>, _day: usize) {
    let machines = parse_input(&lines);
    let mut button_presses = 0;
    for machine in machines {
        button_presses += machine.configure_indicator();
    }
    println!(
        "The fewest button presses required to correctly configure the indicator lights on all of the machines is {button_presses}"
    );
}

pub fn part2(_lines: &Vec<String>, _day: usize) {
    let button_presses = String::from_utf8(
        Command::new("python3")
            .arg("src/python3/day10.py")
            .stdout(Stdio::piped())
            .output()
            .unwrap()
            .stdout,
    )
    .unwrap();

    println!(
        "The fewest button presses required to correctly configure the indicator lights on all of the machines is {button_presses}"
    );
}
