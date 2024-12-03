from collections import deque
from itertools import count

DIRS = [(0,1),(0,-1),(1,0),(-1,0)]

def read():
    with open("input.txt") as fd:
        raw = fd.readlines()
    raw = [list(line.rstrip()) for line in raw]
    return raw


def findStart(grid):
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == 'S':
                return (i,j)
        

def part1(grid):
    rows = len(grid)
    cols = len(grid[0])
    valid = {
        "|": [(1,0),(-1,0)], #down, up
        "-": [(0,-1),(0,1)], #left, right
        "L": [(-1,0),(0,1)], #up, right
        "J": [(0,-1),(-1,0)],#left, up
        "7": [(0,-1),(1,0)], #left, down
        "F": [(0,1),(1,0)],  #right, down
    }
    start_x, start_y = findStart(grid)
    visited = {(start_x, start_y): 0}
    work = deque()
    if (grid[start_x-1][start_y]in {"|", "7", "F"}): #up
        work.append((1, (start_x-1, start_y)))
    if (grid[start_x+1][start_y] in {"|", "L", "J"}): #down
        work.append((1, (start_x+1, start_y)))
    if (grid[start_x][start_y-1] in {"-", "L", "F"}): #left
        work.append((1, (start_x, start_y-1)))
    if (grid[start_x][start_y+1] in {"-", "J", "7"}): #right
        work.append((1, (start_x, start_y+1)))

    while work:
        dist, (x, y) = work.popleft()
        if (x,y) in visited: continue
        visited[(x,y)] = dist
        for dx,dy in valid[grid[x][y]]:
            if not(0<=x+dx<rows and 0<=y+dy<cols): continue
            work.append((dist+1, (x+dx, y+dy)))

    return visited


def part2(grid, visited):
    rows = len(grid)
    cols = len(grid[0])
    res = 0

    for r, row in enumerate(grid):
        inside = False
        for c, val in enumerate(row):
            if (r,c) not in visited:
                res += inside
            else:
                inside = inside ^ (val in "|F7")
    return res




##############
def add(ra, ca, rb, cb):
	return ra + rb, ca + cb

def find_start(grid):
	for r, row in enumerate(grid):
		for c, char in enumerate(row):
			if char == 'S':
				return r, c

def follow_pipes(grid, start_r, start_c):
	U, D, L, R = directions = ((-1, 0), (1, 0), (0, -1), (0, 1))
	possible_pipes = ('|F7', '|LJ', '-FL', '-J7')
	matches = ()

	for (dr, dc), pipes in zip(directions, possible_pipes):
		r, c = start_r + dr, start_c + dc

		if grid[r][c] in pipes:
			matches += ((dr, dc),)

	if   matches == (U, D): start_pipe = '|'
	elif matches == (L, R): start_pipe = '-'
	elif matches == (U, L): start_pipe = 'J'
	elif matches == (U, R): start_pipe = 'L'
	elif matches == (D, L): start_pipe = '7'
	else: start_pipe = 'F'

	r, c = start_r, start_c
	dr, dc = matches[0]
	seen = set([(r, c)])

	for steps in count(1):
		r, c = r + dr, c + dc
		pipe = grid[r][c]
		seen.add((r, c))

		if pipe in 'L7':
			dr, dc = dc, dr
		elif pipe in 'FJ':
			dr, dc = -dc, -dr
		elif pipe == 'S':
			break

	grid[start_r][start_c] = start_pipe
	return seen, steps

def inner_area(grid, main_loop):
	area = 0

	for r, row in enumerate(grid):
		inside = False

		for c, cell in enumerate(row):
			if (r, c) not in main_loop:
				area += inside
			else:
				inside = inside ^ (cell in '|F7')

	return area

data = read()
main_loop, loop_len = follow_pipes(data,*find_start(data))
max_pipe_dist = loop_len//2
print("loop size",len(main_loop))
print("part1", max_pipe_dist)
area = inner_area(data, main_loop)
print('area', area)



##########
data = read()
v = part1(data)
print("og size:", len(v))
print("part1: ", max(v.values()))

print("part2: ", inner_area(data, v))