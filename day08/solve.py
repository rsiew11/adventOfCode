from math import prod


def main():
    with open("./input.txt", "r+") as fd:
        lines = list(
            map(lambda row: [int(tree) for tree in row], 
            fd.read().splitlines())
        )
    print(part1(lines))
    print(part2(lines))

def part1(lines):
    ROWS = len(lines)
    COLS = len(lines[0])
    visibleCount = 0
    seen = set()

    #LEFT
    for r,row in enumerate(lines):
        max_height = -1
        for c,height in enumerate(row):
            if max_height == 9: break
            if height > max_height: # if visible
                visibleCount += 1
                seen.add((r,c))
            max_height = max(max_height, height)
    #RIGHT
    for r,row in enumerate(lines):
        max_height = -1
        for c in range(COLS-1, -1, -1): # reverse order
            height = lines[r][c]
            if max_height == 9: break
            if height > max_height and (r,c) not in seen:
                visibleCount += 1
                seen.add((r,c))
            max_height = max(max_height, height) 

    #TOP
    for c in range(COLS):
        max_height = -1
        for r in range(ROWS):
            height = lines[r][c]
            if max_height == 9: break
            if height > max_height and (r,c) not in seen:
                visibleCount += 1
                seen.add((r,c))
            max_height = max(max_height, height)

    #BOT
    for c in range(COLS):
        max_height = -1
        for r in range(ROWS-1, -1, -1):
            height = lines[r][c]
            if max_height == 9: break
            if height > max_height and (r,c) not in seen:
                visibleCount += 1
            max_height = max(max_height, height)
    return visibleCount

def part2(lines):
    print("\npart2")
    ROWS = len(lines)
    COLS = len(lines[0])

    def scenicScore(row, col):
        scores = [0,0,0,0]
        cur_height = lines[row][col]
        #UP
        r = row-1
        while r >= 0:
            scores[0] += 1
            if lines[r][col] >= cur_height: break
            r -= 1
        #DOWN
        r = row+1
        while r < ROWS:
            scores[1] += 1
            if lines[r][col] >= cur_height: break
            r += 1
        #LEFT
        c = col-1
        while c >= 0:
            scores[2] += 1
            if lines[row][c] >= cur_height: break
            c -= 1
        #RIGHT
        c = col+1
        while c < COLS:
            scores[3] += 1
            if lines[row][c] >= cur_height: break
            c += 1
        return prod(scores)

    return max(scenicScore(r,c) for r in range(1,ROWS-1) for c in range(1,COLS-1))


if __name__ == "__main__":
    main()
    