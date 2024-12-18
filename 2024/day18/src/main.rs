use std::{
    collections::{HashSet, VecDeque},
    env,
    fs::read_to_string,
};

// const ROWS: i32 = 7;
// const COLS: i32 = 7;
// const BYTES: usize = 12;

const ROWS: i32 = 71;
const COLS: i32 = 71;
const BYTES: usize = 1024;

fn parse_input(filename: &str) -> Vec<(i32, i32)> {
    read_to_string(filename)
        .unwrap()
        .lines()
        .map(|l| {
            (
                l.split(",").collect::<Vec<&str>>()[0]
                    .parse::<i32>()
                    .unwrap(),
                l.split(",").collect::<Vec<&str>>()[1]
                    .parse::<i32>()
                    .unwrap(),
            )
        })
        .collect()
}

fn pathfinder(bytes: &Vec<(i32, i32)>, end: usize) -> usize {
    let mut visited = vec![];
    for _ in 0..ROWS {
        let mut row = vec![];
        for _ in 0..COLS {
            row.push(false);
        }
        visited.push(row);
    }

    let mut q = VecDeque::new();
    q.push_back((0 as i32, 0 as i32, 0 as usize));

    for &(x, y) in &bytes[..end] {
        visited[y as usize][x as usize] = true;
    }

    let mut memo = HashSet::new();
    let mut min_path_len = usize::MAX;
    while let Some((x, y, path_len)) = q.pop_front() {
        // println!("visiting: ({x}, {y}) p={path_len} #q={}", q.len());
        if x == COLS - 1 && y == ROWS - 1 {
            if path_len < min_path_len {
                min_path_len = path_len;
            }
            continue;
        }
        if path_len >= min_path_len {
            continue;
        }
        if memo.contains(&(x, y, path_len)) {
            continue;
        }

        memo.insert((x, y, path_len));
        visited[y as usize][x as usize] = true;

        // North
        if y - 1 >= 0 && !visited[(y - 1) as usize][x as usize] {
            q.push_back((x, y - 1, path_len + 1));
        }
        // South
        if y + 1 < ROWS && !visited[(y + 1) as usize][x as usize] {
            q.push_back((x, y + 1, path_len + 1));
        }
        // East
        if x + 1 < COLS && !visited[y as usize][(x + 1) as usize] {
            q.push_back((x + 1, y, path_len + 1));
        }
        // West
        if x - 1 >= 0 && !visited[y as usize][(x - 1) as usize] {
            q.push_back((x - 1, y, path_len + 1));
        }
    }
    min_path_len
}

fn part1(bytes: &Vec<(i32, i32)>) {
    let min_path_len = pathfinder(bytes, BYTES);
    println!("The minimum number of steps to reach the exit is {min_path_len}");
}

fn part2(bytes: &Vec<(i32, i32)>) {
    for (i, _) in bytes.into_iter().enumerate() {
        // println!("Checking the byte ({i})...");
        let min_path_len = pathfinder(bytes, i + 1);
        if min_path_len == usize::MAX {
            println!(
                "The byte that blocks everything is ({}, {})",
                bytes[i].0, bytes[i].1
            );
            break;
        }
    }
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let bytes = parse_input(&args[1]);
    part1(&bytes);
    part2(&bytes);
}
