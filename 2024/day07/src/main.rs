use itertools::Itertools;
use std::{env, fs::read_to_string};

// ===================== BEGINNING OF ROSETTA CODE ==================== //
// struct PermutationIterator<'a, T: 'a> {
//     universe: &'a [T],
//     size: usize,
//     prev: Option<Vec<usize>>,
// }
//
// fn permutations<T>(universe: &[T], size: usize) -> PermutationIterator<T> {
//     PermutationIterator {
//         universe,
//         size,
//         prev: None,
//     }
// }
//
// fn map<T>(values: &[T], ixs: &[usize]) -> Vec<T>
// where
//     T: Clone,
// {
//     ixs.iter().map(|&i| values[i].clone()).collect()
// }
//
// impl<'a, T> Iterator for PermutationIterator<'a, T>
// where
//     T: Clone,
// {
//     type Item = Vec<T>;
//
//     fn next(&mut self) -> Option<Vec<T>> {
//         let n = self.universe.len();
//
//         if n == 0 {
//             return None;
//         }
//
//         match self.prev {
//             None => {
//                 let zeroes: Vec<usize> = std::iter::repeat(0).take(self.size).collect();
//                 let result = Some(map(self.universe, &zeroes[..]));
//                 self.prev = Some(zeroes);
//                 result
//             }
//             Some(ref mut indexes) => match indexes.iter().position(|&i| i + 1 < n) {
//                 None => None,
//                 Some(position) => {
//                     for index in indexes.iter_mut().take(position) {
//                         *index = 0;
//                     }
//                     indexes[position] += 1;
//                     Some(map(self.universe, &indexes[..]))
//                 }
//             },
//         }
//     }
// }
// ===================== END OF ROSETTA CODE ==================== //

fn parse_input(filename: &str) -> Vec<(i64, Vec<i64>)> {
    let mut equations: Vec<(i64, Vec<i64>)> = vec![];
    for line in read_to_string(filename).unwrap().lines() {
        let fields: Vec<&str> = line.split(':').collect();
        let key = fields[0].parse::<i64>().unwrap();
        let value: Vec<i64> = fields[1]
            .split_whitespace()
            .map(|n| n.parse::<i64>().unwrap())
            .collect();
        equations.push((key, value));
    }
    equations
}

fn is_valid_equation(result: i64, numbers: &Vec<i64>, all: bool) -> bool {
    let n = numbers.len() - 1;
    let operators = if all {
        vec!["+", "*", "||"]
    } else {
        vec!["+", "*"]
    };

    let seqs = (1..=n).map(|_| &operators[..]).multi_cartesian_product();
    for seq in seqs {
        let mut attempt: i64 = numbers[0];
        for (i, &op) in seq.into_iter().enumerate() {
            if op == "+" {
                attempt += numbers[i + 1];
            } else if op == "*" {
                attempt *= numbers[i + 1];
            } else {
                attempt = format!("{attempt}{}", numbers[i + 1])
                    .parse::<i64>()
                    .unwrap();
            }
        }
        if attempt == result {
            return true;
        }
    }
    false
}

fn part1(equations: &Vec<(i64, Vec<i64>)>) {
    let mut sum_of_valid: i64 = 0;
    for (result, numbers) in equations.into_iter() {
        if is_valid_equation(*result, &numbers, false) {
            sum_of_valid += result;
        }
    }
    println!("The total calibration result is {sum_of_valid}");
}

fn part2(equations: &Vec<(i64, Vec<i64>)>) {
    let mut sum_of_valid: i64 = 0;
    for (result, numbers) in equations.into_iter() {
        if is_valid_equation(*result, &numbers, true) {
            sum_of_valid += result;
        }
    }
    println!("The total calibration result after adding missing operator is {sum_of_valid}");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let lines = parse_input(&args[1]);
    part1(&lines);
    part2(&lines);
}
