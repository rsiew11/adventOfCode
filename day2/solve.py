
POINT_VALS = {
    "win": 6,
    "draw": 3,
    "lose": 0,
    "A": 1,
    "B": 2,
    "C": 3,
    "X": 1,
    "Y": 2,
    "Z": 3
}

WIN_CON = {
    "A": ""
}


def main():
    with open("input.txt","r") as fd:
                # theirs, mine
        moves = [(line[0],line[-1]) for line in fd.read().splitlines()]



if __name__ == "__main__":
    main()