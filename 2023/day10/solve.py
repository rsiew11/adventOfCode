valid_N = {}
valid_S = {}
valid_E = {}
valid_W = {}


def read():
    with open("input.txt") as fd:
        input = fd.readlines()
    return [[x for x in line.split()] for line in input]



def findStart(grid):
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == 'S':
                return (i,j)
        

def part1(grid):
    start = findStart(grid)

    pass

def part2(grid):
    pass

print("part1: ", part1(read()))
print("part2: ", part2(read()))