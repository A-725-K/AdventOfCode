use cached::proc_macro::cached;
use itertools::Itertools;
use std::{
    collections::{HashSet, VecDeque},
    env,
    fs::read_to_string,
};

fn parse_input(filename: &str) -> Vec<Vec<char>> {
    read_to_string(filename)
        .unwrap()
        .lines()
        .map(|l| l.chars().collect::<Vec<char>>())
        .collect()
}

fn generate_graphs() -> (Vec<Vec<char>>, Vec<Vec<char>>) {
    let numpad = vec![
        vec!['7', '8', '9'],
        vec!['4', '5', '6'],
        vec!['1', '2', '3'],
        vec!['x', '0', 'A'],
    ];
    let directional = vec![vec!['x', '^', 'A'], vec!['<', 'v', '>']];

    (numpad, directional)
}

fn find_shortest(
    graph: &Vec<Vec<char>>,
    start: (i32, i32),
    end: char,
) -> (Vec<String>, (i32, i32)) {
    let rows = graph.len() as i32;
    let cols = graph[0].len() as i32;

    let mut best_paths = HashSet::new();
    let mut min_len = usize::MAX;
    let mut end_x = i32::MAX;
    let mut end_y = i32::MAX;

    let mut q = VecDeque::new();
    q.push_back((start, String::new()));

    while let Some(((curr_x, curr_y), curr_path)) = q.pop_front() {
        // println!("Visiting: ({}, {}) --- p={:?}", curr_x, curr_y, curr_path);
        if curr_x < 0 || curr_y < 0 || curr_x >= cols || curr_y >= rows {
            continue;
        }
        if curr_path.len() > min_len {
            continue;
        }
        if graph[curr_y as usize][curr_x as usize] == 'x' {
            continue;
        }
        if graph[curr_y as usize][curr_x as usize] == end {
            end_x = curr_x;
            end_y = curr_y;
            if curr_path.len() <= min_len {
                min_len = curr_path.len();
                let mut best_path = curr_path.clone();
                best_path.push('A');
                best_paths.insert(best_path);
            }
            continue;
        }

        for ((new_x, new_y), new_dir) in [
            ((curr_x, curr_y + 1), 'v'),
            ((curr_x + 1, curr_y), '>'),
            ((curr_x, curr_y - 1), '^'),
            ((curr_x - 1, curr_y), '<'),
        ] {
            let mut new_path = curr_path.clone();
            new_path.push(new_dir);
            q.push_back(((new_x, new_y), new_path));
        }
    }

    // println!("#best_paths: {}", best_paths.len());
    (best_paths.into_iter().collect_vec(), (end_x, end_y))
}

fn find_all_shortest(graph: &Vec<Vec<char>>, code: &Vec<char>, start: (i32, i32)) -> Vec<String> {
    let mut possibilities = vec![];
    let mut pos = (start.0, start.1);
    for &c in code {
        let (best_paths, end) = find_shortest(graph, pos, c);
        possibilities.push(best_paths);
        pos = (end.0 as i32, end.1 as i32);
    }
    let all_best_paths = possibilities
        .clone()
        .into_iter()
        .multi_cartesian_product()
        .map(|path| path.join(""))
        .collect_vec();

    // println!("paths={all_best_paths:?}");
    all_best_paths
}

fn get_numeric_code(code: &Vec<char>) -> usize {
    code.into_iter().collect::<String>()[0..3]
        .parse::<usize>()
        .unwrap()
}

fn part1(codes: &Vec<Vec<char>>) {
    let (numpad, directional) = generate_graphs();
    let mut sum_of_code_complexity = 0;
    for code in codes {
        let mut all_paths = HashSet::new();
        let mut prev_robot = find_all_shortest(&numpad, code, (2, 3));
        for _ in 0..2 {
            // println!("fst_robot={fst_robot:?}");
            for sequence in prev_robot {
                let snd_robot =
                    find_all_shortest(&directional, &sequence.chars().collect_vec(), (2, 0));
                // println!("snd_robot={snd_robot:?}");
                for snd_seq in snd_robot {
                    all_paths.insert(snd_seq);
                }
            }
            prev_robot = all_paths.into_iter().collect_vec();
            all_paths = HashSet::new();
        }
        let shortest_path_len = prev_robot
            .into_iter()
            .min_by(|a, b| a.len().cmp(&b.len()))
            .unwrap()
            .len();
        // println!("code={code:?}, shortest={shortest_path_len}");
        sum_of_code_complexity += shortest_path_len * get_numeric_code(code);
    }
    println!("The sum of the complexities of the codes is {sum_of_code_complexity}");
}

#[cached]
fn compute_optimal_path_len(seq: String, depth: usize) -> usize {
    let (_, directional) = generate_graphs();
    let mut curr = (2, 0);
    let mut optimal_len = 0;
    if depth == 1 {
        for c in seq.chars() {
            let (bps, end) = find_shortest(&directional, curr, c);
            optimal_len += bps[0].len();
            curr = end;
        }
        return optimal_len;
    }
    for c in seq.chars() {
        let (bps, end) = find_shortest(&directional, curr, c);
        let mut best_lens = vec![];
        for bp in bps {
            best_lens.push(compute_optimal_path_len(bp, depth - 1));
        }
        curr = end;
        optimal_len += best_lens.into_iter().min().unwrap();
    }
    optimal_len
}

fn part2(codes: &Vec<Vec<char>>) {
    let (numpad, _) = generate_graphs();
    let mut sum_of_code_complexity = 0;
    for code in codes {
        let fst_robot = find_all_shortest(&numpad, code, (2, 3));
        let res = fst_robot
            .into_iter()
            .map(|input| compute_optimal_path_len(input, 25))
            .min()
            .unwrap();
        // println!("code={code:?}, shortest={res}");
        sum_of_code_complexity += res * get_numeric_code(code);
    }
    println!("The sum of the complexities of the codes after all the robots finish is {sum_of_code_complexity}");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let lines = parse_input(&args[1]);
    part1(&lines);
    part2(&lines);
}
