use std::{
    collections::{BinaryHeap, HashMap, HashSet, VecDeque},
    env,
    fs::read_to_string,
};

// #[derive(Debug, PartialEq, Eq, Clone, Copy, Hash, PartialOrd, Ord)]
// enum Direction {
//     N,
//     S,
//     E,
//     W,
// }

type Maze = Vec<Vec<char>>;

fn parse_input(filename: &str) -> Maze {
    read_to_string(filename)
        .unwrap()
        .lines()
        .map(|l| l.chars().collect())
        .collect()
}

// fn print_visited(maze: &Maze, visited: &HashSet<(usize, usize, Direction)>) {
//     for y in 0..maze.len() {
//         for x in 0..maze[0].len() {
//             if maze[y][x] == '#' {
//                 print!("#");
//             } else if visited.contains(&(x, y, Direction::N)) {
//                 print!("^");
//             } else if visited.contains(&(x, y, Direction::W)) {
//                 print!("<");
//             } else if visited.contains(&(x, y, Direction::E)) {
//                 print!(">");
//             } else if visited.contains(&(x, y, Direction::S)) {
//                 print!("V");
//             } else {
//                 print!(".");
//             }
//         }
//         println!("");
//     }
// }
//
// Too slow :(
// BFS + backtrack might take days to find the solution...
// fn pathfinder(
//     maze: &Maze,
//     x: usize,
//     y: usize,
//     direction: Direction,
//     curr_score: usize,
//     min_score: &mut usize,
//     visited: &mut HashSet<(usize, usize, Direction)>,
// ) {
//     // println!("Visiting: ({x}, {y}, {:?}) curr={curr_score}", direction);
//     // print_visited(maze, visited);
//     if maze[y][x] == 'E' {
//         if curr_score < *min_score {
//             *min_score = curr_score;
//         }
//         // println!("END ---> curr={curr_score} -- min={}", *min_score);
//         return;
//     }
//     if visited.contains(&(x, y, direction)) {
//         return;
//     }
//     visited.insert((x, y, direction));
//     // North
//     if maze[y - 1][x] != '#' && direction != Direction::S {
//         let mut new_score = curr_score;
//         if direction == Direction::N {
//             new_score += 1;
//         } else {
//             new_score += 1001;
//         }
//         pathfinder(maze, x, y - 1, Direction::N, new_score, min_score, visited);
//     }
//     // South
//     if maze[y + 1][x] != '#' && direction != Direction::N {
//         let mut new_score = curr_score;
//         if direction == Direction::S {
//             new_score += 1;
//         } else {
//             new_score += 1001;
//         }
//         pathfinder(maze, x, y + 1, Direction::S, new_score, min_score, visited);
//     }
//     // East
//     if maze[y][x + 1] != '#' && direction != Direction::W {
//         let mut new_score = curr_score;
//         if direction == Direction::E {
//             new_score += 1;
//         } else {
//             new_score += 1001;
//         }
//         pathfinder(maze, x + 1, y, Direction::E, new_score, min_score, visited);
//     }
//     // West
//     if maze[y][x - 1] != '#' && direction != Direction::E {
//         let mut new_score = curr_score;
//         if direction == Direction::W {
//             new_score += 1;
//         } else {
//             new_score += 1001;
//         }
//         pathfinder(maze, x - 1, y, Direction::W, new_score, min_score, visited);
//     }
//     visited.remove(&(x, y, direction));
// }
//
// fn part1(maze: &Maze) {
//     let mut min_score = usize::MAX;
//     let mut visited = HashSet::new();
//     pathfinder(
//         maze,
//         1,
//         maze.len() - 2,
//         Direction::E,
//         0,
//         &mut min_score,
//         &mut visited,
//     );
//     println!("The lowest score a reindeer can get is {min_score}");
// }

fn part1(maze: &Maze) {
    // Heap structure: (cost, x, y, displacement_row, displacement_col)
    // displacement_row/col explanation:
    // North: (0, 1)
    // South: (0, -1)
    // East: (1, 0)
    // West: (-1, 0)
    let mut prio_queue = BinaryHeap::new();
    prio_queue.push((0 as i32, 1, (maze.len() - 2) as i32, 0 as i32, 1 as i32));
    let mut min_score = i32::MAX;
    let mut visited = HashSet::new();

    // Use Djikstra algorithm to find best (minimum) path
    while let Some((score, x, y, disp_row, disp_col)) = prio_queue.pop() {
        visited.insert((x, y, disp_row, disp_col));
        if maze[y as usize][x as usize] == 'E' {
            min_score = -score;
            break;
        }

        // Use negative costs because the BinaryHeap is a max-heap
        for (new_score, new_x, new_y, new_dr, new_dc) in vec![
            (score - 1, x + disp_col, y + disp_row, disp_row, disp_col),
            (score - 1000, x, y, -disp_col, disp_row),
            (score - 1000, x, y, disp_col, -disp_row),
        ] {
            if maze[new_y as usize][new_x as usize] == '#' {
                continue;
            }
            if visited.contains(&(new_x, new_y, new_dr, new_dc)) {
                continue;
            }
            prio_queue.push((new_score, new_x, new_y, new_dr, new_dc));
        }
    }
    println!("The lowest score a reindeer can get is {min_score}");
}

// Thanks hyper neutrino, your explanations are very precious:
// https://www.youtube.com/watch?v=BJhpteqlVPM
fn part2(maze: &Maze) {
    // To get the next cheapest option
    let mut prio_queue = BinaryHeap::new();
    prio_queue.push((
        0 as i32,
        1 as i32,
        (maze.len() - 2) as i32,
        0 as i32,
        1 as i32,
    ));

    // For each state, the cheapest option to get there
    let mut lowest_scores = HashMap::new();
    lowest_scores.insert((1 as i32, (maze.len() - 2) as i32, 0 as i32, 1 as i32), 0);

    // List of previous states for a state
    let mut backtrack = HashMap::new();

    // Lowest score to find all best paths
    let mut min_score = i32::MAX;

    // List of all paths that led to the end
    let mut end_states = HashSet::new();

    // Get next best step
    while let Some((score, x, y, disp_row, disp_col)) = prio_queue.pop() {
        // If the current score is bigger than the best score, we're not
        // interested in this path
        if score
            > *lowest_scores
                .get(&(x, y, disp_row, disp_col))
                .unwrap_or(&i32::MAX)
        {
            continue;
        }

        // If the current tile is the end tile...
        if maze[y as usize][x as usize] == 'E' {
            // If the current best score is bigger than the minimum it means
            // that we've already explored all possible minimum paths and there
            // is no point in continuing
            if -score > min_score {
                break;
            }
            // Adjust the best cost and save the end state
            min_score = -score;
            end_states.insert((x, y, disp_row, disp_col));
        }

        // Try to move a step or rotate
        for (new_score, new_x, new_y, new_dr, new_dc) in vec![
            (score - 1, x + disp_col, y + disp_row, disp_row, disp_col),
            (score - 1000, x, y, -disp_col, disp_row),
            (score - 1000, x, y, disp_col, -disp_row),
        ] {
            if maze[new_y as usize][new_x as usize] == '#' {
                continue;
            }
            // If the current score is worse than the minimum for this state
            // we want to ignore this state because we already computed a better
            // one, ...
            let minimum = *lowest_scores
                .get(&(new_x, new_y, new_dr, new_dc))
                .unwrap_or(&i32::MAX);
            if -new_score > minimum {
                continue;
            }
            // ..., otherwise we found a better path, so let's start save it
            if -new_score < minimum {
                // First empty the current one because we found a better path
                backtrack.insert(
                    (new_x, new_y, new_dr, new_dc),
                    HashSet::<(i32, i32, i32, i32)>::new(),
                );
                lowest_scores.insert((new_x, new_y, new_dr, new_dc), -new_score);
            }
            backtrack
                .get_mut(&(new_x, new_y, new_dr, new_dc))
                .unwrap()
                .insert((x, y, disp_row, disp_col));
            prio_queue.push((new_score, new_x, new_y, new_dr, new_dc));
        }
    }

    let mut states = VecDeque::new();
    for es in end_states.to_owned().into_iter() {
        states.push_back(es);
    }
    let mut visited = end_states.clone();

    // Find all the best paths backwards with a flood-fill visit
    while let Some(state) = states.pop_front() {
        for &prev in backtrack
            .get(&state)
            .unwrap_or(&HashSet::<(i32, i32, i32, i32)>::new())
        {
            if visited.contains(&prev) {
                continue;
            }
            visited.insert(state);
            states.push_back(prev);
        }
    }
    let mut best_seats = HashSet::new();
    for (x, y, _, _) in visited {
        best_seats.insert((x, y));
    }

    println!(
        "There are {} seats that guarantee the best experience.",
        best_seats.len()
    );
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let maze = parse_input(&args[1]);
    part1(&maze);
    part2(&maze);
}
