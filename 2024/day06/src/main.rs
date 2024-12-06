use std::{collections::HashSet, env, fmt, fs::read_to_string};

#[derive(PartialEq, Clone, Copy, Debug)]
enum Direction {
    N,
    S,
    E,
    W,
}

impl fmt::Display for Direction {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{:?}", self)
    }
}

fn parse_input(filename: &str) -> (Vec<Vec<char>>, i32, i32) {
    let mut maze: Vec<Vec<char>> = vec![];
    let (mut start_x, mut start_y): (i32, i32) = (0, 0);
    for (y, line) in read_to_string(filename)
        .unwrap()
        .lines()
        .into_iter()
        .enumerate()
    {
        let mut l: Vec<char> = vec![];
        for (x, c) in line.chars().into_iter().enumerate() {
            if c == '^' {
                start_x = x as i32;
                start_y = y as i32;
            }
            l.push(c);
        }
        maze.push(l);
    }
    (maze, start_x, start_y)
}

fn get_next_direction(curr: Direction) -> Direction {
    match curr {
        Direction::N => Direction::E,
        Direction::E => Direction::S,
        Direction::S => Direction::W,
        Direction::W => Direction::N,
    }
}

fn can_move(
    maze: &Vec<Vec<char>>,
    x: &i32,
    y: &i32,
    rows: &i32,
    cols: &i32,
    direction: &Direction,
) -> bool {
    if *direction == Direction::N {
        if *y - 1 < 0 || maze[(*y - 1) as usize][*x as usize] != '#' {
            return true;
        }
        return false;
    }
    if *direction == Direction::E {
        if *x + 1 >= *cols || maze[*y as usize][(*x + 1) as usize] != '#' {
            return true;
        }
        return false;
    }
    if *direction == Direction::S {
        if *y + 1 >= *rows || maze[(*y + 1) as usize][*x as usize] != '#' {
            return true;
        }
        return false;
    }
    // Direction::W
    if *x - 1 < 0 || maze[*y as usize][(*x - 1) as usize] != '#' {
        return true;
    }
    return false;
}

fn move_in_direction(x: i32, y: i32, curr: Direction) -> (i32, i32) {
    match curr {
        Direction::N => (x, y - 1),
        Direction::E => (x + 1, y),
        Direction::S => (x, y + 1),
        Direction::W => (x - 1, y),
    }
}

fn part1(maze: &Vec<Vec<char>>, mut x: i32, mut y: i32, rows: &i32, cols: &i32) -> Vec<Vec<bool>> {
    let mut visited: Vec<Vec<bool>> = vec![];
    for _ in 0..*rows {
        let mut row: Vec<bool> = vec![];
        for _ in 0..*cols {
            row.push(false);
        }
        visited.push(row);
    }

    let mut current_direction = Direction::N;

    while x >= 0 && y >= 0 && x < *rows && y < *rows {
        visited[y as usize][x as usize] = true;
        if can_move(maze, &x, &y, rows, cols, &current_direction) {
            (x, y) = move_in_direction(x, y, current_direction);
        } else {
            current_direction = get_next_direction(current_direction);
        }
    }

    let mut tiles_visited = 0;
    for i in 0..*rows {
        for j in 0..*cols {
            if visited[i as usize][j as usize] {
                tiles_visited += 1;
            }
        }
    }
    println!("The guard visited {tiles_visited} tiles");
    visited
}

fn try_find_loop(maze: &Vec<Vec<char>>, mut x: i32, mut y: i32, rows: &i32, cols: &i32) -> bool {
    let mut visited: HashSet<String> = HashSet::new();
    let mut current_direction = Direction::N;

    while x >= 0 && y >= 0 && x < *rows && y < *rows {
        let k = format!("{}-{}-{}", x, y, current_direction);
        if let Some(_) = visited.get(&k) {
            return true;
        }
        visited.insert(k);
        if can_move(maze, &x, &y, rows, cols, &current_direction) {
            (x, y) = move_in_direction(x, y, current_direction);
        } else {
            current_direction = get_next_direction(current_direction);
        }
    }

    false
}

fn part2(
    mut maze: Vec<Vec<char>>,
    start_x: i32,
    start_y: i32,
    rows: &i32,
    cols: &i32,
    visited: &Vec<Vec<bool>>,
) {
    let mut possible_obstacles = 0;

    // Observation: the block must appear in one of the tiles that the guard
    // visit during their turn. It is pointless trying to put it in a tile that
    // is not touched, since `find_loop` will always return false in these case,
    // therefore we can skip considering them.
    // With this optimization I went from ~16sec to ~3.5sec.
    for y in 0..*rows {
        for x in 0..*cols {
            if visited[y as usize][x as usize] // this optimization is not "flour of my bag ;)"
                && maze[y as usize][x as usize] != '#'
            {
                let prev = maze[y as usize][x as usize];
                maze[y as usize][x as usize] = '#';
                let (xx, yy): (i32, i32) = (start_x, start_y);
                if try_find_loop(&maze, xx, yy, rows, cols) {
                    possible_obstacles += 1;
                }
                maze[y as usize][x as usize] = prev;
            }
        }
    }
    println!("We can position {possible_obstacles} to trick the guards");
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let (maze, start_x, start_y) = parse_input(&args[1]);
    let (rows, cols): (i32, i32) = (maze.len() as i32, maze[0].len() as i32);
    let (sx, sy) = (start_x, start_y);
    let visited = part1(&maze, start_x, start_y, &rows, &cols);
    part2(maze, sx, sy, &rows, &cols, &visited);
}
