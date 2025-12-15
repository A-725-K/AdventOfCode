from z3 import *

get_ints = lambda n, ln: Ints(' '.join([f'x_{ln}_{i}' for i in range(n)]))

s = Solver()
result = 0
for li, line in enumerate(open('src/inputs/day10/input').read().split("\n")):
    if line == "": continue
    fields = line.strip().split(" ")[1:]
    switches = [eval(el.replace('(', '[').replace(')', ']')) for el in fields[:-1]]
    joltages = eval(fields[-1].replace('{', '(').replace('}', ')'))
    # print("[DEBUG]: switches:", switches)
    # print("[DEBUG]: joltages:", joltages)

    s.reset()
    xs = get_ints(len(switches), li)
    s.add([ x >= 0 for x in xs])

    # add upper limit for each button
    for j, switch in enumerate(switches):
        min_val = min([joltages[el] for el in switch])
        s.add(xs[j] <= min_val)

    # add equations that combines buttons that increment same joltage
    for i, joltage in enumerate(joltages):
        idxs = [idx for idx, switch in enumerate(switches) if i in switch]
        s.add(Sum([xs[i] for i in idxs]) == joltage)

    # print("[DEBUG]:", s.assertions())

    # iterate over solutions and choose the minimum
    min_sol = None
    while s.check() == sat:
        model = s.model()
        sol = [model[x].as_long() for x in xs]
        min_sol = sum(sol)
        # print(f"[DEBUG]: {sol} --> {min_sol}")

        s.add(Or([x != model[x] for x in xs]))
        s.add(Sum(xs) < min_sol)
    result += min_sol
print(result)

# Example:
# s.add(xs[0] <= 7)
# s.add(xs[1] <= 4)
# s.add(xs[2] <= 5)
# s.add(xs[3] <= 5)
# s.add(xs[4] <= 3)
# s.add(xs[5] <= 3)
# s.add(
#     xs[4] + xs[5] == 3,
#     xs[1] + xs[5] == 4,
#     xs[2] + xs[3] + xs[4] == 5,
#     xs[0] + xs[1] + xs[3] == 7
# )
