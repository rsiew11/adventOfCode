from collections import deque

class Solution():
    def __init__(self):
        self.grid = self.read("input.txt")
        self.ROWS = len(self.grid)
        self.COLS = len(self.grid[0])
        self.start = self.find('S')
        self.end = self.find('E')
        self.grid[self.start[0]][self.start[1]] = ord('a')
        self.grid[self.end[0]][self.end[1]] = ord('z')

    def read(self, f):
        with open(f, "r+") as fd:
            return list(map(
                lambda row: [ord(c) for c in row],
                fd.read().splitlines())
            )
        
    def find(self, char):
        char = ord(char)
        for i, row in enumerate(self.grid):
            for j, c in enumerate(row):
                if c == char:
                    return [i,j]
    
    def canVisit(self, src, dest):
        if not(0 <= dest[0] < self.ROWS and 0 <= dest[1] < self.COLS):
            return False
        if self.grid[dest[0]][dest[1]] <= self.grid[src[0]][src[1]] + 1:
            return True
        else:
            return False
    
    def solve(self, locs):
        dirs = [(1,0), (-1,0), (0,1), (0,-1)]
        visited = set()
        work = deque()
        for loc in locs:
            work.append((loc[0], loc[1], 0))

        while work:
            x,y,d = work.popleft()
            if (x, y) in visited: 
                continue
            visited.add((x, y))

            if x == self.end[0] and y == self.end[1]:
                return d
            
            for dx,dy in dirs:
                dest = (x+dx, y+dy, d+1)
                if self.canVisit((x,y), dest):
                    work.append(dest)

    def part1(self):
        return self.solve([self.start])

    def part2(self):
        return self.solve([(r,c) for r, row in enumerate(self.grid) 
         for c in range(len(row)) if self.grid[r][c] == ord('a')])


s = Solution()
print(f"part1: {s.part1()}")
print(f"part2: {s.part2()}")