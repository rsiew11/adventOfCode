package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time int
	dist int
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func atoi(a string) int {
	num, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	} else {
		return num
	}
}

func formatData(d []string) []race {
	times := strings.Fields(d[0])[1:]
	dists := strings.Fields(d[1])[1:]
	res := make([]race, len(times))
	for i, time := range times {
		t, d := atoi(time), atoi(dists[i])
		res[i] = race{t, d}
	}
	return res
}

func part1(races []race) int {
	wins := make([]int, len(races))
	for raceId, race := range races {
		speed := 1
		traveled := 0
		for hold := 1; hold < race.time; hold++ {
			traveled = speed * (race.time - hold)
			if traveled > race.dist {
				wins[raceId]++
			}
			speed++
		}
	}

	res := 1
	for _, win := range wins {
		res *= win
	}
	return res
}

func part2(races []race) int {
	time := ""
	dist := ""
	for _, race := range races {
		time += strconv.Itoa(race.time)
		dist += strconv.Itoa(race.dist)
	}
	t := atoi(time)
	d := atoi(dist)
	speed := 1
	traveled := 0
	wins := 0
	for hold := 1; hold < t; hold++ {
		traveled = speed * (t - hold)
		if traveled > d {
			wins++
		}
		speed++
	}
	return wins
}

func main() {
	data := readFile("./input.txt")
	m := formatData(data)
	fmt.Println("part1: ", part1(m))
	fmt.Println("part2: ", part2(m))
}
