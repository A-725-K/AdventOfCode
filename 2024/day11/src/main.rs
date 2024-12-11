// use cached::proc_macro::cached;
use std::{collections::HashMap, env, fs::read_to_string};

const K: i64 = 2024;

fn parse_input(filename: &str) -> Vec<i64> {
    read_to_string(filename)
        .unwrap()
        .trim_end()
        .split(' ')
        .map(|n| n.parse::<i64>().unwrap())
        .collect()
}

fn apply_rule(n: i64) -> Vec<i64> {
    if n == 0 {
        return vec![1];
    }
    let n_str = format!("{n}");
    if n_str.len() % 2 != 0 {
        return vec![n * K];
    }
    let mid = n_str.len() / 2;
    let (fst, snd) = n_str.split_at(mid);
    vec![fst.parse::<i64>().unwrap(), snd.parse::<i64>().unwrap()]
}

// Stupid solution that generates the actual stones
fn part1(stones: &Vec<i64>, n_iter: usize) {
    let mut stones_count = 0;
    for &stone in stones {
        let mut seq: Vec<i64> = vec![stone];
        for _ in 0..n_iter {
            let mut new_seq = vec![];
            for s in seq {
                let res = apply_rule(s);
                for new_s in res {
                    new_seq.push(new_s);
                }
            }
            seq = new_seq;
        }
        stones_count += seq.len();
    }
    println!("After blinking {n_iter} times, there are {stones_count} stones");
}

// Even better by using 'cached' crate
// #[cached]
// fn cached_count(stone: i64, current_iter: usize) -> i64 {
//     if current_iter == 0 {
//         return 1;
//     }
//     // if let Some(&already_counted) = stonehenge.get(&(stone, current_iter)) {
//     //     return already_counted;
//     // }
//     if stone == 0 {
//         let count = cached_count(1, current_iter - 1);
//         // stonehenge.insert((stone, current_iter), count);
//         return count;
//     }
//     let stone_str = format!("{stone}");
//     let n = stone_str.len();
//     if n % 2 == 0 {
//         let (fst, snd) = stone_str.split_at(n / 2);
//         let count = cached_count(fst.parse::<i64>().unwrap(), current_iter - 1)
//             + cached_count(snd.parse::<i64>().unwrap(), current_iter - 1);
//         // stonehenge.insert((stone, current_iter), count);
//         return count;
//     }
//     let count = cached_count(stone * K, current_iter - 1);
//     // stonehenge.insert((stone, current_iter), count);
//     return count;
// }
//
// Smart implementation
fn cached_count(
    stonehenge: &mut HashMap<(i64, usize), i64>,
    stone: i64,
    current_iter: usize,
) -> i64 {
    if current_iter == 0 {
        return 1;
    }
    if let Some(&already_counted) = stonehenge.get(&(stone, current_iter)) {
        return already_counted;
    }
    if stone == 0 {
        let count = cached_count(stonehenge, 1, current_iter - 1);
        stonehenge.insert((stone, current_iter), count);
        return count;
    }
    let stone_str = format!("{stone}");
    let n = stone_str.len();
    if n % 2 == 0 {
        let (fst, snd) = stone_str.split_at(n / 2);
        let count = cached_count(stonehenge, fst.parse::<i64>().unwrap(), current_iter - 1)
            + cached_count(stonehenge, snd.parse::<i64>().unwrap(), current_iter - 1);
        stonehenge.insert((stone, current_iter), count);
        return count;
    }
    let count = cached_count(stonehenge, stone * K, current_iter - 1);
    stonehenge.insert((stone, current_iter), count);
    return count;
}

fn part2(stones: &Vec<i64>, n_iter: usize) {
    let mut stones_count = 0;
    let mut stonehenge: HashMap<(i64, usize), i64> = HashMap::new();
    for &stone in stones {
        stones_count += cached_count(&mut stonehenge, stone, n_iter);
    }
    println!("After blinking {n_iter} times, there are {stones_count} stones");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let stones = parse_input(&args[1]);
    part1(&stones, 25);
    part2(&stones, 75);
}
