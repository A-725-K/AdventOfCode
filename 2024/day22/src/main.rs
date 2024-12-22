use std::{
    collections::{HashMap, HashSet},
    env,
    fs::read_to_string,
};

const ITERATIONS: usize = 2000;
const MODULO: usize = 16777216;

fn parse_input(filename: &str) -> Vec<usize> {
    read_to_string(filename)
        .unwrap()
        .lines()
        .map(|n| n.parse::<usize>().unwrap())
        .collect()
}

fn evolve(mut n: usize) -> (usize, Vec<(usize, i32)>) {
    let mut changes = vec![];
    let mut prev = (n % 10) as i32;
    for _ in 0..ITERATIONS {
        let x = n * 64;
        n ^= x;
        n %= MODULO;

        let y = n / 32;
        n ^= y;
        n %= MODULO;

        let z = n * 2048;
        n ^= z;
        n %= MODULO;

        changes.push((n % 10, (n % 10) as i32 - prev));
        prev = (n % 10) as i32;
    }
    (n, changes)
}

fn part1(secret_numbers: &Vec<usize>) {
    let mut sum_of_secret_numbers = 0;
    for &sn in secret_numbers {
        let (n, _) = evolve(sn);
        sum_of_secret_numbers += n;
    }
    println!("The sum of the secret numbers is {sum_of_secret_numbers}");
}

// fn are_seq_the_same(this_seq: &[(usize, i32)], other_seq: &[(usize, i32)]) -> bool {
//     for (idx, el) in this_seq.into_iter().enumerate() {
//         if el.1 != other_seq[idx].1 {
//             return false;
//         }
//     }
//     true
// }
//
// fn get_most_bananas(prices: &HashMap<usize, Vec<(usize, i32)>>) -> usize {
//     let mut max_bananas = 0;
//     // let mut best_seq = vec![];
//     let mut memo = HashSet::new();
//     for (i, seq_of_changes) in prices.values().enumerate() {
//         println!("Checking seq: {i}");
//         for (_, this_seq) in seq_of_changes.windows(4).enumerate() {
//             let key = this_seq
//                 .into_iter()
//                 .map(|x| format!("{}", x.1))
//                 .collect::<Vec<String>>()
//                 .join(",");
//             if memo.contains(&key) {
//                 continue;
//             }
//             let mut bananas_with_seq = 0;
//             for (_, other_prices) in prices {
//                 for (j, other_seq) in other_prices.windows(4).enumerate() {
//                     if are_seq_the_same(this_seq, other_seq) {
//                         bananas_with_seq += other_prices[j + 3].0;
//                         break;
//                     }
//                 }
//             }
//             if bananas_with_seq > max_bananas {
//                 max_bananas = bananas_with_seq;
//                 // best_seq = Vec::from(this_seq);
//             }
//             memo.insert(key);
//         }
//     }
//     // println!("Best seq: {best_seq:?}");
//     max_bananas
// }

// After re-reading the text and some optimizations :)
fn get_most_bananas_optimized(prices: &HashMap<usize, Vec<(usize, i32)>>) -> usize {
    let mut max_bananas = 0;
    let mut memo = HashMap::new();
    for seq_of_changes in prices.values() {
        let mut already_checked = HashSet::new();
        for this_seq in seq_of_changes.windows(4) {
            let key = this_seq
                .into_iter()
                .map(|x| format!("{}", x.1))
                .collect::<Vec<String>>()
                .join(",");
            // If I find a subsequence twice in the same set of prices, I should
            // not count it twice
            if already_checked.contains(&key) {
                continue;
            }
            already_checked.insert(key.clone());
            let bananas = memo.entry(key).or_insert(0);
            *bananas += this_seq[3].0;
        }
    }
    for (_seq, bananas) in memo {
        if bananas > max_bananas {
            max_bananas = bananas;
            // best_seq = _seq;
        }
    }
    // println!("Best seq: {best_seq:?} --> {max_bananas}");
    max_bananas
}

fn part2(secret_numbers: &Vec<usize>) {
    let mut prices = HashMap::new();
    for &sn in secret_numbers {
        let (_, changes) = evolve(sn);
        prices.insert(sn, changes);
    }
    let bananas = get_most_bananas_optimized(&prices);
    println!("The most bananas I can get is {bananas}");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let secret_numbers = parse_input(&args[1]);
    part1(&secret_numbers);
    part2(&secret_numbers);
}
