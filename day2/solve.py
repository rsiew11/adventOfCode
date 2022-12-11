
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

def read():
    with open("input.txt","r") as fd:
        return [(line[0],line[-1]) for line in fd.read().splitlines()]

def part1(moves):
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

def part2(moves):
    win_con = {
        "A": "B",
        "B": "C",
        "C": "A"
    }
    lose_con = {
        "A": "C",
        "B": "A",
        "C": "B"
    }
    score = 0
    for opp, res in moves:
        if res == "X": # lose
            score += POINT_VALS["lose"] + POINT_VALS[lose_con[opp]]
        elif res == "Y": # draw
            score += POINT_VALS["draw"] + POINT_VALS[opp]
        else: # lose
            score += POINT_VALS["win"] + POINT_VALS[win_con[opp]]
    return score


if __name__ == "__main__":
    data = read()
    print(part1(data))
    print(part2(data))

