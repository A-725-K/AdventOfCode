use std::{collections::HashSet, env, fs::read_to_string};

type Garden = Vec<Vec<char>>;
type Region = (char, Vec<(usize, usize)>);

fn parse_input(filename: &str) -> Garden {
    read_to_string(filename)
        .unwrap()
        .lines()
        .map(|line| line.chars().collect())
        .collect()
}

fn get_regions(garden: &Garden, rows: usize, cols: usize) -> Vec<Region> {
    let mut regions = vec![];
    let mut visited: Vec<Vec<bool>> = vec![];
    for _ in 0..rows {
        let mut row = vec![];
        for _ in 0..cols {
            row.push(false);
        }
        visited.push(row);
    }

    for y in 0..rows {
        for x in 0..cols {
            if visited[y][x] {
                continue;
            }
            let curr_symbol = garden[y][x];
            let mut region = vec![];
            let mut to_visit = vec![(x, y)];

            while !to_visit.is_empty() {
                let (curr_x, curr_y) = to_visit.pop().unwrap();
                if visited[curr_y][curr_x] {
                    continue;
                }
                visited[curr_y][curr_x] = true;
                region.push((curr_x, curr_y));

                // North
                if (curr_y as i32) - 1 >= 0
                    && !visited[curr_y - 1][curr_x]
                    && garden[curr_y - 1][curr_x] == curr_symbol
                {
                    to_visit.push((curr_x, curr_y - 1));
                }
                // South
                if curr_y + 1 < rows
                    && !visited[curr_y + 1][curr_x]
                    && garden[curr_y + 1][curr_x] == curr_symbol
                {
                    to_visit.push((curr_x, curr_y + 1));
                }
                // East
                if curr_x + 1 < cols
                    && !visited[curr_y][curr_x + 1]
                    && garden[curr_y][curr_x + 1] == curr_symbol
                {
                    to_visit.push((curr_x + 1, curr_y));
                }
                // West
                if (curr_x as i32) - 1 >= 0
                    && !visited[curr_y][curr_x - 1]
                    && garden[curr_y][curr_x - 1] == curr_symbol
                {
                    to_visit.push((curr_x - 1, curr_y));
                }
            }
            regions.push((curr_symbol, region));
        }
    }
    regions
}

fn compute_region_cost(region: &Region, garden: &Garden, rows: usize, cols: usize) -> usize {
    let area = region.1.len();
    let mut perimeter = 4 * area;
    for (sector_x, sector_y) in region.1.clone().into_iter() {
        // North
        if (sector_y as i32) - 1 >= 0 && garden[sector_y - 1][sector_x] == region.0 {
            perimeter -= 1;
        }
        // South
        if sector_y + 1 < rows && garden[sector_y + 1][sector_x] == region.0 {
            perimeter -= 1;
        }
        // East
        if sector_x + 1 < cols && garden[sector_y][sector_x + 1] == region.0 {
            perimeter -= 1;
        }
        // West
        if (sector_x as i32) - 1 >= 0 && garden[sector_y][sector_x - 1] == region.0 {
            perimeter -= 1;
        }
    }
    area * perimeter
}

fn part1(garden: &Garden, rows: usize, cols: usize) {
    let mut fences_cost = 0;
    for region in get_regions(garden, rows, cols) {
        fences_cost += compute_region_cost(&region, garden, rows, cols);
    }
    println!("To fence all regions in the garden the cost is {fences_cost}$");
}

fn compute_region_cost_with_bulk_discount(reg: &Region) -> usize {
    let region: Vec<(i32, i32)> = reg
        .1
        .clone()
        .into_iter()
        .map(|(x, y)| (x as i32, y as i32))
        .collect();
    let area = region.len();

    // Observation 1: the number of segments in the outline, is the same as the
    //                number of corners in the polygon.
    // Observation 2: if a vertex appears an odd number of times in a region,
    //                it is a corner.
    let mut vertices = HashSet::new();
    for (sector_x, sector_y) in region.to_owned() {
        for dir in [(-1, -1), (1, -1), (1, 1), (-1, 1)] {
            let vertex = ((sector_x * 2) as i32 + dir.0, (sector_y * 2) as i32 + dir.1);
            vertices.insert(vertex);
        }
    }
    // Not enough... it is missing few cases...
    // let size = vertices
    //     .values()
    //     .into_iter()
    //     .filter(|&n| *n % 2 != 0)
    //     .count();

    let mut sizes = 0;
    for (vx, vy) in vertices.clone() {
        // Find the squares close to the current vertex
        let mut close_squares = vec![];
        for dir in [(-1, -1), (1, -1), (1, 1), (-1, 1)] {
            let p = ((vx + dir.0) / 2, (vy + dir.1) / 2);
            close_squares.push(region.contains(&p));
        }
        let corner_count: usize = close_squares.clone().into_iter().map(|e| e as usize).sum();
        if corner_count == 1 || corner_count == 3 {
            sizes += 1;
        } else if corner_count == 2
            && (close_squares == [true, false, true, false]
                || close_squares == [false, true, false, true])
        {
            sizes += 2;
        }
    }
    sizes * area
}

fn part2(garden: &Garden, rows: usize, cols: usize) {
    let mut fences_cost = 0;
    for region in get_regions(garden, rows, cols) {
        fences_cost += compute_region_cost_with_bulk_discount(&region);
    }
    println!("To fence all regions in the garden with bulk discount the cost is {fences_cost}$");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let garden = parse_input(&args[1]);
    let (rows, cols) = (garden.len(), garden[0].len());
    part1(&garden, rows, cols);
    part2(&garden, rows, cols);
}
