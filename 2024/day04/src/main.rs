use std::{env, fs::read_to_string, slice::Iter};

enum Direction {
    N,
    S,
    E,
    W,
    NE,
    NW,
    SE,
    SW,
}

impl Direction {
    pub fn iter() -> Iter<'static, Direction> {
        static DIRECTIONS: [Direction; 8] = [
            Direction::N,
            Direction::S,
            Direction::E,
            Direction::W,
            Direction::NE,
            Direction::NW,
            Direction::SE,
            Direction::SW,
        ];
        DIRECTIONS.iter()
    }
}

fn parse_input(filename: &str) -> Vec<Vec<char>> {
    let mut crosswords: Vec<Vec<char>> = vec![];
    for line in read_to_string(filename).unwrap().lines() {
        let mut l: Vec<char> = vec![];
        for c in line.chars() {
            l.push(c);
        }
        crosswords.push(l);
    }
    crosswords
}

fn search_xmas(
    crosswords: &Vec<Vec<char>>,
    i: i32,
    j: i32,
    rows: i32,
    cols: i32,
    idx: i32,
    dir: &Direction,
) -> bool {
    if i < 0 || i >= rows || j < 0 || j >= cols {
        return false;
    }
    if (idx == 0 && crosswords[i as usize][j as usize] == 'X')
        || (idx == 1 && crosswords[i as usize][j as usize] == 'M')
        || (idx == 2 && crosswords[i as usize][j as usize] == 'A')
    {
        let (new_row, new_col) = match dir {
            Direction::N => (i - 1, j),
            Direction::S => (i + 1, j),
            Direction::E => (i, j + 1),
            Direction::W => (i, j - 1),
            Direction::NE => (i - 1, j + 1),
            Direction::NW => (i - 1, j - 1),
            Direction::SE => (i + 1, j + 1),
            Direction::SW => (i + 1, j - 1),
        };
        return search_xmas(crosswords, new_row, new_col, rows, cols, idx + 1, dir);
    }
    if idx == 3 && crosswords[i as usize][j as usize] == 'S' {
        return true;
    }
    false
}

fn find_xmas(crosswords: &Vec<Vec<char>>, i: i32, j: i32, rows: i32, cols: i32) -> i32 {
    let mut count = 0;
    for direction in Direction::iter() {
        if search_xmas(crosswords, i, j, rows, cols, 0, direction) {
            count += 1;
        }
    }
    count
}

fn part1(crosswords: &Vec<Vec<char>>) {
    let (rows, cols): (i32, i32) = (crosswords.len() as i32, crosswords[0].len() as i32);
    let mut count = 0;
    for i in 0..rows {
        for j in 0..cols {
            count += find_xmas(crosswords, i, j, rows, cols)
        }
    }
    println!("The word XMAS appears {count} times");
}

fn find_x_mas(crosswords: &Vec<Vec<char>>, i: i32, j: i32) -> bool {
    if crosswords[i as usize][j as usize] == 'A' {
        let ne_sw = (crosswords[(i - 1) as usize][(j + 1) as usize] == 'M'
            && crosswords[(i + 1) as usize][(j - 1) as usize] == 'S')
            || (crosswords[(i - 1) as usize][(j + 1) as usize] == 'S'
                && crosswords[(i + 1) as usize][(j - 1) as usize] == 'M');
        let nw_se = (crosswords[(i - 1) as usize][(j - 1) as usize] == 'M'
            && crosswords[(i + 1) as usize][(j + 1) as usize] == 'S')
            || (crosswords[(i - 1) as usize][(j - 1) as usize] == 'S'
                && crosswords[(i + 1) as usize][(j + 1) as usize] == 'M');
        return ne_sw && nw_se;
    }
    false
}

fn part2(crosswords: &Vec<Vec<char>>) {
    let (rows, cols): (i32, i32) = (crosswords.len() as i32, crosswords[0].len() as i32);
    let mut count = 0;
    for i in 1..rows - 1 {
        for j in 1..cols - 1 {
            if find_x_mas(crosswords, i, j) {
                count += 1;
            }
        }
    }
    println!("There are {count} X-MAS shapes");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let crosswords = parse_input(&args[1]);
    part1(&crosswords);
    part2(&crosswords);
}
