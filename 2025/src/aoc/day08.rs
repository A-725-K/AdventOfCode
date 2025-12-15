use std::{cmp::Ordering, collections::HashMap, fmt::Debug};
use union_find::{QuickUnionUf, UnionBySize, UnionFind};

fn cmp_f64(a: &(usize, usize, f64), b: &(usize, usize, f64)) -> Ordering {
    if a.2 < b.2 {
        return Ordering::Less;
    }
    if a.2 > b.2 {
        return Ordering::Greater;
    }
    return Ordering::Equal;
}

#[derive(Clone, Copy)]
struct Junction {
    x: usize,
    y: usize,
    z: usize,
    id: usize,
}

impl Debug for Junction {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "({},{},{})", self.x, self.y, self.z)
    }
}

impl Junction {
    fn distance(&self, other: &Junction) -> f64 {
        ((self.x as f64 - other.x as f64).powf(2.0)
            + (self.y as f64 - other.y as f64).powf(2.0)
            + (self.z as f64 - other.z as f64).powf(2.0))
        .sqrt()
    }
}

fn parse_input(lines: &Vec<String>) -> Vec<Junction> {
    let mut id = 0;
    lines
        .iter()
        .map(|line| {
            let fields: Vec<_> = line.split(",").collect();
            id += 1;
            Junction {
                x: fields[0].parse::<usize>().unwrap(),
                y: fields[1].parse::<usize>().unwrap(),
                z: fields[2].parse::<usize>().unwrap(),
                id: id - 1,
            }
        })
        .collect()
}

fn find_junction_by_id(junctions: &Vec<Junction>, id: usize) -> Junction {
    *(junctions.iter().find(|j| j.id == id).unwrap())
}

fn connect_junctions(junctions: &Vec<Junction>) -> Vec<(usize, usize, f64)> {
    let n = junctions.len();
    let mut all_pairs = vec![];
    for i in 0..n {
        for j in i + 1..n {
            all_pairs.push((
                junctions[i].id,
                junctions[j].id,
                junctions[i].distance(&junctions[j]),
            ))
        }
    }
    all_pairs.sort_by(cmp_f64);
    all_pairs
}

fn collect_sizes(circuits: &mut QuickUnionUf<UnionBySize>, n: usize) -> Vec<usize> {
    let mut circuit_sizes = HashMap::new();
    for i in 0..n {
        let root = circuits.find(i);
        circuit_sizes
            .entry(root)
            .and_modify(|csz| *csz += 1)
            .or_insert(1.to_owned());
    }
    let mut sizes: Vec<_> = circuit_sizes.values().cloned().collect();
    sizes.sort();
    sizes
}

pub fn part1(lines: &Vec<String>, _day: usize) {
    let junctions = parse_input(lines);
    let iterations = 1000;
    let all_pairs = connect_junctions(&junctions);

    let mut circuits = QuickUnionUf::<UnionBySize>::new(junctions.len());
    for it in 0..iterations {
        let next = all_pairs[it];
        let c1 = circuits.find(next.0);
        let c2 = circuits.find(next.1);
        circuits.union(c1, c2);
    }

    let circuits_sizes = collect_sizes(&mut circuits, junctions.len());
    let n = circuits_sizes.len();
    println!(
        "The product of the sizes of the 3 biggest circuits is {}",
        circuits_sizes[n - 1] * circuits_sizes[n - 2] * circuits_sizes[n - 3],
    );
}

pub fn part2(lines: &Vec<String>, _day: usize) {
    let junctions = parse_input(lines);
    let all_pairs = connect_junctions(&junctions);

    let mut circuits = QuickUnionUf::<UnionBySize>::new(junctions.len());
    let mut final_product = 0;
    for next in all_pairs {
        let c1 = circuits.find(next.0);
        let c2 = circuits.find(next.1);
        circuits.union(c1, c2);
        let circuit_sizes = collect_sizes(&mut circuits, junctions.len());
        if circuit_sizes.len() == 1 {
            let p1 = find_junction_by_id(&junctions, next.0);
            let p2 = find_junction_by_id(&junctions, next.1);
            final_product = p1.x * p2.x;
            break;
        }
    }
    println!(
        "If you multiply together the X coordinates of the last two junctions you get {final_product}"
    );
}
