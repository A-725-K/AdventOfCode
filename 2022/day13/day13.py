import functools


LT, EQ, GT = -1, 0, 1


def cmp(l, r):
    if l < r:
        return LT
    elif l > r:
        return GT
    return EQ


def comparePackets(left, right):
    if type(left) == int:
        if type(right) == int:
            # print(f'1 Comparing {left} ** {right}')
            return cmp(left, right)
        # print(f'2 Comparing {left} ** {right}')
        return comparePackets([left], right)

    if type(right) == int:
        # print(f'3 Comparing {left} ** {right}')
        return comparePackets(left, [right])
    
    # print(f'4 Comparing {left} ** {right}')
    for r_idx, rr in enumerate(right):
        if r_idx >= len(left):
            break
        res = comparePackets(left[r_idx], rr)
        if res != EQ:
            return res

    return cmp(len(left), len(right))

    
def part1(lines):
    index, packetsInOrder = 1, []

    l, r = None, None
    for i, line in enumerate(lines):
        if i%3 == 0:
            l = eval(line)
        elif i%3 == 1:
            r = eval(line)
        else:
            # print(f'-------- {index}')
            if comparePackets(l, r) < 0:
                packetsInOrder.append(index)
            index += 1

    print(f'The sum of the indexes of in-order packets is: {sum(packetsInOrder)}')


def part2(lines):
    lines = list(map(lambda l: eval(l), filter(lambda x: x != '', lines)))
    lines.append([2])
    lines.append([6])
    lines = sorted(lines, key=functools.cmp_to_key(comparePackets))

    decoder_key = (lines.index([2]) + 1) * (lines.index([6]) + 1)

    print(f'The decoder key for the message is: {decoder_key}')


def main():
    lines = []
    with open('input', 'r') as f:
        lines = f.read().splitlines()

    part1(lines)
    part2(lines)


if __name__ == '__main__':
    main()

