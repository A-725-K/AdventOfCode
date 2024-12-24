use std::{
    collections::{HashMap, VecDeque},
    env,
    fs::read_to_string,
};

struct Operation {
    op1: String,
    op2: String,
    sym: String,
    res_reg: String,
}

fn parse_input(filename: &str) -> (HashMap<String, i64>, Vec<Operation>) {
    let mut registers = HashMap::new();
    let mut operations = vec![];
    let mut read_registers = true;
    for line in read_to_string(filename).unwrap().lines() {
        if line == "" {
            read_registers = false;
            continue;
        }
        if read_registers {
            let fields = line.split(": ").map(String::from).collect::<Vec<String>>();
            registers.insert(fields[0].clone(), fields[1].parse::<i64>().unwrap());
        } else {
            let fields = line.split(" ").map(String::from).collect::<Vec<String>>();
            let op = Operation {
                op1: fields[0].clone(),
                op2: fields[2].clone(),
                sym: fields[1].clone(),
                res_reg: fields[4].clone(),
            };
            operations.push(op);
        }
    }
    (registers, operations)
}

fn part1(registers: &mut HashMap<String, i64>, operations: &Vec<Operation>) {
    let mut q = VecDeque::new();
    for op in operations {
        q.push_back(op);
    }

    while let Some(op) = q.pop_front() {
        if registers.contains_key(&op.op1) && registers.contains_key(&op.op2) {
            let op1 = registers.get(&op.op1).unwrap();
            let op2 = registers.get(&op.op2).unwrap();
            let res = if op.sym == "AND" {
                op1 & op2
            } else if op.sym == "OR" {
                op1 | op2
            } else if op.sym == "XOR" {
                op1 ^ op2
            } else {
                panic!("Unknown operation")
            };
            registers.insert(op.res_reg.clone(), res);
        } else {
            q.push_back(op);
        }
    }

    let result = number_from_registers("z", &registers);
    println!("The decimal number produced by the wiring is {result}");
}

fn number_from_registers(letter: &str, registers: &HashMap<String, i64>) -> i64 {
    let mut keys = registers
        .keys()
        .into_iter()
        .filter(|k| k.starts_with(letter))
        .collect::<Vec<&String>>();
    keys.sort();
    let mut result = 0;
    let mut shift = 0;
    for k in keys {
        // println!("{k}: {}", registers.get(k).unwrap());
        let i = registers.get(k).unwrap();
        result = result | (i << shift);
        shift += 1;
    }
    result
}

/// PART 2

fn craft(letter: &str, n: usize) -> String {
    letter.to_string() + format!("{:0>2}", n).as_str()
}

fn verify_immediate_carry(
    reg: String,
    num: usize,
    formulas: &HashMap<String, (String, String, String)>,
) -> bool {
    if !formulas.contains_key(&reg) {
        return false;
    }

    let (op, left, right) = formulas.get(&reg).unwrap();

    // The immediate carry depends only on the two operands and both must be
    // true
    if op != "AND" {
        return false;
    }
    let mut ops = [left, right];
    ops.sort();
    ops == [&craft("x", num), &craft("y", num)]
}

fn verify_carry_backward(
    reg: String,
    num: usize,
    formulas: &HashMap<String, (String, String, String)>,
) -> bool {
    if !formulas.contains_key(&reg) {
        return false;
    }

    let (op, left, right) = formulas.get(&reg).unwrap();

    if op != "AND" {
        return false;
    }
    (verify_intermediate_xor(left.clone(), num, formulas)
        && verify_carry_bit(right.clone(), num, formulas))
        || (verify_intermediate_xor(right.clone(), num, formulas)
            && verify_carry_bit(left.clone(), num, formulas))
}

fn verify_carry_bit(
    reg: String,
    num: usize,
    formulas: &HashMap<String, (String, String, String)>,
) -> bool {
    if !formulas.contains_key(&reg) {
        return false;
    }

    let (op, left, right) = formulas.get(&reg).unwrap();

    // The first carry bit is true only if both x00 and y00 are true
    if num == 1 {
        let mut ops = [left, right];
        ops.sort();
        return op == "AND" && ops == ["x00", "y00"];
    }

    // For all the other bits it has to be an OR between the current carry and
    // the ones that has been carried
    if op != "OR" {
        return false;
    }
    (verify_immediate_carry(left.clone(), num - 1, formulas)
        && verify_carry_backward(right.clone(), num - 1, formulas))
        || (verify_immediate_carry(right.clone(), num - 1, formulas)
            && verify_carry_backward(left.clone(), num - 1, formulas))
}

fn verify_intermediate_xor(
    reg: String,
    num: usize,
    formulas: &HashMap<String, (String, String, String)>,
) -> bool {
    if !formulas.contains_key(&reg) {
        return false;
    }

    let (op, left, right) = formulas.get(&reg).unwrap();
    if op != "XOR" {
        return false;
    }

    let mut ops = [left, right];
    ops.sort();
    return ops == [&craft("x", num), &craft("y", num)];
}

fn verify_z(reg: String, num: usize, formulas: &HashMap<String, (String, String, String)>) -> bool {
    // After swapping some keys might not exists on the right side of the
    // input operations
    if !formulas.contains_key(&reg) {
        return false;
    }

    let (op, left, right) = formulas.get(&reg).unwrap();
    if op != "XOR" {
        return false;
    }

    // Check that the LSB is the XOR between x00 and y00
    if num == 0 {
        let mut ops = [left, right];
        ops.sort();
        return ops == ["x00", "y00"];
    }

    (verify_intermediate_xor(left.clone(), num, formulas)
        && verify_carry_bit(right.clone(), num, formulas))
        || (verify_intermediate_xor(right.clone(), num, formulas)
            && verify_carry_bit(left.clone(), num, formulas))
}

fn fail_at(formulas: &HashMap<String, (String, String, String)>) -> usize {
    let mut idx = 0;
    while idx <= 45 {
        if !verify_z(craft("z", idx), idx, &formulas) {
            break;
        }
        idx += 1;
    }
    return idx;
}

fn swap_values(map: &mut HashMap<String, (String, String, String)>, k1: &String, k2: &String) {
    let a = map.get_mut(k1).unwrap() as *mut (String, String, String);
    let b = map.get_mut(k2).unwrap() as *mut (String, String, String);
    unsafe {
        std::ptr::swap(a, b);
    }
}

fn part2(operations: &Vec<Operation>) {
    // To understand how addition works bit-wise, read here:
    //  https://www.uop.edu.jo/PDF%20File/petra%20university%20Digital_Design_-_A_Comprehensive_Guide_to_Digital_Electronics_and_Computer_System_Architecture-Part18.pdf
    let mut formulas = HashMap::new();
    for op in operations {
        formulas.insert(
            op.res_reg.clone(),
            (op.sym.clone(), op.op1.clone(), op.op2.clone()),
        );
    }
    let mut failed_registers = vec![];
    for _ in 0..4 {
        let failure_idx = fail_at(&formulas);
        for reg1 in formulas.clone().keys() {
            let mut other = String::new();
            for reg2 in formulas.clone().keys() {
                if reg1 == reg2 {
                    continue;
                }

                swap_values(&mut formulas, reg1, reg2);
                if fail_at(&formulas) > failure_idx {
                    other = reg2.to_string();
                    break;
                }
                swap_values(&mut formulas, reg1, reg2);
            }
            if other != "" {
                // println!("{reg1} <=> {other}");
                failed_registers.push(reg1.to_string());
                failed_registers.push(other.to_string());
                break;
            }
        }
    }
    failed_registers.sort();
    println!("The faulty registers are: {}", failed_registers.join(","));
}

pub fn main() {
    let args: Vec<String> = env::args().collect();
    let (mut registers, operations) = parse_input(&args[1]);
    part1(&mut registers, &operations);
    part2(&operations);
}
