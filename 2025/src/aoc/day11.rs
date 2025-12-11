use std::collections::{HashMap, HashSet};

fn parse_input(lines: &Vec<String>) -> HashMap<String, Vec<&str>> {
    lines.iter().fold(HashMap::new(), |mut acc, line| {
        let fields: Vec<_> = line.split(" ").collect();
        acc.insert(fields[0].replace(":", ""), fields[1..].to_vec());
        acc
    })
}

fn visit(graph: &HashMap<String, Vec<&str>>, start: &str, end: &str) -> usize {
    let mut seen = HashSet::new();
    visit_rec(&graph, String::from(start), end, &mut seen)
}
fn visit_rec(
    graph: &HashMap<String, Vec<&str>>,
    curr: String,
    end: &str,
    seen: &mut HashSet<String>,
) -> usize {
    if curr == end {
        return 1;
    }
    if seen.contains(&curr.clone()) {
        return 0;
    }
    seen.insert(curr.clone());

    let mut paths = 0;
    for adj in graph.get(&String::from(curr.clone())).unwrap() {
        paths += visit_rec(&graph, adj.to_string(), end, seen);
    }
    // Backtrack only if you found a path to the end node, otherwise does not
    // make sense to pursue any more path in that sub-tree
    if paths != 0 {
        seen.remove(&curr);
    }
    paths
}

pub fn part1(lines: &Vec<String>, _day: usize) {
    let mut graph = parse_input(&lines);
    graph.insert(String::from("out"), vec![]);
    let paths = visit(&graph, "you", "out");

    println!("There are {paths} different paths from 'you' to 'out' nodes");
}

pub fn part2(lines: &Vec<String>, _day: usize) {
    let mut graph = parse_input(&lines);
    graph.insert(String::from("out"), vec![]);

    let fft2dac = visit(&graph, "fft", "dac");
    // println!("0 -- fft -> dac: {fft2dac}");

    let paths = if fft2dac == 0 {
        // Case 1: start --> ... -> dac --> ... -> fft -> ... -> out
        visit(&graph, "svr", "dac") * visit(&graph, "dac", "fft") * visit(&graph, "fft", "out")
    } else {
        // Case 2: start --> ... -> fft --> ... -> dac -> ... -> out
        let svr2fft = visit(&graph, "svr", "fft");
        // println!("2.1 -- svr -> fft: {svr2fft}");
        let dac2out = visit(&graph, "dac", "out");
        // println!("2.2 -- dac -> out: {dac2out}");
        svr2fft * fft2dac * dac2out
    };

    println!(
        "There are {paths} different paths from 'svr' to 'out' nodes that use 'dac' and 'fft' modules"
    );
}
