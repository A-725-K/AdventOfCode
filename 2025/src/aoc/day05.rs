use std::cmp;

#[derive(Debug, Clone)]
struct Range {
    start: usize,
    end: usize,
}

fn parse_input(lines: &Vec<String>) -> (Vec<Range>, Vec<usize>) {
    let mut inventory = vec![];
    let mut aliments = vec![];

    let mut i = 0;
    loop {
        if lines[i].is_empty() {
            i += 1;
            break;
        }
        let fields: Vec<_> = lines[i].split("-").collect();
        inventory.push(Range {
            start: fields[0].parse::<usize>().unwrap(),
            end: fields[1].parse::<usize>().unwrap(),
        });
        i += 1;
    }

    while i < lines.len() {
        aliments.push(lines[i].parse::<usize>().unwrap());
        i += 1;
    }
    (inventory, aliments)
}

fn is_fresh(aliment: usize, inventory: &Vec<Range>) -> bool {
    for range in inventory {
        if aliment >= range.start && aliment <= range.end {
            return true;
        }
    }
    false
}

fn merge_ranges(inventory: &mut Vec<Range>) -> Vec<Range> {
    inventory.sort_by_key(|rng| rng.start);

    let mut inv_stack: Vec<Range> = vec![inventory[0].clone()];
    for range in inventory {
        // If current interval overlaps with the last merged interval, merge them
        if range.start <= inv_stack.last().unwrap().end {
            let last = inv_stack.last_mut().unwrap();
            last.end = cmp::max(range.end, last.end);
        } else {
            inv_stack.push(range.clone());
        }
    }

    inv_stack
}

pub fn part1(lines: &Vec<String>, _day: usize) {
    let (inventory, aliments) = parse_input(&lines);
    let mut fresh_aliments = 0;
    for aliment in aliments {
        if is_fresh(aliment, &inventory) {
            fresh_aliments += 1;
        }
    }
    println!("There are {fresh_aliments} fresh aliments available");
}

pub fn part2(lines: &Vec<String>, _day: usize) {
    let (mut inventory, _) = parse_input(&lines);
    let mut fresh_aliments = 0;
    let merged_inventory = merge_ranges(&mut inventory);
    for range in merged_inventory {
        fresh_aliments += range.end - range.start + 1;
    }
    println!("There are in total {fresh_aliments} fresh aliments");
}
