
def read():
    with open("input.txt") as fd:
        input = fd.readlines()
    return [[int(x) for x in line.split()] for line in input]


def getExtraps(line):
    extraps = [[line[i+1] - line[i] for i in range(len(line)-1)]]
    while not all(v == 0 for v in extraps[-1]):
        cur = extraps[-1]
        extraps.append([cur[i+1] - cur[i] for i in range(len(cur)-1)])
    return extraps

def part1(data):
    res = 0
    for line in data:
        extraps = getExtraps(line)
        last = 0
        for i in range(len(extraps)-1, -1, -1):
            last = last + extraps[i][-1]
        last = last + line[-1]
        res += last 
        
    return res

def part2(data):
    res = 0
    for line in data:
        extraps = getExtraps(line)
        first = 0
        for i in range(len(extraps)-1, -1, -1):
            first = extraps[i][0] - first
        first = line[0] - first
        res += first 
        
    return res

print("part1: ", part1(read()))
print("part2: ", part2(read()))