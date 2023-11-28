package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	data := readFile("./input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part1(input []string) int {
	pairs := 0
	for _, line := range input {
		l0, r0, l1, r1 := findBounds(line)
		if (l0 <= l1 && r1 <= r0) || (l1 <= l0 && r0 <= r1) {
			pairs++
		}
	}
	return pairs
}

func part2(input []string) int {
	overlap := 0
	for _, line := range input {
		l0, r0, l1, r1 := findBounds(line)
		switch {
		case l1 <= l0 && l0 <= r1:
			overlap++
		case l1 <= r0 && r0 <= r1:
			overlap++
		case l0 <= l1 && l1 <= r0:
			overlap++
		case l0 <= r1 && r1 <= r0:
			overlap++
		default:
			continue
		}
	}
	return overlap
}

func findBounds(line string) (int, int, int, int) {
	elf0, elf1, _ := strings.Cut(line, ",")
	l0Str, r0Str, _ := strings.Cut(elf0, "-")
	l1Str, r1Str, _ := strings.Cut(elf1, "-")
	l0, _ := strconv.Atoi(l0Str)
	r0, _ := strconv.Atoi(r0Str)
	l1, _ := strconv.Atoi(l1Str)
	r1, _ := strconv.Atoi(r1Str)
	return l0, r0, l1, r1
}
