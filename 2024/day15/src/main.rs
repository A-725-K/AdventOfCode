use std::{env, fs::read_to_string};

// TODO: implement animation :)

type Warehouse = Vec<Vec<char>>;

fn print_warehouse(warehouse: &Warehouse) {
    for y in 0..warehouse.len() {
        for x in 0..warehouse[0].len() {
            print!("{}", warehouse[y][x]);
        }
        println!("");
    }
}

fn parse_input(filename: &str) -> (Warehouse, Vec<char>, (usize, usize)) {
    let mut warehouse: Warehouse = vec![];
    let mut read_warehouse = true;
    let mut movements = vec![];
    let mut robot_pos = (0, 0);

    for (y, line) in read_to_string(filename).unwrap().lines().enumerate() {
        if line == "" {
            read_warehouse = false;
            continue;
        }
        if read_warehouse {
            let row: Vec<char> = line.chars().collect();
            if let Some(x) = row.iter().position(|&c| c == '@') {
                robot_pos = (x, y);
            }
            warehouse.push(row);
        } else {
            line.chars().for_each(|m| movements.push(m));
        }
    }
    (warehouse, movements, robot_pos)
}

fn can_move_in_direction(warehouse: &mut Warehouse, x: usize, y: usize, dir: char) -> bool {
    if warehouse[y][x] == '#' {
        return false;
    }
    if warehouse[y][x] == '.' {
        return true;
    }
    let mut new_dir = (x, y);
    match dir {
        '^' => new_dir.1 -= 1,
        '>' => new_dir.0 += 1,
        '<' => new_dir.0 -= 1,
        'v' => new_dir.1 += 1,
        _ => {}
    }
    if can_move_in_direction(warehouse, new_dir.0, new_dir.1, dir) {
        warehouse[new_dir.1][new_dir.0] = warehouse[y][x];
        return true;
    }
    return false;
}

fn compute_gps_cost(warehouse: &Warehouse) -> usize {
    let mut gps_coords_sum = 0;
    for y in 0..warehouse.len() {
        for x in 0..warehouse[0].len() {
            if warehouse[y][x] == 'O' || warehouse[y][x] == '[' {
                gps_coords_sum += x + y * 100;
            }
        }
    }
    gps_coords_sum
}

fn part1(
    warehouse: &mut Warehouse,
    movements: &Vec<char>,
    robot: &mut (usize, usize),
    debug: bool,
) {
    if debug {
        println!("--------- BEFORE ----------");
        print_warehouse(warehouse);
    }
    for &m in movements {
        warehouse[robot.1][robot.0] = '.';
        if m == '^' {
            if can_move_in_direction(warehouse, robot.0, robot.1 - 1, m) {
                robot.1 -= 1;
            }
        } else if m == '>' {
            if can_move_in_direction(warehouse, robot.0 + 1, robot.1, m) {
                robot.0 += 1;
            }
        } else if m == '<' {
            if can_move_in_direction(warehouse, robot.0 - 1, robot.1, m) {
                robot.0 -= 1;
            }
        } else if m == 'v' {
            if can_move_in_direction(warehouse, robot.0, robot.1 + 1, m) {
                robot.1 += 1;
            }
        } else {
            panic!("Unknown movement!");
        }
        warehouse[robot.1][robot.0] = '@';
    }
    if debug {
        println!("--------- AFTER ----------");
        print_warehouse(warehouse);
    }

    let gps_coords_sum = compute_gps_cost(warehouse);
    println!("After the robot mess the sum of GPS of boxes is {gps_coords_sum}");
}

fn generate_new_warehouse(warehouse: &Warehouse) -> (Warehouse, (usize, usize)) {
    let mut new_warehouse = vec![];
    let mut robot_pos = (0, 0);
    for y in 0..warehouse.len() {
        let mut row = vec![];
        for x in 0..warehouse[0].len() {
            match warehouse[y][x] {
                '#' => {
                    row.push('#');
                    row.push('#');
                }
                'O' => {
                    row.push('[');
                    row.push(']');
                }
                '.' => {
                    row.push('.');
                    row.push('.');
                }
                '@' => {
                    row.push('@');
                    row.push('.');
                    robot_pos = (x * 2, y);
                }
                _ => {}
            }
        }
        new_warehouse.push(row);
    }
    (new_warehouse, robot_pos)
}

// Left and right movement can be handled in the same way as before (in the last
// else-if statement of the function). For up and down movements we need extra
// care.
fn can_move_in_direction_big(warehouse: &mut Warehouse, x: usize, y: usize, dir: char) -> bool {
    if warehouse[y][x] == '#' {
        return false;
    }
    if warehouse[y][x] == '.' {
        return true;
    }
    let mut new_dir = (x, y);
    match dir {
        '^' => new_dir.1 -= 1,
        '>' => new_dir.0 += 1,
        '<' => new_dir.0 -= 1,
        'v' => new_dir.1 += 1,
        _ => {}
    }
    if warehouse[y][x] == '[' && (dir == '^' || dir == 'v') {
        let left = can_move_in_direction_big(warehouse, new_dir.0, new_dir.1, dir);
        let right = can_move_in_direction_big(warehouse, new_dir.0 + 1, new_dir.1, dir);
        if left && right {
            warehouse[new_dir.1][new_dir.0] = warehouse[y][x];
            warehouse[new_dir.1][new_dir.0 + 1] = warehouse[y][x + 1];
            warehouse[y][x + 1] = '.';
            return true;
        }
    } else if warehouse[y][x] == ']' && (dir == '^' || dir == 'v') {
        let left = can_move_in_direction_big(warehouse, x - 1, new_dir.1, dir);
        let right = can_move_in_direction_big(warehouse, new_dir.0, new_dir.1, dir);
        if left && right {
            warehouse[new_dir.1][new_dir.0] = warehouse[y][x];
            warehouse[new_dir.1][new_dir.0 - 1] = warehouse[y][x - 1];
            warehouse[y][x - 1] = '.';
            return true;
        }
    } else if can_move_in_direction_big(warehouse, new_dir.0, new_dir.1, dir) {
        warehouse[new_dir.1][new_dir.0] = warehouse[y][x];
        return true;
    }
    return false;
}

fn part2(warehouse: &Warehouse, movements: &Vec<char>, debug: bool) {
    let (mut big_warehouse, mut robot) = generate_new_warehouse(warehouse);

    if debug {
        println!("--------- BEFORE ----------");
        print_warehouse(&big_warehouse);
    }
    for &m in movements {
        big_warehouse[robot.1][robot.0] = '.';
        // It is necessary to use backtrack strategy to handle the following
        // situation:
        //
        // ################
        // ##........######
        // ##......[][]..## [2] [3]
        // ##.......[]...##   [1]
        // ##.......@....##
        // ################
        // MOVEMENT: ^
        //
        // in this case `can_move_in_direction_big` will return true for both
        // left and right case for (2) causing it to move towards the top of the
        // warehouse. Unfortunately (3) cannot move up, therefore the whole
        // block 1-2-3 cannot move up but (2) has been already lifted up.
        // The whole computation, though, returns false, therefore at this level
        // we know that we moved something when we shouldn't have. The easiest
        // way to undo the wrong movement is to restore to the previous situation
        // since we know that there has been no movement at all. Probably not
        // the most efficient solution, but it works!
        let backtrack_copy = big_warehouse.clone();
        if m == '^' {
            if can_move_in_direction_big(&mut big_warehouse, robot.0, robot.1 - 1, m) {
                robot.1 -= 1;
            } else {
                big_warehouse = backtrack_copy;
            }
        } else if m == '>' {
            if can_move_in_direction_big(&mut big_warehouse, robot.0 + 1, robot.1, m) {
                robot.0 += 1;
            }
        } else if m == '<' {
            if can_move_in_direction_big(&mut big_warehouse, robot.0 - 1, robot.1, m) {
                robot.0 -= 1;
            }
        } else if m == 'v' {
            if can_move_in_direction_big(&mut big_warehouse, robot.0, robot.1 + 1, m) {
                robot.1 += 1;
            } else {
                big_warehouse = backtrack_copy;
            }
        } else {
            panic!("Unknown movement!");
        }
        big_warehouse[robot.1][robot.0] = '@';
    }
    if debug {
        println!("--------- AFTER ----------");
        print_warehouse(&big_warehouse);
    }

    let gps_coords_sum = compute_gps_cost(&big_warehouse);
    println!(
        "After the robot mess in the big warehouse the sum of GPS of boxes is {gps_coords_sum}"
    );
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let (warehouse, movements, mut robot_pos) = parse_input(&args[1]);
    let debug = false;
    part1(
        warehouse.clone().as_mut(),
        &movements,
        &mut robot_pos,
        debug,
    );
    part2(&warehouse, &movements, debug);
}
