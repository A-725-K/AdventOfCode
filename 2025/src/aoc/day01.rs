#[derive(Debug)]
enum Direction {
    Left,
    Right,
}

#[derive(Debug)]
struct Rotation {
    direction: Direction,
    value: i32,
}

fn parse_input(lines: &Vec<String>) -> Vec<Rotation> {
    let mut rotations = vec![];
    for line in lines {
        let direction = if line.chars().nth(0).unwrap() == 'L' {
            Direction::Left
        } else {
            Direction::Right
        };
        let value = line[1..].parse::<i32>().unwrap();
        rotations.push(Rotation { direction, value })
    }
    rotations
}

pub fn part1(lines: &Vec<String>, _day: usize) {
    let mut dial: i32 = 50;
    let rotations = parse_input(lines);
    let mut password = 0;
    for rotation in rotations {
        match rotation.direction {
            Direction::Left => dial = (dial - rotation.value) % 100,
            Direction::Right => dial = (dial + rotation.value) % 100,
        }
        if dial == 0 {
            password += 1;
        }
    }
    println!("The password is {password}");
}

pub fn part2(lines: &Vec<String>, _day: usize) {
    let mut dial: i32 = 50;
    let rotations = parse_input(lines);
    let mut password = 0;
    for rotation in rotations {
        for _ in 0..rotation.value {
            match rotation.direction {
                Direction::Left => dial -= 1,
                Direction::Right => dial += 1,
            }
            dial %= 100;
            if dial == 0 {
                password += 1;
            }
        }
    }
    println!("The password with method 0x434C49434B (CLICK) is {password}");
}
