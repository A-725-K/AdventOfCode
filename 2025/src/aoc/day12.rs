type Shape = (Vec<Vec<char>>, usize);
type Region = (Vec<usize>, usize, usize);

// fn print_present(present: &Shape) {
//     present
//         .0
//         .iter()
//         .for_each(|r| println!("{}", r.iter().collect::<String>()));
//     println!("");
// }

// WTF is not needed??? xDxDxDxDxD
// fn rotate_present(present: &Shape) -> Vec<Shape> {
//     let mut rotations: Vec<Shape> = vec![];
//     let n = present.len();
//
//     for rotation in 0..3 {
//         let mut m = vec![vec!['.'; n]; n];
//
//         for i in 0..n {
//             for j in 0..n {
//                 if rotation == 0 {
//                     m[(n - j - 1) % n][i] = present[i][j];
//                 } else {
//                     m[(n - j - 1) % n][i] = rotations[rotation - 1][i][j];
//                 }
//             }
//         }
//         rotations.push(m);
//     }
//
//     rotations
// }

fn parse_input(lines: &Vec<String>) -> (Vec<Shape>, Vec<Region>) {
    let mut presents = vec![];
    let mut regions = vec![];

    let mut parse_presents = false;
    let mut curr_present = vec![];
    for line in lines {
        // parse presents
        if line.is_empty() {
            parse_presents = false;
            let blocks = curr_present.iter().fold(0, |acc, r: &Vec<_>| {
                acc + r.iter().filter(|&c| *c == '#').count()
            });
            presents.push((curr_present.clone(), blocks));
            // presents.extend(rotate_present(&curr_present));
            curr_present.clear();
            continue;
        }
        if line.chars().last().unwrap() == ':' {
            parse_presents = true;
            continue;
        }

        if parse_presents {
            curr_present.push(line.chars().collect::<Vec<char>>());
        } else {
            let fields: Vec<_> = line.split(":").map(|f| f.trim()).collect();
            let sizes: Vec<_> = fields[0].split("x").collect();
            let (rows, cols) = (
                sizes[1].parse::<usize>().unwrap(),
                sizes[0].parse::<usize>().unwrap(),
            );
            let num_presents = fields[1]
                .split(" ")
                .map(|f| f.parse::<usize>().unwrap())
                .collect();
            regions.push((num_presents, rows, cols));
        }
    }

    (presents, regions)
}

fn can_fit_presents_in_region(presents: &Vec<Shape>, region_data: &Region) -> bool {
    let mut block_count = 0;
    for (idx, num_presents) in region_data.0.iter().enumerate() {
        block_count += presents[idx].1 * num_presents;
    }
    block_count <= region_data.1 * region_data.2
}

pub fn part1(lines: &Vec<String>, _day: usize) {
    let (presents, regions) = parse_input(&lines);
    let mut valid_regions = 0;

    for region in regions {
        if can_fit_presents_in_region(&presents, &region) {
            valid_regions += 1;
        }
    }
    println!("You can fit all the presents in {valid_regions} regions");
}

pub fn part2(_lines: &Vec<String>, _day: usize) {
    println!("As always no part 2 for last day :)");
}
