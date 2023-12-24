import sys
import sympy


class Hailstone:
    def __init__(self, x, y, z, vx, vy, vz):
        self.x, self.y, self.z = x, y, z
        self.vx, self.vy, self.vz = vx, vy, vz

    def __repr__(self):
        return f"x={self.x}, y={self.y}, z={self.z} -- " \
            f"vx={self.vx}, vy={self.vy}, vz={self.vz}"


def parse_input():
    hailstones = []
    with open(sys.argv[1], 'r') as f:
        while line := f.readline():
            hs = map(int, line.replace("@", ",").replace(" ", "").split(","))
            hailstones.append(Hailstone(*hs))
    return hailstones


def gather_equations(hailstones, symbols):
    to_solve = []
    (xr, yr, zr, vxr, vyr, vzr) = symbols
    for hs in hailstones:
        to_solve.append((xr-hs.x)*(hs.vy-vyr) - (yr-hs.y)*(hs.vx-vxr))
        to_solve.append((yr-hs.y)*(hs.vz-vzr) - (zr-hs.z)*(hs.vy-vyr))
    return to_solve


def part1():
    # solved already in Go
    pass


def part2(hailstones):
    xr, yr, zr, vxr, vyr, vzr = sympy.symbols("xr, yr, zr, vxr, vyr, vzr")
    eqs = gather_equations(hailstones, (xr, yr, zr, vxr, vyr, vzr))
    solution = sympy.solve(eqs)
    assert len(solution) == 1
    solution = solution[0]
    result = solution[xr] + solution[yr] + solution[zr]
    print(f"The sum of coordinates of the rock I throw is: {result}", end="")


def main():
    hailstones = parse_input()
    part1()
    part2(hailstones)


if __name__ == "__main__":
    main()
