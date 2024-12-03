from collections import Counter

def read():
    with open("./input.txt") as fd:
        data = fd.readlines()
    l1 = []
    l2 = []
    for line in data:
        v1, v2 = line.split()
        l1.append(int(v1))
        l2.append(int(v2))
    return sorted(l1), sorted(l2)

def p1():
    l1, l2 = read()
    return sum(abs(v1-v2) for v1,v2 in zip(l1,l2))

def p2():
    l1, l2 = read()
    l2 = Counter(l2)
    return sum(v * l2[v] for v in l1)

if __name__ == "__main__":
    print(p1())
    print(p2())

