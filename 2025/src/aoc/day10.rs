use std::collections::{HashSet, VecDeque};

#[derive(Debug)]
struct Machine {
    indicator: usize,
    switches: Vec<usize>,
    jswitches: Vec<Vec<usize>>,
    joltages: Vec<usize>,
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

    fn is_same_joltage(&self, other: &Vec<usize>) -> bool {
        self.joltages
            .iter()
            .zip(other)
            .fold(true, |acc, (a, b)| acc && a == b)
    }

    fn is_valid_joltage(&self, other: &Vec<usize>) -> bool {
        self.joltages
            .iter()
            .zip(other)
            .fold(true, |acc, (a, b)| acc && a >= b)
    }

    fn configure_joltages(&self) -> usize {
        let mut q = VecDeque::new();
        let mut seen = HashSet::new();
        let mut button_pressed = 0;
        let mut init_state = vec![];
        for _ in 0..self.joltages.len() {
            init_state.push(0);
        }

        q.push_back((init_state, 0));
        while !q.is_empty() {
            let (curr_state, current_level) = q.pop_front().unwrap();
            println!(
                "curr_state={:?} desired={:?} switches={:?} current_level={current_level} qsize={}",
                curr_state,
                self.joltages,
                self.jswitches,
                q.len()
            );
            if self.is_same_joltage(&curr_state) {
                button_pressed = current_level;
                break;
            }
            if !self.is_valid_joltage(&curr_state) {
                continue;
            }
            let key = format!("{:?}", curr_state);
            if seen.contains(&key) {
                continue;
            }
            seen.insert(key);

            for jswitch in self.jswitches.clone() {
                let mut new_state = curr_state.clone();
                let mut inc = 1;
                while self.is_valid_joltage(&new_state) {
                    for btn in jswitch.clone() {
                        new_state[btn] += 1;
                    }
                    q.push_back((new_state.clone(), current_level + inc));
                    inc += 1;
                }
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
                jswitches: jswitches,
                joltages: joltages,
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

pub fn part2(lines: &Vec<String>, _day: usize) {
    let machines = parse_input(&lines);
    // println!("{:?}", machines);
    let mut button_presses = 0;
    for (i, machine) in machines.iter().enumerate() {
        println!(">>>> {i}/{}", machines.len());
        button_presses += machine.configure_joltages();
    }
    println!("{button_presses}");
}
