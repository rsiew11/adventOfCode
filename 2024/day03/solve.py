import re

def readRaw():
    with open("./input.txt") as fd:
        data = fd.readlines()
    line = "".join(data).replace("\n", "")
    return line

def read():
    valid = set(map(str, range(10)))
    valid.add(",")
    line = readRaw()
    pairs = []
    for i in range(len(line)):
        if line[i:i+4] != "mul(": continue
        l = i+4
        r = l
        while line[r] in valid:
            r += 1
        if r < len(line) and line[r] == ")":
            parts = line[l:r].split(",")
            if len(parts) == 2:
                pairs.append((i, r, int(parts[0]), int(parts[1])))
    return pairs


def p1():
    pairs = read()
    return sum(x*y for _,_,x,y in pairs)


def p2():
    data = read()
    line = readRaw()
    do = True
    intervals = []
    l,r = 0,0
    for i in range(len(line)): 
        if do == True:
            if line[i:i+7] == "don't()":
                r = i+7
                do = False
                intervals.append((l,r))
        else:
            if line[i:i+4] == "do()":
                l = i
                do = True
    if do == True:
        intervals.append((l, len(line)))
    res = 0
    for lb,rb in intervals:
        for l,_,x,y in data:
            if lb < l < rb:
                res += x*y
    return res
        





if __name__ == "__main__":
    print("part1",p1())
    print("part2", p2())