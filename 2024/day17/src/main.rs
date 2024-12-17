use day17::cpu;
use std::{env, fs::read_to_string};

fn parse_input(filename: &str) -> cpu::Cpu {
    let file = read_to_string(filename).unwrap();
    let lines = file.lines().collect::<Vec<&str>>();
    let reg_a = lines[0].split(": ").collect::<Vec<&str>>()[1]
        .parse::<usize>()
        .unwrap();
    let reg_b = lines[1].split(": ").collect::<Vec<&str>>()[1]
        .parse::<usize>()
        .unwrap();
    let reg_c = lines[2].split(": ").collect::<Vec<&str>>()[1]
        .parse::<usize>()
        .unwrap();
    let prog = lines[4].split(": ").collect::<Vec<&str>>()[1]
        .split(",")
        .map(|n| n.parse::<usize>().unwrap())
        .collect::<Vec<usize>>();
    cpu::Cpu::new(reg_a, reg_b, reg_c, prog)
}

fn part1(cpu: &mut cpu::Cpu) {
    cpu.run();
    print!("The output of the program is: ");
    cpu.print_out();
}

// Of course brute-force approach is not feasible...
fn part2(cpu: &mut cpu::Cpu) {
    let mut generate_curr_digit = Vec::new();

    // 1024 == 2^10, find all the possible values that generate 2 as the first
    // digit and then try to iteratively find all the others.
    for i in 0..1024 {
        cpu.rewind(i);
        cpu.run();
        if cpu.out[0] == cpu.prog[0] {
            generate_curr_digit.push(i);
        }
    }

    // Every 10 bits influence a single digit in the output
    let mut curr_char = 0;
    while curr_char < cpu.prog.len() {
        let mut generate_next_digit = Vec::new();
        for value in generate_curr_digit {
            for i in 0..8 {
                let num = (i << (7 + 3 * curr_char)) | value;
                cpu.rewind(num);
                cpu.run();
                // We can check that the previous char hasn't changed by
                // manipulating the current shifted 10 bits
                // assert!(cpu.out[curr_char - 1] == cpu.prog[curr_char - 1]);
                if cpu.out.len() > curr_char && cpu.out[curr_char] == cpu.prog[curr_char] {
                    generate_next_digit.push(num);
                }
            }
        }
        generate_curr_digit = generate_next_digit;
        curr_char += 1;
    }

    // `generate_curr_digit` contains all the correct solutions plus the ones
    // that have the program itself as prefix and then generate few other digits.
    // Need to filter to consider only the ones that correspond with the actual
    // program.
    let results = generate_curr_digit
        .into_iter()
        .filter(|&a| {
            cpu.rewind(a);
            cpu.run();
            cpu.has_replicated()
        })
        .collect::<Vec<usize>>();
    let reg_a = results.into_iter().min().unwrap();
    println!("The value of A register to replicate itself is: {reg_a}");
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let mut cpu = parse_input(&args[1]);
    part1(&mut cpu);
    part2(&mut cpu);
}
