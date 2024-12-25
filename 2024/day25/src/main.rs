use std::{env, fs::read_to_string};

fn parse_input(filename: &str) -> (Vec<Vec<usize>>, Vec<Vec<usize>>) {
    let mut keys = vec![];
    let mut locks = vec![];

    let lines = read_to_string(filename)
        .unwrap()
        .lines()
        .map(String::from)
        .collect::<Vec<String>>();

    for i in (0..lines.len()).step_by(8) {
        let mut item = vec![];
        for x in 0..5 {
            let mut count = 0;
            for y in 0..7 {
                if lines[i + y].chars().collect::<Vec<char>>()[x] == '#' {
                    count += 1;
                }
            }
            count -= 1;
            item.push(count);
        }
        if lines[i].chars().collect::<Vec<char>>()[0] == '#' {
            locks.push(item);
        } else {
            keys.push(item);
        }
    }

    (keys, locks)
}

fn are_compatible(key: &Vec<usize>, lock: &Vec<usize>) -> bool {
    key.into_iter()
        .zip(lock.into_iter())
        .map(|x| x.0 + x.1)
        .filter(|&x| x <= 5)
        .count()
        == key.len()
}

fn part1(keys: &Vec<Vec<usize>>, locks: &Vec<Vec<usize>>) {
    let mut fitting_combination = 0;
    for key in keys {
        for lock in locks {
            if are_compatible(key, lock) {
                fitting_combination += 1;
            }
        }
    }
    println!("There are {fitting_combination} fitting combinations of keys and locks");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let (keys, locks) = parse_input(&args[1]);
    part1(&keys, &locks);
}
