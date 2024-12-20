use std::{
    collections::{HashMap, HashSet, VecDeque},
    env,
    fs::read_to_string,
};

fn parse_input(filename: &str) -> (Vec<Vec<char>>, (usize, usize)) {
    let maze = read_to_string(filename)
        .unwrap()
        .lines()
        .map(|l| l.chars().collect())
        .collect::<Vec<Vec<char>>>();
    let mut start = (0, 0);
    'out: for y in 0..maze.len() {
        for x in 0..maze[0].len() {
            if maze[y][x] == 'S' {
                start = (x, y);
                break 'out;
            }
        }
    }
    (maze, start)
}

// fn generate_visited(rows: usize, cols: usize) -> Vec<Vec<bool>> {
//     let mut visited = vec![];
//     for _ in 0..rows {
//         let mut row = vec![];
//         for _ in 0..cols {
//             row.push(false);
//         }
//         visited.push(row);
//     }
//     visited
// }
//
// fn walk_normal(
//     maze: &Vec<Vec<char>>,
//     x: usize,
//     y: usize,
//     path_len: usize,
//     visited: &mut Vec<Vec<bool>>,
// ) -> usize {
//     if maze[y][x] == 'E' {
//         return path_len;
//     }
//
//     visited[y][x] = true;
//
//     // North
//     if maze[y - 1][x] != '#' && !visited[y - 1][x] {
//         return walk_normal(maze, x, y - 1, path_len + 1, visited);
//     }
//     // South
//     if maze[y + 1][x] != '#' && !visited[y + 1][x] {
//         return walk_normal(maze, x, y + 1, path_len + 1, visited);
//     }
//     // East
//     if maze[y][x + 1] != '#' && !visited[y][x + 1] {
//         return walk_normal(maze, x + 1, y, path_len + 1, visited);
//     }
//     // West
//     if maze[y][x - 1] != '#' && !visited[y][x - 1] {
//         return walk_normal(maze, x - 1, y, path_len + 1, visited);
//     }
//
//     0
// }
//
// fn try_solve_part_1(maze: &Vec<Vec<char>>, start: (usize, usize)) {
//     let mut visited = generate_visited(maze.len(), maze[0].len());
//     let min_path_len = walk_normal(maze, start.0, start.1, 0, &mut visited);
//     visited = generate_visited(maze.len(), maze[0].len());
//     let mut cheats = HashMap::new();
//     let mut q = VecDeque::new();
//     q.push_back((start.0 as i32, start.1 as i32, 0, false));
//
//     while let Some((x, y, path_len, has_cheated)) = q.pop_front() {
//         if path_len >= min_path_len {
//             continue;
//         }
//         if maze[y as usize][x as usize] == 'E' {
//             let diff = min_path_len - path_len;
//             println!(">>>>> FOUND A PATH THAT SAVES {diff} PS");
//             let curr_cheat = cheats.entry(diff).or_insert(0);
//             *curr_cheat += 1;
//             continue;
//         }
//         if visited[y as usize][x as usize] {
//             continue;
//         }
//
//         visited[y as usize][x as usize] = true;
//
//         // North
//         if maze[(y - 1) as usize][x as usize] != '#' {
//             //&& !visited[(y - 1) as usize][x as usize] {
//             q.push_back((x, y - 1, path_len + 1, has_cheated));
//         }
//         if maze[(y - 1) as usize][x as usize] == '#'
//             && !has_cheated
//             && y - 2 >= 0
//             && maze[(y - 2) as usize][x as usize] != '#'
//             && !visited[(y - 2) as usize][x as usize]
//         {
//             q.push_back((x, y - 2, path_len + 2, true));
//         }
//         // South
//         if maze[(y + 1) as usize][x as usize] != '#' {
//             //&& !visited[(y + 1) as usize][x as usize] {
//             q.push_back((x, y + 1, path_len + 1, has_cheated));
//         }
//         if maze[(y + 1) as usize][x as usize] == '#'
//             && !has_cheated
//             && y + 2 < maze.len() as i32
//             && maze[(y + 2) as usize][x as usize] != '#'
//             && !visited[(y + 2) as usize][x as usize]
//         {
//             q.push_back((x, y + 2, path_len + 2, true));
//         }
//         // East
//         if maze[y as usize][(x + 1) as usize] != '#' {
//             //&& !visited[y as usize][(x + 1) as usize] {
//             q.push_back((x + 1, y, path_len + 1, has_cheated));
//         }
//         if maze[y as usize][(x + 1) as usize] == '#'
//             && !has_cheated
//             && x + 2 < maze[0].len() as i32
//             && maze[y as usize][(x + 2) as usize] != '#'
//             && !visited[y as usize][(x + 2) as usize]
//         {
//             q.push_back((x + 2, y, path_len + 2, true));
//         }
//         // West
//         if maze[y as usize][(x - 1) as usize] != '#' {
//             // && !visited[y as usize][(x - 1) as usize] {
//             q.push_back((x - 1, y, path_len + 1, has_cheated));
//         }
//         if maze[y as usize][(x - 1) as usize] == '#'
//             && !has_cheated
//             && x - 2 >= 0
//             && maze[y as usize][(x - 2) as usize] != '#'
//             && !visited[y as usize][(x - 2) as usize]
//         {
//             q.push_back((x - 2, y, path_len + 2, true));
//         }
//         visited[y as usize][x as usize] = false;
//
//         // println!("q={}", q.len());
//     }
//
//     let mut keys: Vec<&usize> = cheats.keys().collect();
//     keys.sort();
//     for k in keys {
//         println!("{k}: {}", cheats.get(&k).unwrap());
//     }
// }

fn compute_distances(
    maze: &Vec<Vec<char>>,
    start: (usize, usize),
    rows: i32,
    cols: i32,
) -> Vec<Vec<i32>> {
    let (mut x, mut y) = (start.0 as i32, start.1 as i32);
    let mut distances = (0..rows)
        .map(|_| (0..cols).map(|_| -1).collect::<Vec<i32>>())
        .collect::<Vec<Vec<i32>>>();
    distances[y as usize][x as usize] = 0;
    while maze[y as usize][x as usize] != 'E' {
        for (new_x, new_y) in [(x, y - 1), (x, y + 1), (x + 1, y), (x - 1, y)] {
            if new_x < 0 || new_y < 0 || new_x >= cols || new_y >= rows {
                continue;
            }
            if maze[new_y as usize][new_x as usize] == '#' {
                continue;
            }
            if distances[new_y as usize][new_x as usize] >= 0 {
                continue;
            }
            distances[new_y as usize][new_x as usize] = distances[y as usize][x as usize] + 1;
            x = new_x;
            y = new_y;
        }
    }
    distances
}

// https://www.youtube.com/watch?v=tWhwcORztSY
fn part1(maze: &Vec<Vec<char>>, start: (usize, usize)) {
    let rows = maze.len() as i32;
    let cols = maze[0].len() as i32;
    const THRESHOLD: i32 = 100;

    // First pass: compute the distances for each tile
    let distances = compute_distances(maze, start, rows, cols);

    let mut cheats_count = 0;
    // Second pass: try to cheat
    for y in 0..rows {
        for x in 0..cols {
            if maze[y as usize][x as usize] == '#' {
                continue;
            }
            for (new_x, new_y) in [
                (x, y - 2),
                (x, y + 2),
                (x + 2, y),
                (x - 2, y),
                (x + 1, y + 1),
                (x + 1, y - 1),
                (x - 1, y + 1),
                (x - 1, y - 1),
            ] {
                if new_x < 0 || new_y < 0 || new_x >= cols || new_y >= rows {
                    continue;
                }
                // Avoid segmentation faults
                if maze[new_y as usize][new_x as usize] == '#' {
                    continue;
                }
                if distances[y as usize][x as usize] - distances[new_y as usize][new_x as usize]
                    >= THRESHOLD + 2
                // 2 is the time allowed to cheat
                {
                    cheats_count += 1;
                }
            }
        }
    }

    // println!("{:?}", distances);
    println!("There are {cheats_count} cheats that can save you {THRESHOLD} picoseconds");
}

fn part2(maze: &Vec<Vec<char>>, start: (usize, usize)) {
    let rows = maze.len() as i32;
    let cols = maze[0].len() as i32;
    const THRESHOLD: i32 = 100;

    // First pass: compute the distances for each tile
    let distances = compute_distances(maze, start, rows, cols);

    let mut cheats_count = 0;
    // Second pass: try to cheat
    for y in 0..rows {
        for x in 0..cols {
            if maze[y as usize][x as usize] == '#' {
                continue;
            }
            for radius in 2..=20 {
                for delta_row in 0..=radius {
                    let delta_col = radius - delta_row;
                    // Using a set because otherwise some points are visited
                    // twice.
                    for (new_x, new_y) in HashSet::from([
                        (x + delta_col, y + delta_row),
                        (x + delta_col, y - delta_row),
                        (x - delta_col, y + delta_row),
                        (x - delta_col, y - delta_row),
                    ]) {
                        if new_x < 0 || new_y < 0 || new_x >= cols || new_y >= rows {
                            continue;
                        }
                        // Avoid segmentation faults
                        if maze[new_y as usize][new_x as usize] == '#' {
                            continue;
                        }
                        if distances[y as usize][x as usize]
                            - distances[new_y as usize][new_x as usize]
                            >= THRESHOLD + radius
                        // 2 is the time allowed to cheat
                        {
                            cheats_count += 1;
                        }
                    }
                }
            }
        }
    }
    println!("There are {cheats_count} cheats that lasts up to 20 picseconds that can save you {THRESHOLD} picoseconds");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let (maze, start) = parse_input(&args[1]);
    part1(&maze, start);
    part2(&maze, start);
}
