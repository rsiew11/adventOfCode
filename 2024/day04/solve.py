
def read():
    with open("./input.txt") as fd:
        return [list(line.rstrip()) for line in fd.readlines()]
    


def p1():
    grid = read()
    rows, cols = len(grid), len(grid[0])
    count = 0

    def check(r,c,dr,dc):
        word = []
        for i in range(4):  
            nr, nc = r + (i * dr), c + (i * dc)
            if 0 <= nr < rows and 0 <= nc < cols:
                word.append(grid[nr][nc])
            else:
                return False 
        return "".join(word) == "XMAS"

    for x, row in enumerate(grid):
        for y, val in enumerate(row):
            if val != "X": continue
            for dx, dy in [(0, 1), (0, -1), (1, 0), (-1, 0), (1, 1), (1, -1), (-1, 1), (-1, -1)]:
                if check(x, y, dx, dy): count += 1
    return count


def p2():
    grid = read()
    rows, cols = len(grid), len(grid[0])
    count = 0
    def check(r,c):
        if not(0 <= r-1 and r+1 < rows): return False
        if not(0 <= c-1 and c+1 < cols): return False
        return (
            sorted([grid[r-1][c-1], grid[r+1][c+1]]) == ["M", "S"] and 
            sorted([grid[r-1][c+1], grid[r+1][c-1]]) == ["M", "S"]
        )
    
    for x, row in enumerate(grid):
        for y, val in enumerate(row):
            if val != "A": continue
            if check(x,y): count += 1
    return count


if __name__ == "__main__":
    print(p1())
    print(p2())