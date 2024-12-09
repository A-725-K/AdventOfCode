use std::{collections::HashMap, env, fs::read_to_string};

fn parse_input(filename: &str) -> (Vec<Vec<i32>>, Vec<i32>, String) {
    let s = read_to_string(filename).unwrap();
    let memorymap = s.trim_end();
    let mut slacks: Vec<i32> = vec![];
    let mut files: Vec<Vec<i32>> = vec![];
    let mut file_id = 0;
    for (idx, c) in memorymap.chars().enumerate() {
        let mut n = c.to_digit(10).unwrap() as i32;
        if idx % 2 == 0 {
            let mut curr_files = vec![];
            while n > 0 {
                curr_files.push(file_id);
                n -= 1;
            }
            files.push(curr_files);
            file_id += 1;
        } else {
            slacks.push(n);
        }
    }
    (files, slacks, s.clone())
}

/// WAY TOO CONVOLUTED!!!!! I ALSO DIDN'T UNDERSTAND HOW FILES WITH ID > 10
/// WERE HANDLED IN MEMORY :(
// fn parse_input(filename: &str) -> (Vec<String>, Vec<u32>) {
//     let s = read_to_string(filename).unwrap();
//     let memorymap = s.trim_end();
//     let mut slacks: Vec<u32> = vec![];
//     let mut files: Vec<String> = vec![];
//
//     let mut file_id = 0;
//     for (idx, c) in memorymap.chars().enumerate() {
//         if idx % 2 == 0 {
//             let file_id_str: Vec<char> = format!("{file_id}").chars().collect();
//             let mut complete_file_id = String::from("");
//             for i in 0..c.to_digit(10).unwrap() {
//                 let file_idx = (i as usize) % file_id_str.len();
//                 complete_file_id.push(file_id_str[file_idx]);
//             }
//             files.push(complete_file_id);
//             file_id += 1;
//         } else {
//             slacks.push(c.to_digit(10).unwrap());
//         }
//     }
//     (files, slacks)
// }
//
// fn part1(files: &Vec<String>, slacks: &Vec<u32>) {
//     let mut memory_defrag = String::from(&files[0]);
//     let mut memory_defrag_crc: usize = 0;
//     let mut memory_idx: usize = 1;
//     let mut file_to_move_idx = files.len() - 1;
//     let mut file_to_read_idx = 1;
//     let mut file_idx = (files[file_to_move_idx].len() - 1) as i32;
//     for &s in slacks {
//         let mut slack = s as i32;
//         let mut file_to_move: Vec<char> = files[file_to_move_idx].chars().collect();
//         'HERE: while slack > 0 {
//             while file_idx >= 0 {
//                 // println!(
//                 //     "{:?} -- {} -- {}",
//                 //     file_to_move, file_to_move[file_idx as usize], file_idx
//                 // );
//                 memory_defrag.push(file_to_move[file_idx as usize]);
//                 // let c = file_to_move[file_idx as usize].to_digit(10).unwrap() as u64;
//                 // println!("{memory_idx}*{c}");
//                 // memory_defrag_crc += memory_idx * c;
//                 memory_idx += 1;
//                 file_idx -= 1;
//                 slack -= 1;
//                 if slack == 0 {
//                     break 'HERE;
//                 }
//             }
//             file_to_move_idx -= 1;
//             file_idx = (files[file_to_move_idx].len() - 1) as i32;
//             file_to_move = files[file_to_move_idx].chars().collect();
//         }
//         // println!(
//         //     "ADDING FILE ---> {} == {}",
//         //     files[file_to_read_idx], file_idx
//         // );
//         if file_to_read_idx >= file_to_move_idx {
//             // println!("FILE_IDX={file_idx} -- FILE_TO_MOVE_IDX={file_to_move_idx}");
//             break;
//         }
//         for c in files[file_to_read_idx].chars() {
//             memory_defrag.push(c);
//         }
//         file_to_read_idx += 1;
//     }
//     let file_to_finish_to_move: Vec<char> = files[file_to_move_idx].chars().collect();
//     while file_idx >= 0 {
//         memory_defrag.push(file_to_finish_to_move[file_idx as usize]);
//         // let c = file_to_finish_to_move[file_idx as usize]
//         //     .to_digit(10)
//         //     .unwrap() as u64;
//         //
//         // memory_defrag_crc += memory_idx * c;
//         // memory_idx += 1;
//         // println!("{memory_idx}*{c}");
//         file_idx -= 1;
//     }
//     println!("{memory_defrag}");
//     for (i, c) in memory_defrag.chars().enumerate() {
//         memory_defrag_crc += (i * c.to_digit(10).unwrap() as usize);
//     }
//     println!("The checksum after defragmenting the file system is {memory_defrag_crc}");
// }

fn part1(fs: &Vec<Vec<i32>>, slacks: &Vec<i32>) {
    let mut files = fs.clone();
    let mut memory_defrag_crc: usize = 0;
    let mut memory_defrag = vec![];
    let mut file_to_move_idx = (files.len() - 1) as i32;
    let mut file_to_read_idx = 1;
    let mut file_to_move = files[file_to_move_idx as usize].clone();
    let mut file_idx = (file_to_move.len() - 1) as i32;

    while !files[0].is_empty() {
        memory_defrag.push(files[0].pop().unwrap());
    }

    for &s in slacks {
        let mut slack = s;
        while slack > 0 {
            if file_idx < 0 {
                file_to_move_idx -= 1;
                file_to_move = files[file_to_move_idx as usize].clone();
                if file_to_move.is_empty() {
                    break;
                }
                file_idx = (file_to_move.len() - 1) as i32;
            }
            memory_defrag.push(file_to_move[file_idx as usize]);
            file_idx -= 1;
            slack -= 1;
            if slack == 0 {
                break;
            }
        }
        if file_to_read_idx >= file_to_move_idx {
            break;
        }
        while !files[file_to_read_idx as usize].is_empty() {
            memory_defrag.push(files[file_to_read_idx as usize].pop().unwrap());
        }
        file_to_read_idx += 1;
    }
    while file_idx >= 0 {
        memory_defrag.push(files[file_to_read_idx as usize][file_idx as usize]);
        file_idx -= 1;
    }
    // println!("{:?}", memory_defrag);

    for (idx, mem) in memory_defrag.into_iter().enumerate() {
        memory_defrag_crc += idx * mem as usize;
    }
    println!("The checksum after naively defragmenting the file system is {memory_defrag_crc}");
}

fn part2(s: String) {
    let mut memory_defrag: HashMap<i32, (i32, i32)> = HashMap::new();
    let mut slacks = vec![];

    let mut file_idx = 0;
    let mut position = 0;
    for (idx, ch) in s.trim_end().chars().enumerate() {
        let n = ch.to_digit(10).unwrap() as i32;
        if idx % 2 == 0 {
            memory_defrag.insert(file_idx, (position, n));
            file_idx += 1;
        } else {
            if n > 0 {
                slacks.push((position, n));
            }
        }
        position += n;
    }

    while file_idx > 0 {
        file_idx -= 1;
        let (current_position, space_required) = memory_defrag.get_mut(&file_idx).unwrap();
        for (idx, (slack_start, slack_len)) in slacks.to_owned().into_iter().enumerate() {
            // Remove all blanks on the right of the current one
            if slack_start >= *current_position {
                slacks.resize(idx, (-2, -2));
                break;
            }
            // There is space: let's move the file
            if *space_required <= slack_len {
                *current_position = slack_start;
                *space_required = *space_required;
                if *space_required == slack_len {
                    // No more blanks in this slot
                    slacks.remove(idx);
                } else {
                    // Update the slack and update the space left in that slack
                    slacks[idx] = (slack_start + *space_required, slack_len - *space_required);
                }
                break;
            }
        }
    }
    let mut memory_defrag_crc: i64 = 0;
    for (file_idx, (current_position, space_on_disk)) in memory_defrag.into_iter() {
        for idx in current_position..current_position + space_on_disk {
            memory_defrag_crc += (file_idx * idx) as i64;
        }
    }
    println!("The checksum after properly defragmenting the file system is {memory_defrag_crc}");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let (files, slacks, inp) = parse_input(&args[1]);
    part1(&files, &slacks);
    part2(inp);
}
