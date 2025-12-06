pub fn part1(lines: &Vec<String>, _day: usize) {
    let mut worksheet = vec![];
    let mut operations = vec![];

    for (i, line) in lines.iter().enumerate() {
        if i == lines.len() - 1 {
            operations = line.split_whitespace().collect();
        } else {
            let row = line
                .split_whitespace()
                .map(|n| n.parse::<usize>().unwrap())
                .collect::<Vec<usize>>();
            worksheet.push(row);
        }
    }
    let n = worksheet[0].len();
    let mut results = vec![];
    for i in 0..n {
        if operations[i] == "+" {
            results.push(0);
        } else if operations[i] == "*" {
            results.push(1);
        } else {
            panic!("Unknown operation");
        }
    }
    for i in 0..n {
        for j in 0..worksheet.len() {
            if operations[i] == "+" {
                results[i] += worksheet[j][i];
            } else if operations[i] == "*" {
                results[i] *= worksheet[j][i];
            } else {
                panic!("Unknown operation");
            }
        }
    }
    println!(
        "The grand total of all the operations is {}",
        results.iter().fold(0, |acc, el| { acc + el })
    );
}

fn digits2n(digits: &Vec<usize>) -> usize {
    digits.iter().fold(0, |acc, n| acc * 10 + n)
}

pub fn part2(lines: &Vec<String>, _day: usize) {
    let mut grand_result = vec![];
    let all_chars = lines
        .iter()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut start_idx = (all_chars[0].len() - 1) as i32; // assuming all lines are the same length
    let mut end_idx = start_idx;
    let operations_idx = lines.len() - 1;
    while start_idx >= 0 {
        // find the max length of the numbers for this operation
        while all_chars[operations_idx][start_idx as usize] == ' ' {
            start_idx -= 1;
        }

        let mut result = if all_chars[operations_idx][start_idx as usize] == '+' {
            0
        } else if all_chars[operations_idx][start_idx as usize] == '*' {
            1
        } else {
            panic!("Unknown operation")
        };

        // parse current operation
        for i in start_idx..end_idx + 1 {
            let mut digits = vec![];
            // consider only digits that are present, ignore if whitespace in the column
            for j in 0..operations_idx {
                if all_chars[j][i as usize] != ' ' {
                    digits.push(all_chars[j][i as usize].to_digit(10).unwrap() as usize);
                }
            }
            let n = digits2n(&digits);
            if all_chars[operations_idx][start_idx as usize] == '+' {
                result += n;
            } else if all_chars[operations_idx][start_idx as usize] == '*' {
                result *= n;
            } else {
                panic!("Unknown operation");
            }
        }
        grand_result.push(result);

        // consume blank column
        start_idx -= 1;
        end_idx = start_idx - 1;
    }
    println!(
        "The grand result after parsing correctly the worksheet is {}",
        grand_result.iter().fold(0, |acc, el| { acc + el })
    );
}
