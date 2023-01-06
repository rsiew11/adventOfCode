


def main():
    with open("./input.txt", "r+") as fd:
        lines = list(
            map(lambda row: [int(tree) for tree in row], 
            fd.read().splitlines())
        )
    ROWS = len(lines)
    COLS = len(lines[0])

            # left/right + top/box - corners
    visibleCount = (ROWS * 2) + (COLS * 2) - 4 

    #LEFT
    for row in lines:
        max_height = 0
        for height in row:
            if max_height == 9:
                break
            if height > max_height: # if visible
                visibleCount += 1
            max_height = max(max_height, height)
    #RIGHT
    for row in lines:
        for height in row[::-1]

    #TOP

    #BOT

    return 0



if __name__ == "__main__":
    print(main())
    