use std::{collections::HashMap, env, fs::read_to_string};

fn parse_input(filename: &str) -> (HashMap<usize, Vec<usize>>, Vec<Vec<usize>>) {
    let mut rules: HashMap<usize, Vec<usize>> = HashMap::new();
    let mut queues: Vec<Vec<usize>> = vec![];
    let mut parse_rules = true;
    for line in read_to_string(filename).unwrap().lines() {
        if line == "" {
            parse_rules = false;
            continue;
        }
        if parse_rules {
            let fields: Vec<usize> = line
                .split('|')
                .map(|n| n.parse::<usize>().unwrap())
                .collect();
            rules
                .entry(fields[0])
                .and_modify(|l| l.push(fields[1]))
                .or_insert(vec![fields[1]]);
        } else {
            let queue: Vec<usize> = line
                .split(',')
                .map(|n| n.parse::<usize>().unwrap())
                .collect();
            queues.push(queue);
        }
    }
    (rules, queues)
}

fn is_correct(queue: &Vec<usize>, rules: &HashMap<usize, Vec<usize>>) -> bool {
    let n = queue.len();
    for i in 1..n {
        if let Some(el_rules) = rules.get(&queue[i]) {
            for r in el_rules {
                if queue[0..i].contains(r) {
                    return false;
                }
            }
        }
    }
    true
}

fn part1(rules: &HashMap<usize, Vec<usize>>, queues: &Vec<Vec<usize>>) -> Vec<usize> {
    let mut correct_updates: Vec<usize> = vec![];
    let mut wrong_updates: Vec<usize> = vec![];
    for (idx, queue) in queues.iter().enumerate() {
        if is_correct(&queue, rules) {
            correct_updates.push(idx);
        } else {
            wrong_updates.push(idx)
        }
    }

    let mut sum_of_middle_pages = 0;
    for idx in correct_updates {
        let mid = queues[idx].len() / 2;
        sum_of_middle_pages += queues[idx][mid];
    }

    println!("The sum of the middle pages of the correct queues is {sum_of_middle_pages}");
    wrong_updates
}

fn fix_queue(queue: &Vec<usize>, rules: &HashMap<usize, Vec<usize>>) -> Vec<usize> {
    let n = queue.len();
    let mut fixed_queue = vec![0; n];

    for &el in queue {
        let el_rules = match rules.get(&el) {
            Some(v) => v.clone(),
            None => vec![],
        };
        let mut occurences = 0;
        for r in el_rules.iter() {
            if queue.contains(&r) {
                occurences += 1;
            }
        }
        fixed_queue[n - occurences - 1] = el;
    }

    fixed_queue
}

fn part2(rules: &HashMap<usize, Vec<usize>>, queues: &Vec<Vec<usize>>, wrong_updates: &Vec<usize>) {
    let mut sum_of_middle_pages_of_fixed = 0;
    for &idx in wrong_updates.iter() {
        let new_queue = fix_queue(&queues[idx], rules);
        let mid = new_queue.len() / 2;
        sum_of_middle_pages_of_fixed += new_queue[mid];
    }
    println!("The sum of the middle pages of the fixed queues is {sum_of_middle_pages_of_fixed}");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let (rules, queues) = parse_input(&args[1]);
    // println!("rules:\n{:?}\nqueues:\n{:?}", rules, queues);
    let wrong_updates = part1(&rules, &queues);
    part2(&rules, &queues, &wrong_updates);
}
