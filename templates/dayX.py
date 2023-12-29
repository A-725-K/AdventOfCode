import sys


def parse_input(filename):
    lines = []
    with open(filename, "r") as f:
        lines = list(map(lambda s: s.replace("\n", ""), f.readlines()))
    return lines


def part1(lines):
    pass


def part2(lines):
    pass


def main():
    lines = parse_input(sys.argv[1])
    part1(lines)
    part2(lines)


if __name__ == "__main__":
    main()
