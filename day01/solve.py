

def main():
    calorie_counts = []
    with open("./input.txt","r") as fd:
        count = 0
        for line in fd.readlines():
            if line != "\n":
                count += int(line.rstrip())
            else:
                calorie_counts.append(count)
                count = 0
    return max(calorie_counts)

if __name__ == "__main__":
    print(main())