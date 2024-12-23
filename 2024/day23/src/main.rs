use std::{
    collections::{HashMap, HashSet, VecDeque},
    env,
    fs::read_to_string,
};

type Graph = HashMap<String, HashSet<String>>;

fn parse_input(filename: &str) -> Graph {
    let mut lan = HashMap::new();
    for line in read_to_string(filename).unwrap().lines() {
        let pcs = line.split("-").map(String::from).collect::<Vec<String>>();
        let pc0 = lan.entry(pcs[0].clone()).or_insert(HashSet::new());
        pc0.insert(pcs[1].clone());
        let pc1 = lan.entry(pcs[1].clone()).or_insert(HashSet::new());
        pc1.insert(pcs[0].clone());
    }
    lan
}

fn find_cycles(
    g: &Graph,
    visited: &mut HashMap<String, bool>,
    n: usize,
    curr_node: String,
    start: String,
    path: String,
    paths: &mut HashSet<String>,
) {
    // println!("Visiting: {curr_node} --- e=({start}) n={n}");
    let mut vs = visited.to_owned();
    let v = vs.get_mut(&curr_node).unwrap();
    *v = true;

    if n == 0 {
        *v = false;
        if g.get(&curr_node).unwrap().contains(&start) {
            let mut final_path = path.clone();
            final_path.pop();
            let mut parts = final_path
                .split(",")
                .map(String::from)
                .collect::<Vec<String>>();
            parts.sort();
            // Check that every computer is connected to each other and also that there are no
            // duplicate PCs in the same lan. Probably not really efficient but paths shouldn't be
            // too long...
            for i in 0..parts.len() - 1 {
                for j in i + 1..parts.len() {
                    if parts[i] == parts[j] || !g.get(&parts[i]).unwrap().contains(&parts[j]) {
                        return;
                    }
                }
            }
            paths.insert(parts.join(","));
        }
        return;
    }

    for adj in g.get(&curr_node).unwrap() {
        let va = visited.get(&adj.to_string()).unwrap();
        if !*va && !path.contains(adj) {
            let mut new_path = path.clone();
            new_path.push_str(&adj);
            new_path.push_str(",");

            find_cycles(
                g,
                visited,
                n - 1,
                adj.to_string(),
                start.to_string(),
                new_path,
                paths,
            );
        }
    }

    *v = false;
}

fn part1(lan: &Graph) {
    let mut paths = HashSet::new();
    let mut visited = HashMap::new();
    for node in lan.keys() {
        visited.insert(node.to_string(), false);
    }

    for node in lan.keys().skip(2) {
        find_cycles(
            lan,
            &mut visited,
            3,
            node.to_string(),
            node.to_string(),
            String::new(),
            &mut paths,
        );
        visited.insert(node.to_string(), true);
    }

    let mut lans_with_t_pc = 0;
    'p: for path in paths {
        let pcs = path.split(",");
        for pc in pcs {
            if pc.starts_with("t") {
                lans_with_t_pc += 1;
                continue 'p;
            }
        }
    }
    println!("There are {lans_with_t_pc} possible LANs where an historian might be");
}

// It probably works but it takes the universe age to compute the result :(
// fn reach_everything(g: &Graph, path: &String) -> bool {
//     let nodes = path.split(",").map(String::from).collect::<Vec<String>>();
//     let mut visited = HashMap::new();
//     for v in g.keys() {
//         visited.insert(v.to_string(), false);
//     }
//
//     let mut q = VecDeque::new();
//     for v in nodes {
//         q.push_back(v.to_string());
//     }
//     while let Some(curr) = q.pop_front() {
//         if *visited.get(&curr).unwrap() {
//             continue;
//         }
//         visited.insert(curr.to_string(), true);
//
//         for adj in g.get(&curr).unwrap() {
//             q.push_back(adj.to_string());
//         }
//     }
//
//     let ret = visited.values().filter(|&v| !v).count() == 0;
//     // println!("path={path} {visited:?} ret={ret}");
//     ret
// }
//
// fn part2(lan: &Graph) {
//     let mut size = 3;
//     let n_nodes = lan.keys().count();
//     let mut best_path = String::new();
//     let mut best_len = 0;
//     while size < n_nodes - 1 {
//         let mut paths = HashSet::new();
//         let mut visited = HashMap::new();
//         for node in lan.keys() {
//             visited.insert(node.to_string(), false);
//         }
//
//         for node in lan.keys() {
//             find_cycles(
//                 lan,
//                 &mut visited,
//                 size,
//                 node.to_string(),
//                 node.to_string(),
//                 String::new(),
//                 &mut paths,
//             );
//             visited.insert(node.to_string(), true);
//         }
//
//         println!("here: s={size} paths={}", paths.len());
//         let mut all = false;
//         for path in paths {
//             all = all || reach_everything(&lan, &path);
//             if all {
//                 if path.len() > best_len {
//                     best_len = path.len();
//                     best_path = path;
//                     break;
//                 }
//             }
//         }
//         if !all {
//             break;
//         }
//         size += 1;
//     }
//     println!("The password for the lan is {best_path}");
// }

fn find_connected_lans(
    g: &Graph,
    node: String,
    connected: HashSet<String>,
    possible_lans: &mut HashSet<String>,
) {
    let mut pk = connected.to_owned().into_iter().collect::<Vec<String>>();
    pk.sort();
    let path_key = pk.join(",");
    if possible_lans.contains(&path_key) {
        return;
    }
    possible_lans.insert(path_key);

    for adj in g.get(&node).unwrap() {
        if connected.contains(&adj.to_string()) {
            continue;
        }
        let mut connected_to_all = true;
        for other in connected.to_owned() {
            if !g.get(&other).unwrap().contains(adj) {
                connected_to_all = false;
                break;
            }
        }
        if !connected_to_all {
            continue;
        }
        let mut new_connected = connected.clone();
        new_connected.insert(adj.to_string());
        find_connected_lans(g, adj.to_string(), new_connected, possible_lans);
    }
}

fn part2(lan: &Graph) {
    let mut possible_lans = HashSet::new();
    let mut nodes = vec![];
    for k in lan.keys() {
        nodes.push(k);
    }

    println!("#nodes={}", nodes.len());
    for (i, v) in nodes.into_iter().enumerate() {
        println!("({i}) Checking node {v}...");
        let mut connected = HashSet::new();
        connected.insert(v.clone());
        find_connected_lans(lan, v.clone(), connected, &mut possible_lans);
    }

    let best_lan = possible_lans.into_iter().max_by_key(|p| p.len()).unwrap();
    let mut parts = best_lan.split(",").collect::<Vec<&str>>();
    parts.sort();
    let password = parts.join(",");

    println!("The password for the LAN party is: {password}");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let lan = parse_input(&args[1]);
    part1(&lan);
    part2(&lan);
}
