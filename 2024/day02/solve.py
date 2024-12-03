
def read():
    with open("./input.txt") as fd:
        data = fd.readlines()
    return [list(map(int, line.split(' '))) for line in data]

def p1():
    lines = read()
    inc = {1,2,3}
    dec = {-1,-2,-3}
    safe = 0
    for line in lines:
        diffs = [b-a for a,b in zip(line, line[1:])]
        if   all(diff in inc for diff in diffs): safe += 1
        elif all(diff in dec for diff in diffs): safe += 1
    return safe

def p2():
    lines = read()
    inc, dec = {1,2,3}, {-1,-2,-3}
    safe = 0

    def isSafe(line):
        diffs = [b-a for a,b in zip(line, line[1:])]
        return (all(diff in inc for diff in diffs) or 
                all(diff in dec for diff in diffs))

    for line in lines:
        if isSafe(line):
            safe += 1
            continue

        variations = [line[:i] + line[i+1:] for i in range(len(line))]
        if any(isSafe(var) for var in variations): safe += 1

    return safe

if __name__ == "__main__":
    print(p1())
    print(p2())

