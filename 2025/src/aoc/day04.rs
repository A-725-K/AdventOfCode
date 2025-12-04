fn parse_input(lines: &Vec<String>) -> Vec<Vec<char>> {
    lines.iter().map(|line| line.chars().collect()).collect()
}

fn can_be_accessed_by_forklift(diagram: &Vec<Vec<char>>, row: i32, col: i32) -> bool {
    let rows = diagram.len() as i32;
    let cols = diagram[0].len() as i32;

    let mut paper_rolls = 0;
    for r in row - 1..row + 2 {
        if r < 0 || r >= rows {
            continue;
        }
        for c in col - 1..col + 2 {
            if c < 0 || c >= cols || (r == row && c == col) {
                continue;
            }
            if diagram[r as usize][c as usize] == '@' {
                paper_rolls += 1;
            }
        }
    }

    paper_rolls < 4
}

pub fn part1(lines: &Vec<String>, _day: usize) {
    let diagram = parse_input(lines);
    let mut paper_rolls_accessible_by_forklift = 0;
    for (row, row_elems) in diagram.iter().enumerate() {
        for (col, &elem) in row_elems.iter().enumerate() {
            if elem == '@' && can_be_accessed_by_forklift(&diagram, row as i32, col as i32) {
                paper_rolls_accessible_by_forklift += 1;
            }
        }
    }
    println!(
        "There are {paper_rolls_accessible_by_forklift} paper rolls that can be accessed by the forklift"
    );
}

pub fn part2(lines: &Vec<String>, _day: usize) {
    let mut diagram = parse_input(lines);
    let mut paper_rolls_accessible_by_forklift = 0;

    loop {
        let mut to_remove = vec![];
        for (row, row_elems) in diagram.iter().enumerate() {
            for (col, &elem) in row_elems.iter().enumerate() {
                if elem == '@' && can_be_accessed_by_forklift(&diagram, row as i32, col as i32) {
                    to_remove.push((row, col));
                }
            }
        }
        if to_remove.is_empty() {
            break;
        }

        paper_rolls_accessible_by_forklift += to_remove.len();
        for (row, col) in to_remove {
            diagram[row][col] = '.';
        }
    }

    println!("There are {paper_rolls_accessible_by_forklift} that can be removed using forklifts");
}
