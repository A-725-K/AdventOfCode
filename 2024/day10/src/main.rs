use std::{collections::HashSet, env, fs::read_to_string};

fn parse_input(filename: &str) -> Vec<Vec<i32>> {
    let mut maze: Vec<Vec<i32>> = vec![];
    for line in read_to_string(filename).unwrap().lines() {
        let mut row = vec![];
        for c in line.chars() {
            if c == '.' {
                row.push(-1);
            } else {
                row.push(c.to_digit(10).unwrap() as i32);
            }
        }
        maze.push(row);
    }
    maze
}

fn count_trailheads(
    maze: &Vec<Vec<i32>>,
    x: i32,
    y: i32,
    current: i32,
    ends: &mut HashSet<(i32, i32)>,
) -> (i32, i32) {
    if current == 9 {
        if ends.contains(&(x, y)) {
            return (0, 1);
        }
        // println!("FOUND --> ({x}, {y})");
        ends.insert((x, y));
        return (1, 1);
    }

    let (mut count, mut count_distinct) = (0, 0);
    let next = current + 1;
    // North
    if y - 1 >= 0 && maze[(y - 1) as usize][x as usize] == next {
        let res = count_trailheads(maze, x, y - 1, next, ends);
        count += res.0;
        count_distinct += res.1;
    }
    // South
    if y + 1 <= (maze.len() - 1) as i32 && maze[(y + 1) as usize][x as usize] == next {
        let res = count_trailheads(maze, x, y + 1, next, ends);
        count += res.0;
        count_distinct += res.1;
    }
    // East
    if x + 1 <= (maze[0].len() - 1) as i32 && maze[y as usize][(x + 1) as usize] == next {
        let res = count_trailheads(maze, x + 1, y, next, ends);
        count += res.0;
        count_distinct += res.1;
    }
    // West
    if x - 1 >= 0 && maze[y as usize][(x - 1) as usize] == next {
        let res = count_trailheads(maze, x - 1, y, next, ends);
        count += res.0;
        count_distinct += res.1;
    }

    (count, count_distinct)
}

fn day_10(maze: &Vec<Vec<i32>>) {
    let mut trailheads = 0;
    let mut distinct_trailheads = 0;
    for (y, row) in maze.into_iter().enumerate() {
        for (x, &start) in row.into_iter().enumerate() {
            if start == 0 {
                let mut ends: HashSet<(i32, i32)> = HashSet::new();
                let counts = count_trailheads(maze, x as i32, y as i32, start, &mut ends);
                trailheads += counts.0;
                distinct_trailheads += counts.1;
            }
        }
    }
    println!("There are {trailheads} trailheads in the topographic map");
    println!("There are {distinct_trailheads} distinct trailheads in the topographic map");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let maze = parse_input(&args[1]);
    day_10(&maze);
}
