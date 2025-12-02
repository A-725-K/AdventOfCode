use std::collections::HashSet;

fn parse_input(line: &String) -> Vec<Range> {
    line.split(",")
        .map(|range_str| {
            let range_fields: Vec<_> = range_str.split("-").collect();
            Range {
                start: range_fields[0].parse::<usize>().unwrap(),
                end: range_fields[1].parse::<usize>().unwrap(),
            }
        })
        .collect()
}

fn get_number_of_digits(n: &usize) -> u32 {
    let (mut digits, mut i) = (0, *n);
    while i > 0 {
        i /= 10;
        digits += 1;
    }
    digits
}

fn repeat_num(n: usize, reps: usize) -> usize {
    format!("{n}").repeat(reps).parse::<usize>().unwrap()
}

#[derive(Debug)]
struct Range {
    start: usize,
    end: usize,
}

impl Range {
    fn invalid_in_range(&self) -> usize {
        let mut sum_of_invalid = 0;

        let mut begin_p = get_number_of_digits(&self.start);
        begin_p /= 2;
        if begin_p != 0 {
            begin_p -= 1;
        }
        let end_p = get_number_of_digits(&self.end);

        let begin = 10u64.pow(begin_p) as usize;
        let end = 10u64.pow(end_p - 1) as usize;
        for i in begin..end + 1 {
            let new_n = repeat_num(i, 2);
            if new_n < self.start {
                continue;
            }
            if new_n > self.end + 1 {
                break;
            }
            if new_n >= self.start && new_n <= self.end {
                sum_of_invalid += new_n;
            }
        }
        sum_of_invalid
    }

    fn invalid_in_range_repeated(&self) -> usize {
        let mut sum_of_invalid = 0;
        let digits = get_number_of_digits(&self.end);
        let end = 10u64.pow((digits as f64 / 2.0).ceil() as u32);
        let mut seen = HashSet::new();
        for i in 1..end + 1 {
            let mut repetitions = 0;
            let mut new_n = 0;
            while new_n < self.end + 1 {
                repetitions += 1;
                new_n = repeat_num(i as usize, repetitions);
                if new_n < self.start {
                    continue;
                }
                if new_n > self.end + 1 {
                    break;
                }

                if new_n >= self.start && new_n <= self.end && !seen.contains(&new_n) {
                    if new_n <= 100 && !new_n.is_multiple_of(11) {
                        continue;
                    }
                    sum_of_invalid += new_n;
                    seen.insert(new_n);
                }
            }
        }
        sum_of_invalid
    }
}

fn find_invalid_ids(line: &String, repeated: bool) {
    let ranges = parse_input(&line);
    let mut sum_of_invalid_ids = 0;
    for rng in ranges {
        if repeated {
            sum_of_invalid_ids += rng.invalid_in_range_repeated();
        } else {
            sum_of_invalid_ids += rng.invalid_in_range();
        }
    }
    println!("The sum of all invalid IDs is {sum_of_invalid_ids}");
}

pub fn part1(lines: &Vec<String>, _day: usize) {
    find_invalid_ids(&lines[0], false);
}

pub fn part2(lines: &Vec<String>, _day: usize) {
    find_invalid_ids(&lines[0], true);
}
