use std::{
    cmp,
    collections::{HashSet, VecDeque},
};

type Point = (i64, i64);
// type Edge = (i64, i64, i64, i64);
// type Rectangle = (Point, Point, Point, Point);

fn parse_input(lines: &Vec<String>) -> Vec<Point> {
    lines
        .iter()
        .map(|line| {
            let fields: Vec<_> = line.split(",").collect();
            (
                fields[0].parse::<i64>().unwrap(),
                fields[1].parse::<i64>().unwrap(),
            )
        })
        .collect()
}

pub fn part1(lines: &Vec<String>, _day: usize) {
    let red_tiles = parse_input(&lines);
    let mut max_rectangle_area = 0;
    let n = red_tiles.len();

    for i in 0..n - 1 {
        for j in i + 1..n {
            let tile1 = red_tiles[i];
            let tile2 = red_tiles[j];
            let area = (tile1.0 - tile2.0 + 1).abs() * (tile1.1 - tile2.1 + 1).abs();
            max_rectangle_area = cmp::max(area, max_rectangle_area);
        }
    }

    println!("The biggest rectangle of red tiles consists of {max_rectangle_area} tiles");
}

// generate all possible rectangles using the vertices (red tiles)
// bug?????
// fn generate_rectangles(red_tiles: &Vec<Point>) -> Vec<(Rectangle, i64)> {
//     let n = red_tiles.len();
//     let mut rectangles = vec![];
//     for i in 0..n - 1 {
//         for j in i + 1..n {
//             let tile1 = red_tiles[i];
//             let tile2 = red_tiles[j];
//
//             rectangles.push((
//                 (tile1, tile2, (tile1.0, tile2.1), (tile2.0, tile1.1)),
//                 (tile1.0 - tile2.0 + 1).abs() * (tile1.1 - tile2.1 + 1).abs(),
//             ));
//         }
//     }
//     rectangles.sort_by_key(|r| r.1);
//     rectangles.reverse();
//     rectangles
// }

// Ray casting v.1
// fn angle2d(x1: f64, y1: f64, x2: f64, y2: f64) -> f64 {
//     let th1 = y1.atan2(x1);
//     let th2 = y2.atan2(x2);
//     let mut delta_th = th2 - th1;
//     while delta_th > PI {
//         delta_th -= 2.0 * PI;
//     }
//     while delta_th < -PI {
//         delta_th += 2.0 * PI;
//     }
//     delta_th
// }
// fn is_inside(vertices: &Vec<Point>, p: Point) -> bool {
//     let mut angle = 0.0;
//     let n = vertices.len();
//     for i in 0..n {
//         let p1x = vertices[i].0 - p.0;
//         let p1y = vertices[i].1 - p.1;
//         let p2x = vertices[(i + 1) % n].0 - p.0;
//         let p2y = vertices[(i + 1) % n].1 - p.1;
//         angle += angle2d(p1x as f64, p1y as f64, p2x as f64, p2y as f64);
//     }
//     angle.abs() >= PI
// }
// // Ray casting v.2
// fn intersect(x: f64, y: f64, xi: f64, yi: f64, xj: f64, yj: f64) -> bool {
//     ((yi > y) != (yj > y)) && ((x) < ((xj - xi) * (y - yi)) / (yj - yi) + xi)
// }
// fn is_inside(vertices: &Vec<Point>, p: Point) -> bool {
//     let mut inside = false;
//     let mut i = 0 as usize;
//     let mut j = 0 as usize;
//
//     let (x, y) = p;
//     while i < vertices.len() {
//         let (xi, yi) = vertices[i];
//         let (xj, yj) = vertices[j];
//         if intersect(
//             x as f64, y as f64, xi as f64, yi as f64, xj as f64, yj as f64,
//         ) {
//             inside = !inside;
//         }
//         j = i;
//         i += 1;
//     }
//     inside
// }
// // Ray casting v.3
// fn is_inside(vertices: &Vec<Point>, p: Point) -> bool {
//     let n = vertices.len();
//     let mut inside = false;
//     let mut p1 = vertices[0];
//
//     let (x, y) = p;
//     for i in 1..n + 1 {
//         let p2 = vertices[i % n];
//         if y > cmp::min(p1.1, p2.1) {
//             if y <= cmp::max(p1.1, p2.1) {
//                 if x <= cmp::max(p1.0, p2.0) {
//                     let intersect = (y as f64 - p1.1 as f64) * (p2.0 as f64 - p1.0 as f64)
//                         / (p2.1 as f64 - p1.1 as f64)
//                         + p1.0 as f64;
//                     if p1.0 == p2.0 || ((x as f64) < intersect) {
//                         inside = !inside;
//                     }
//                 }
//             }
//         }
//         p1 = p2;
//     }
//     inside
// }

// 21515 is problematic but wrong
// 21598 is problematic but wrong
// 21941 is problematic but wrong
// 22017 is problematic but wrong... etc.. many more...
// pub fn part2(lines: &Vec<String>, _day: usize) {
//     let red_tiles = parse_input(&lines);
//     let rectangles = generate_rectangles(&red_tiles);
//     let mut area_max_within_borders = 0;
//
//     let n = rectangles.len();
//     let mut idx = 48455;
//     while idx < n {
//         let rectangle = rectangles[idx];
//         // for (idx, rectangle) in rectangles.iter().enumerate() {
//         println!("{}/{} -- {}", idx, rectangles.len(), rectangle.1);
//         // println!(">>>>>>>>>> {:?}", rectangle);
//         let minx = cmp::min(
//             rectangle.0.0.0,
//             cmp::min(rectangle.0.1.0, cmp::min(rectangle.0.2.0, rectangle.0.3.0)),
//         );
//         let miny = cmp::min(
//             rectangle.0.0.1,
//             cmp::min(rectangle.0.1.1, cmp::min(rectangle.0.2.1, rectangle.0.3.1)),
//         );
//         let maxx = cmp::max(
//             rectangle.0.0.0,
//             cmp::max(rectangle.0.1.0, cmp::max(rectangle.0.2.0, rectangle.0.3.0)),
//         );
//         let maxy = cmp::max(
//             rectangle.0.0.1,
//             cmp::max(rectangle.0.1.1, cmp::max(rectangle.0.2.1, rectangle.0.3.1)),
//         );
//         println!("minx={minx} maxx={maxx} -- miny={miny} maxy={maxy}");
//
//         let mut found = true;
//         // check all points inside... maybe too much?
//         'points_in_rect: for i in minx..maxx + 1 {
//             if i % 5000 == 0 {
//                 println!(">>> rows={i}/{maxx}");
//             }
//             for j in miny..maxy + 1 {
//                 // if (i == minx && j == miny)
//                 //     || (i == minx && j == maxy)
//                 //     || (i == maxx && j == miny)
//                 //     || (i == maxx && j == maxy)
//                 // {
//                 // if i == minx || i == maxx || j == miny || j == maxy {
//                 // continue;
//                 // }
//                 // println!("-- checking: {i},{j}");
//
//                 if !is_inside(&red_tiles, (i, j))
//                     && (i != minx && i != maxx)
//                     && (j != miny && j != maxy)
//                 {
//                     // println!("{i},{j} is outside!");
//                     found = false;
//                     break 'points_in_rect;
//                 }
//             }
//         }
//         if found {
//             area_max_within_borders = rectangle.1;
//             break;
//         }
//         idx += 1;
//     }
//
//     println!(
//         "The largest rectangle area consisting only of red and green tiles is {area_max_within_borders}"
//     );
//
//     // WHILE THIS ALGORITHM SEEMS SOUND, THE IMPLEMENTATION I PROVIDED WAS NOT :( IT WAS NOT
//     // EFFICIENT ENOUGH AND IT MISSED CLEARLY SOME EDGE CASES, BECAUSE ON THE CORRECT OUTPUT
//     // IT FAILED...
//     // generate all rectangles using red_tiles sorted by area
//     // for each rectangle:
//     //   for each internal point of the rectangle: (just exclude vertices)
//     //     if point is external (ray casting)
//     //       go to next rectangle
//     //   area_max = rectangle.area()
//     //   break
// }

// get all the edges that constitute the perimeter in the form (minx, miny, maxx, maxy)
// fn get_perimeter(vertices: &Vec<Point>) -> Vec<Edge> {
//     let n = vertices.len();
//     let mut edges = vec![];
//     for i in 0..n - 1 {
//         let p1 = vertices[i];
//         let p2 = vertices[i + 1];
//         edges.push((
//             cmp::min(p1.0, p2.0),
//             cmp::min(p1.1, p2.1),
//             cmp::max(p1.0, p2.0),
//             cmp::max(p1.1, p2.1),
//         ));
//     }
//     // connect first and last
//     edges.push((
//         cmp::min(vertices[n - 1].0, vertices[0].0),
//         cmp::min(vertices[n - 1].1, vertices[0].1),
//         cmp::max(vertices[n - 1].0, vertices[0].0),
//         cmp::max(vertices[n - 1].1, vertices[0].1),
//     ));
//     edges
// }
// // https://observablehq.com/@jwolondon/advent-of-code-2025-day-9
// fn intersect(edges: &Vec<Edge>, minx: i64, miny: i64, maxx: i64, maxy: i64) -> bool {
//     for &(edge_minx, edge_miny, edge_maxx, edge_maxy) in edges {
//         if edge_minx > minx && edge_minx < maxx && edge_miny > miny && edge_maxy < maxy {
//             return true;
//         }
//
//         if edge_miny == edge_maxy {
//             if edge_miny > miny
//                 && edge_miny < maxy
//                 && cmp::max(edge_minx, edge_maxx) >= minx
//                 && cmp::min(edge_minx, edge_maxx) <= maxx
//             {
//                 return true;
//             }
//         } else {
//             if edge_minx > minx
//                 && edge_minx < maxx
//                 && cmp::max(edge_miny, edge_maxy) >= miny
//                 && cmp::min(edge_miny, edge_maxy) <= maxy
//             {
//                 return true;
//             }
//         }
//     }
//     false
// }
//
// probably correct, it will takes longer than the Sun to extinguish....
// pub fn part2(lines: &Vec<String>, _day: usize) {
//     let red_tiles = parse_input(&lines);
//     let rectangles = generate_rectangles(&red_tiles);
//     let perimeter = get_perimeter(&red_tiles);
//     let mut area_max_within_borders = 0;
//
//     for rectangle in rectangles.clone() {
//         let minx = cmp::min(
//             rectangle.0.0.0,
//             cmp::min(rectangle.0.1.0, cmp::min(rectangle.0.2.0, rectangle.0.3.0)),
//         );
//         let miny = cmp::min(
//             rectangle.0.0.1,
//             cmp::min(rectangle.0.1.1, cmp::min(rectangle.0.2.1, rectangle.0.3.1)),
//         );
//         let maxx = cmp::max(
//             rectangle.0.0.0,
//             cmp::max(rectangle.0.1.0, cmp::max(rectangle.0.2.0, rectangle.0.3.0)),
//         );
//         let maxy = cmp::max(
//             rectangle.0.0.1,
//             cmp::max(rectangle.0.1.1, cmp::max(rectangle.0.2.1, rectangle.0.3.1)),
//         );
//
//         if !intersect(&perimeter, minx, miny, maxx, maxy) {
//             area_max_within_borders = rectangle.1;
//             break;
//         }
//     }
//
//     println!(
//         "The largest rectangle area consisting only of red and green tiles is {area_max_within_borders}"
//     );
// }

fn build_prefix_sum<F>(rows: usize, cols: usize, func: F) -> Vec<Vec<i64>>
where
    F: Fn(usize, usize) -> i64,
{
    let mut pfx_sum = vec![vec![0; cols]; rows];
    for x in 0..rows {
        for y in 0..cols {
            let left = if x > 0 { pfx_sum[x - 1][y] } else { 0 };
            let top = if y > 0 { pfx_sum[x][y - 1] } else { 0 };
            let topleft = if x > 0 && y > 0 {
                pfx_sum[x - 1][y - 1]
            } else {
                0
            };
            pfx_sum[x][y] = left + top - topleft + func(x, y);
        }
    }
    pfx_sum
}

// if the compressed area is the same as the actual area of the rectangle it means that it is
// fully contained by the perimeter and it is a valid one
fn is_rectangle_valid(
    pfx_sum: &Vec<Vec<i64>>,
    xs: &Vec<i64>,
    ys: &Vec<i64>,
    x1: i64,
    y1: i64,
    x2: i64,
    y2: i64,
) -> bool {
    let mut compressed_xs = [
        xs.iter().position(|x| *x == x1).unwrap() * 2,
        xs.iter().position(|x| *x == x2).unwrap() * 2,
    ];
    compressed_xs.sort();
    let mut compressed_ys = [
        ys.iter().position(|y| *y == y1).unwrap() * 2,
        ys.iter().position(|y| *y == y2).unwrap() * 2,
    ];
    compressed_ys.sort();

    let left = if compressed_xs[0] > 0 {
        pfx_sum[compressed_xs[0] - 1][compressed_ys[1]]
    } else {
        0
    };
    let top = if compressed_ys[0] > 0 {
        pfx_sum[compressed_xs[1]][compressed_ys[0] - 1]
    } else {
        0
    };
    let topleft = if compressed_xs[0] > 0 && compressed_ys[0] > 0 {
        pfx_sum[compressed_xs[0] - 1][compressed_ys[0] - 1]
    } else {
        0
    };
    let compressed_area = pfx_sum[compressed_xs[1]][compressed_ys[1]] - left - top + topleft;
    let full_area = ((compressed_xs[1] - compressed_xs[0] + 1)
        * (compressed_ys[1] - compressed_ys[0] + 1)) as i64;
    compressed_area == full_area
}

// flood fill to determine inside points: the general idea is that if a point is not outside
// of the perimeter then it must be inside
fn flood_fill(compressed_grid: &Vec<Vec<i64>>) -> HashSet<Point> {
    let mut q = VecDeque::new();
    let mut outside = HashSet::new();
    q.push_back((-1, -1));
    outside.insert((-1, -1));

    while !q.is_empty() {
        let (cx, cy) = q.pop_front().unwrap();
        for (xx, yy) in [(cx + 1, cy), (cx - 1, cy), (cx, cy - 1), (cx, cy + 1)] {
            // in range considering a buffer of 1 block after the end
            if xx < -1
                || yy < -1
                || xx > compressed_grid.len() as i64
                || yy > compressed_grid[0].len() as i64
            {
                continue;
            }
            // hit the wall
            if xx >= 0
                && xx < compressed_grid.len() as i64
                && yy >= 0
                && yy < compressed_grid[0].len() as i64
                && compressed_grid[xx as usize][yy as usize] == 1
            {
                continue;
            }
            let p = (xx, yy);
            if outside.contains(&p) {
                continue;
            }
            outside.insert(p);
            q.push_back(p);
        }
    }
    outside
}

fn coordinate_compression(red_tiles: &Vec<Point>, xs: &Vec<i64>, ys: &Vec<i64>) -> Vec<Vec<i64>> {
    let mut compressed_grid = Vec::new();
    for _ in 0..xs.len() * 2 - 1 {
        let mut row = Vec::new();
        for _ in 0..ys.len() * 2 - 1 {
            row.push(0);
        }
        compressed_grid.push(row);
    }

    for ((x1, y1), (x2, y2)) in red_tiles
        .iter()
        .zip([&red_tiles[1..], &red_tiles[..1]].concat().iter())
    {
        let mut compressed_xs = [
            xs.iter().position(|x| x == x1).unwrap() * 2,
            xs.iter().position(|x| x == x2).unwrap() * 2,
        ];
        compressed_xs.sort();
        let mut compressed_ys = [
            ys.iter().position(|y| y == y1).unwrap() * 2,
            ys.iter().position(|y| y == y2).unwrap() * 2,
        ];
        compressed_ys.sort();

        for cx in compressed_xs[0]..compressed_xs[1] + 1 {
            for cy in compressed_ys[0]..compressed_ys[1] + 1 {
                compressed_grid[cx][cy] = 1;
            }
        }
    }
    compressed_grid
}

pub fn part2(lines: &Vec<String>, _day: usize) {
    let red_tiles = parse_input(&lines);
    let mut xs_set = HashSet::new();
    let mut ys_set = HashSet::new();

    for rt in &red_tiles {
        xs_set.insert(rt.0);
        ys_set.insert(rt.1);
    }

    let mut xs = xs_set.into_iter().collect::<Vec<i64>>();
    xs.sort();
    let mut ys = ys_set.into_iter().collect::<Vec<i64>>();
    ys.sort();

    let mut compressed_grid = coordinate_compression(&red_tiles, &xs, &ys);

    // fill the perimeter of the compressed grid, basically consider all red and green tiles
    let outside = flood_fill(&compressed_grid);
    for xx in 0..compressed_grid.len() {
        for yy in 0..compressed_grid[0].len() {
            if !outside.contains(&(xx as i64, yy as i64)) {
                compressed_grid[xx][yy] = 1;
            }
        }
    }

    let red_green_compressed_grid = build_prefix_sum(
        compressed_grid.len(),
        compressed_grid[0].len(),
        |x: usize, y: usize| compressed_grid[x][y],
    );

    let mut max_area = 0;
    for (i, &(x1, y1)) in red_tiles.iter().enumerate() {
        for &(x2, y2) in red_tiles[..i].into_iter() {
            if is_rectangle_valid(&red_green_compressed_grid, &xs, &ys, x1, y1, x2, y2) {
                let area = ((x1 - x2).abs() + 1) * ((y1 - y2).abs() + 1);
                max_area = cmp::max(max_area, area);
            }
        }
    }

    println!(
        "The area of the biggest rectangle of red/green tiles inside the perimeter is {max_area} m^2"
    );
}
