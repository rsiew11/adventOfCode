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

func part2(moves []string, m map[string]loc) int {
	return 0
}

func main() {
	data := readFile("./input.txt")
	moves, m := formatData(data)
	fmt.Println("part1: ", part1(moves, m))
	fmt.Println("part2: ", part2(moves, m))
}
