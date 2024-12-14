use regex::Regex;
use std::{env, fs::read_to_string};

const ROWS: i32 = 103;
const COLS: i32 = 101;
const SECONDS: i32 = 100;

type SurveilledArea = Vec<Vec<i32>>;

#[derive(Debug, Clone)]
struct Robot {
    pos: (i32, i32),
    v: (i32, i32),
}

impl Robot {
    fn move_n_steps(self: &mut Self, n: i32) {
        let new_pos_to_adjust = (self.pos.0 + n * self.v.0, self.pos.1 + n * self.v.1);
        let mut new_x = new_pos_to_adjust.0 % COLS;
        let mut new_y = new_pos_to_adjust.1 % ROWS;
        // Modulo operator (%) accepts negative integers like in Go, not like Python.
        // Need to adjust it to a positive number.
        if new_x < 0 {
            new_x += COLS;
        }
        if new_y < 0 {
            new_y += ROWS;
        }
        let new_pos = (new_x, new_y);
        self.pos = new_pos;
    }
}

fn parse_input(filename: &str) -> Vec<Robot> {
    let mut robots = vec![];
    let regex = Regex::new(r"p=(.*,.*) v=(.*,.*)").unwrap();
    for line in read_to_string(filename).unwrap().lines() {
        let matches = regex.captures(line).unwrap();

        let ps: Vec<&str> = matches.get(1).unwrap().as_str().split(',').collect();
        let pos = (ps[0].parse::<i32>().unwrap(), ps[1].parse::<i32>().unwrap());
        let vs: Vec<&str> = matches.get(2).unwrap().as_str().split(',').collect();
        let v = (vs[0].parse::<i32>().unwrap(), vs[1].parse::<i32>().unwrap());

        robots.push(Robot { pos, v });
    }
    robots
}

fn print_bathroom_area(bathroom_area: &SurveilledArea) {
    for y in 0..ROWS {
        for x in 0..COLS {
            let ba = bathroom_area[y as usize][x as usize];
            if ba == 0 {
                print!(".");
            } else {
                print!("{ba}");
            }
        }
        println!("");
    }
}

fn compute_safety_factor(bathroom_area: &SurveilledArea) -> i32 {
    let (mut q1, mut q2, mut q3, mut q4) = (0, 0, 0, 0);
    let mid_row = ROWS / 2;
    let mid_col = COLS / 2;
    for y in 0..mid_row {
        for x in 0..mid_col {
            q1 += bathroom_area[y as usize][x as usize];
            q2 += bathroom_area[y as usize][(x + mid_col) as usize + 1];
            q3 += bathroom_area[(y + mid_row) as usize + 1][x as usize];
            q4 += bathroom_area[(y + mid_row) as usize + 1][(x + mid_col) as usize + 1];
        }
    }
    q1 * q2 * q3 * q4
}

fn build_bathroom_area(robots: &mut Vec<Robot>) -> SurveilledArea {
    let mut bathroom_area = vec![vec![0; COLS as usize]; ROWS as usize];
    for r in robots {
        bathroom_area[r.pos.1 as usize][r.pos.0 as usize] += 1;
    }
    bathroom_area
}

fn part1(mut robots: Vec<Robot>) {
    for r in &mut robots {
        r.move_n_steps(SECONDS);
    }
    let bathroom_area = build_bathroom_area(&mut robots);
    let safety_factor = compute_safety_factor(&bathroom_area);
    println!(
        "The safety factor of the area after {} seconds is {}",
        SECONDS, safety_factor
    );
}

fn find_tree(bathroom_area: &SurveilledArea) -> bool {
    // This is a bit cheating. I first tried just to find the following shape:
    //   *
    //  * *
    // *   *
    // But then I got more than 1 hit, then I changed to a more dense solution
    // by adding also the internal ones. To understand the shape I first
    // generated several possible candidates into a file and inspected manually
    // the output to understand what I was actually looking for :)
    for y in 0..(ROWS as usize - 2) {
        for x in 2..(COLS as usize - 2) {
            // Top of the three
            if bathroom_area[y][x] > 0
            // First level
                && bathroom_area[y + 1][x - 1] > 0
                && bathroom_area[y + 1][x] > 0
                && bathroom_area[y + 1][x + 1] > 0
            // Second level
                && bathroom_area[y + 2][x - 2] > 0
                && bathroom_area[y + 2][x - 1] > 0
                && bathroom_area[y + 2][x] > 0
                && bathroom_area[y + 2][x + 1] > 0
                && bathroom_area[y + 2][x + 2] > 0
            {
                return true;
            }
        }
    }
    false
}

fn part2(robots: &mut Vec<Robot>) {
    // There can be at most ROWS * COLS iteration after which the pattern repeats
    for seconds in 0..(ROWS * COLS) {
        for r in &mut *robots {
            r.move_n_steps(1);
        }
        let bathroom_area = build_bathroom_area(robots);
        if find_tree(&bathroom_area) {
            print_bathroom_area(&bathroom_area);
            println!(
                "To get a Christmas Tree I need to wait {} seconds",
                // Adjust the skewed iteration number, since I started from 0...
                seconds + 1
            );
            break;
        }
    }
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let mut robots = parse_input(&args[1]);
    part1(robots.clone());
    part2(&mut robots);
}
