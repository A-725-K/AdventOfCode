use cached::proc_macro::cached;

fn parse_input(lines: &Vec<String>) -> (Vec<Vec<char>>, usize, usize) {
    let result: Vec<_> = lines.iter().map(|line| line.chars().collect()).collect();
    (result.clone(), result.len(), result[0].len())
}

pub fn part1(lines: &Vec<String>, _day: usize) {
    let (mut manyfold, rows, cols) = parse_input(lines);
    let mut beam_splits = 0;

    for row in 1..rows {
        for col in 0..cols {
            if manyfold[row][col] == '.'
                && (manyfold[row - 1][col] == 'S' || manyfold[row - 1][col] == '|')
            {
                manyfold[row][col] = '|';
            } else if manyfold[row][col] == '^' && manyfold[row - 1][col] == '|' {
                // assuming all rays are within the borders
                manyfold[row][col - 1] = '|';
                manyfold[row][col + 1] = '|';
                beam_splits += 1;
            }
        }
    }

    println!("The beam has been split {beam_splits} times");
}

// Backtrack algorightm... it works but it is too slow... even if cached... :(
// fn visit(manyfold: &mut Vec<Vec<char>>, row: usize, rows: usize, cols: usize) -> usize {
//     if row >= rows {
//         let mut beam_arrived = false;
//         for col in 0..cols {
//             if manyfold[rows - 1][col] == '|' {
//                 beam_arrived = true;
//                 break;
//             }
//         }
//         if beam_arrived {
//             // for i in 0..rows {
//             //     for j in 0..cols {
//             //         print!("{}", manyfold[i][j]);
//             //     }
//             //     println!();
//             // }
//             // println!();
//             return 1;
//         }
//         return 0;
//     }
//     let mut timelines = 0;
//     for col in 0..cols {
//         if manyfold[row][col] == '.'
//             && (manyfold[row - 1][col] == 'S' || manyfold[row - 1][col] == '|')
//         {
//             manyfold[row][col] = '|';
//         } else if manyfold[row][col] == '^' && manyfold[row - 1][col] == '|' {
//             manyfold[row][col - 1] = '|';
//             timelines += visit(&mut manyfold.clone(), row + 1, rows, cols);
//             manyfold[row][col - 1] = '.';
//             manyfold[row][col + 1] = '|';
//             timelines += visit(&mut manyfold.clone(), row + 1, rows, cols);
//             manyfold[row][col + 1] = '.';
//         }
//     }
//     timelines + visit(manyfold, row + 1, rows, cols)
// }

#[cached]
fn visit(manyfold: Vec<Vec<char>>, row: usize, col: usize, rows: usize, cols: usize) -> usize {
    // the beam has arrived
    if row >= rows {
        return 1;
    }

    // propagate the current beam
    if manyfold[row][col] == '.' || manyfold[row][col] == 'S' {
        return visit(manyfold, row + 1, col, rows, cols);
    }

    // found a split ^: need to count the results on the left and the right and sum them
    visit(manyfold.clone(), row, col - 1, rows, cols) + visit(manyfold, row, col + 1, rows, cols)
}

pub fn part2(lines: &Vec<String>, _day: usize) {
    let (manyfold, rows, cols) = parse_input(lines);
    let start_col = manyfold[0].iter().position(|&c| c == 'S').unwrap();
    let timelines = visit(manyfold, 0, start_col, rows, cols);
    println!("There are {timelines} possible timelines");
}
