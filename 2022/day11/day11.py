class Queue:
    def __init__(self):
        self.queue = []
        self.size = 0

    def enqueue(self, el):
        self.queue.append(el)
        self.size += 1

    def dequeue(self):
        el = self.queue[0]
        self.queue = self.queue[1:]
        self.size -= 1
        return el

    def is_empty(self):
        return self.size == 0


class Monkey:
    def __init__(self, itms, opr, tst, tt, dv):
        self.items = itms
        self.op = opr
        self.test = tst
        self.to_throw = tt
        self.div = dv
        self.monkey_business = 0


def read_test_line(line):
    fields = line.split(' ')
    last = len(fields) - 1
    return int(fields[last])


def read_items(line):
    line = line.strip()
    fields = line.split(' ')[2:]

    q = Queue()
    for s in fields:
        off = len(s) - 1
        if s[off] != ',':
            off += 1

        n = int(s[:off])
        q.enqueue(n)

    return q


def read_operation_line(line):
    line = line.strip()
    fields = line.split(' ')[4:]

    n = None
    try:
        n = int(fields[1])
    except:
        n = -1

    if fields[0] == "+":
        if n < 0:
            return lambda x: x + x
        else:
            return lambda x: x + n
    elif fields[0] == "*":
        if n < 0:
            return lambda x: x * x
        else:
            return lambda x: x * n
    else:
        print('Operation not known')
        exit(2)


def parse_input(lines):
    monkeys = []
    for i in range(0, len(lines), 7):
        itms = read_items(lines[i+1])
        opr = read_operation_line(lines[i+2])
        div = read_test_line(lines[i+3])
        tst = eval(f'lambda x: x%{div} == 0')
        tt = {}
        tt[True] = read_test_line(lines[i+4])
        tt[False] = read_test_line(lines[i+5])
        
        monkeys.append(Monkey(itms, opr, tst, tt, div))
    return monkeys


def print_round(monkeys, round):
    print(f'After round {round}')
    for idx, m in enumerate(monkeys):
        print(f'Monkey {idx}: {m.items.queue}\tbusiness: {m.monkey_business}')
    print()


def part1(lines):
    monkeys = parse_input(lines)

    for _ in range(20):
        for idx, m in enumerate(monkeys):
            monkeys[idx].monkey_business += m.items.size
            while not m.items.is_empty():
                item = m.items.dequeue()
                item = m.op(item)
                item //= 3
                key = m.test(item)
                monkeys[m.to_throw[key]].items.enqueue(item)
        # print_round(monkeys, r)

    monkey_businesses = []
    for m in monkeys:
        monkey_businesses.append(m.monkey_business)
    monkey_businesses = sorted(monkey_businesses, reverse=True)

    print(f'The monkey business value is: {monkey_businesses[0] * monkey_businesses[1]}')


def part2(lines):
    monkeys = parse_input(lines)
    lcm = 1
    for m in monkeys:
        lcm *= m.div
    for r in range(10000):
        for idx, m in enumerate(monkeys):
            monkeys[idx].monkey_business += m.items.size
            while not m.items.is_empty():
                item = m.items.dequeue()
                item = m.op(item)
                key = m.test(item)
                monkeys[m.to_throw[key]].items.enqueue(item % lcm)
        if r % 100 == 0:
            print(f'Round {r}...')
        # print_round(monkeys, r)

    monkey_businesses = []
    for m in monkeys:
        monkey_businesses.append(m.monkey_business)
    monkey_businesses = sorted(monkey_businesses, reverse=True)

    print(f'The monkey business value is: {monkey_businesses[0] * monkey_businesses[1]}')


def main():
    lines = None
    with open('mini_input', 'r') as f:
        lines = f.read().splitlines()

    part1(lines)
    part2(lines)


if __name__ == '__main__':
    main()
