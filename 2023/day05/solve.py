
def read(f):
    with open(f, "r+") as fd:
        return fd.read().strip()

def getDestVal(val, m):
    for entry in m:
        d,s,r = map(int, entry.split())
        if s <= val < s+r:
            return d + (val - s) # dest + how much more is the val above start
    return val

def part1(seeds, maps):
    locations = []
    for seed in seeds:
        val = seed
        for m in maps:
            val = getDestVal(val, m)
        locations.append(val)
    return min(locations)

def part2():
    pass


def main():
    data = read("./input.txt")
    sections = data.split("\n\n") 

    seeds = list(map(int, sections[0].split(":")[1].split()))
    maps = [s.split("\n")[1:] for s in sections[1:]]
    print("part1:", part1(seeds, maps))

    seeds = [(seeds[i], seeds[i+1]) for i in range(0,len(seeds), 2)]
    for s in seeds:
        print(s)



if __name__ == "__main__":
    main()