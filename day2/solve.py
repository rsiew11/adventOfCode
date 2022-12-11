
POINT_VALS = {
    "win": 6,
    "draw": 3,
    "lose": 0,
    "A": 1, # rock
    "B": 2, # paper
    "C": 3, # scissor
    "X": 1, # rock
    "Y": 2, # paper
    "Z": 3  # scissor
}


def part1():
    win_con = {
        "X": "C",
        "Y": "A",
        "Z": "B"
    }
    draw_con = {
        "X": "A",
        "Y": "B",
        "Z": "C"
    }
    with open("input.txt","r") as fd:
        moves = [(line[0],line[-1]) for line in fd.read().splitlines()]

    score = 0
    for opp, me in moves:
        score += POINT_VALS[me]
        if win_con[me] == opp:
            score += POINT_VALS["win"]
        elif draw_con[me] == opp:
            score += POINT_VALS["draw"]
        else:
            score += POINT_VALS["lose"]
    return score

def part2():
    pass


if __name__ == "__main__":
    print(part1())
    print(part2())

