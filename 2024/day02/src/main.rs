use std::{env, fs::read_to_string};

fn parse_input(filename: &str) -> Vec<Vec<i32>> {
    let mut reports: Vec<Vec<i32>> = vec![];
    for line in read_to_string(filename).unwrap().lines() {
        let report: Vec<i32> = line.split(' ').map(|n| n.parse::<i32>().unwrap()).collect();
        reports.push(report);
    }
    reports
}

fn is_safe(report: &Vec<i32>) -> bool {
    let (mut is_desc, mut is_asc) = (true, true);
    for i in 1..report.len() {
        is_desc = is_desc && report[i - 1] > report[i];
        is_asc = is_asc && report[i - 1] < report[i];
        let diff = (report[i - 1] - report[i]).abs();
        if diff < 1 || diff > 3 {
            return false;
        }
    }
    is_asc || is_desc
}

// FALSY ASSUMPTIONS IN HERE:
//  - this solution was not acutally skipping the faulty level "i" and comparing
//    "i-1" and "i+1", but considering the element that should have been removed
//    instead
// fn is_safe_dampener(report: &Vec<i32>) -> bool {
//     let (mut is_desc, mut is_asc, mut diff_ok) = (true, true, true);
//     let (mut is_desc_tol, mut is_asc_tol, mut diff_ok_tol) = (true, true, true);
//     let (mut dti, mut ati): (usize, usize) = (0, 0);
//     println!("{:?}", report);
//     for i in 1..report.len() {
//         is_desc = is_desc && report[i - 1] > report[i];
//         if !is_desc && is_desc_tol {
//             is_desc = true;
//             is_desc_tol = false;
//             if dti == 0 {
//                 dti = i;
//             }
//             // println!("element {i} bonus is_desc");
//         }
//         is_asc = is_asc && report[i - 1] < report[i];
//         if !is_asc && is_asc_tol {
//             is_asc = true;
//             is_asc_tol = false;
//             if ati == 0 {
//                 ati = i;
//             }
//             // println!("element {i} bonus is_asc");
//         }
//         let diff = (report[i - 1] - report[i]).abs();
//         println!("{i} ====== diff={diff}");
//         diff_ok = diff_ok && diff >= 1 && diff <= 3;
//         if !diff_ok && diff_ok_tol {
//             if ati == i || dti == i {
//                 diff_ok_tol = true;
//                 diff_ok = true;
//                 continue;
//             }
//             if diff_ok_tol {
//                 diff_ok_tol = false;
//                 diff_ok = true;
//                 println!("element {i} bonus diff_ok");
//                 continue;
//             }
//             // diff_ok = false;
//             return false;
//         }
//     }
//     let ret = diff_ok && (is_asc || is_desc);
//     println!("result --> asc={is_asc} desc={is_desc} diff_ok={diff_ok} res={ret}\n");
//     ret
// }

fn is_safe_dampener(report: &Vec<i32>) -> bool {
    // O(N^2) It can be better, I know... :(
    for i in 0..report.len() {
        let mut new_report = report.clone();
        new_report.remove(i);
        if is_safe(&new_report) {
            return true;
        }
    }
    false
}

fn part1(reports: Vec<Vec<i32>>) {
    let mut safe_reports = 0;
    for report in reports {
        if is_safe(&report) {
            safe_reports += 1;
        }
    }
    println!("There are {} safe reports", safe_reports);
}

fn part2(reports: Vec<Vec<i32>>) {
    let mut safe_reports_damp = 0;
    for report in reports {
        if is_safe_dampener(&report) {
            safe_reports_damp += 1;
        }
    }
    println!("There are {} safe dampened reports", safe_reports_damp);
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let reports = parse_input(&args[1]);
    part1(reports.clone());
    part2(reports);
}
