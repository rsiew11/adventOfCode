package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type loc struct {
	left  string
	right string
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

func formatData(lines []string) ([]string, map[string]loc) {
	moves := strings.Split(lines[0], "")
	m := make(map[string]loc)
	for _, line := range lines[2:] {
		fields := strings.Split(line, " = ")
		start, dests := fields[0], strings.Split(fields[1], ", ")
		m[start] = loc{dests[0][1:], dests[1][:len(dests[1])-1]}
	}
	return moves, m
}

func part1(moves []string, m map[string]loc) int {
	cur, count := "AAA", 0
	for cur != "ZZZ" {
		for _, move := range moves {
			if move == "R" {
				cur = m[cur].right
			} else {
				cur = m[cur].left
			}
			count++
		}
	}
	return count
}

func findLocations(m map[string]loc, char byte) []string {
	locations := make([]string, 0)
	for k, _ := range m {
		if k[len(k)-1] == char {
			locations = append(locations, k)
		}
	}
	return locations
}

func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(vals []int) int {
	res := 1
	for _, v := range vals {
		res = (v * res) / gcd(v, res)
	}
	return res
}

func part2(moves []string, m map[string]loc) int {
	curLocs := findLocations(m, 'A')
	endLocs := make(map[string]bool)
	for _, l := range findLocations(m, 'Z') {
		endLocs[l] = true
	}
	counts := make([]int, len(curLocs))

	for i, start := range curLocs {
		cur := start
		for cur[len(cur)-1] != 'Z' {
			for _, move := range moves {
				if move == "R" {
					cur = m[cur].right
				} else {
					cur = m[cur].left
				}
				counts[i]++
			}
		}
	}
	return lcm(counts)
}

func main() {
	data := readFile("./input.txt")
	moves, m := formatData(data)
	fmt.Println("part1: ", part1(moves, m))
	fmt.Println("part2: ", part2(moves, m))
}
