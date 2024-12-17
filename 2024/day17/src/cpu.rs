#[derive(Debug, Clone)]
pub struct Cpu {
    pc: usize,
    reg_a: usize,
    reg_b: usize,
    reg_c: usize,

    pub prog: Vec<usize>,
    pub out: Vec<usize>,
}

impl Cpu {
    pub fn new(reg_a: usize, reg_b: usize, reg_c: usize, prog: Vec<usize>) -> Cpu {
        Cpu {
            pc: 0,
            reg_a,
            reg_b,
            reg_c,
            prog: prog.clone(),
            out: vec![],
        }
    }
    pub fn run(&mut self) {
        while self.pc < self.prog.len() {
            let next_operation = self.prog[self.pc];
            let next_operand = self.prog[self.pc + 1];
            let mut has_jumped = false;

            match next_operation {
                0 => self.adv(next_operand),
                1 => self.bxl(next_operand),
                2 => self.bst(next_operand),
                3 => has_jumped = self.jnz(next_operand),
                4 => self.bxc(next_operand),
                5 => self.out(next_operand),
                6 => self.bdv(next_operand),
                7 => self.cdv(next_operand),
                _ => panic!("Unknown instruction code: '{next_operation}'"),
            }

            if !has_jumped {
                self.pc += 2;
            }
        }
    }
    pub fn print_out(&self) {
        println!(
            "{}",
            self.out
                .clone()
                .into_iter()
                .map(|n| format!("{}", n))
                .collect::<Vec<String>>()
                .join(",")
        );
    }
    pub fn rewind(&mut self, reg_a: usize) {
        self.pc = 0;
        self.out = vec![];
        self.reg_a = reg_a;
        self.reg_b = 0;
        self.reg_c = 0;
    }
    pub fn has_replicated(&self) -> bool {
        self.out.len() == self.prog.len()
            && self
                .out
                .iter()
                .zip(self.prog.iter())
                .filter(|&(a, b)| a == b)
                .count()
                == self.prog.len()
    }

    fn get_combo(&self, op: usize) -> usize {
        if op <= 3 {
            return op;
        }
        if op == 4 {
            return self.reg_a;
        }
        if op == 5 {
            return self.reg_b;
        }
        if op == 6 {
            return self.reg_c;
        }
        panic!("Opcode '{op}' not valid, cannot execute this operation!");
    }

    fn adv(&mut self, op: usize) {
        self.reg_a /= (2 as u32).pow(self.get_combo(op) as u32) as usize;
    }
    fn bxl(&mut self, op: usize) {
        self.reg_b ^= op;
    }
    fn bst(&mut self, op: usize) {
        self.reg_b = self.get_combo(op) % 8;
    }
    fn jnz(&mut self, op: usize) -> bool {
        if self.reg_a == 0 {
            return false;
        }
        self.pc = op;
        true
    }
    fn bxc(&mut self, _: usize) {
        self.reg_b ^= self.reg_c;
    }
    fn out(&mut self, op: usize) {
        self.out.push(self.get_combo(op) % 8);
    }
    fn bdv(&mut self, op: usize) {
        self.reg_b = self.reg_a / (2 as u32).pow(self.get_combo(op) as u32) as usize;
    }
    fn cdv(&mut self, op: usize) {
        self.reg_c = self.reg_a / (2 as u32).pow(self.get_combo(op) as u32) as usize;
    }
}
