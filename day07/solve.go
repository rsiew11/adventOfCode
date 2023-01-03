package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const maxSize = 100000

func main() {
	data := readFile("./input.txt")
	fmt.Println(part1(data))
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

func part1(input []string) int {
	listings := dirListings(input)
	lineCnt := len(input)
	dirs := make(map[int]int)
	for _, lineNum := range listings {
		dirs[lineNum] = 0
		line := []rune(input[lineNum])
		// keep going until next human command
		for i := 0; line[0] != '$' && lineNum+i < lineCnt-1; i++ {
			fmt.Println(lineNum, i)
			fields := strings.Fields(input[lineNum+i])
			size, err := strconv.Atoi(fields[0])
			line = []rune(input[lineNum+i])
			if err != nil { // this is a dir listiting or human command
				continue
			} else {
				dirs[lineNum] += size
			}
		}
	}
	sum := 0
	for _, v := range dirs {
		if v <= maxSize {
			sum += v
		}
	}
	return sum
}
func dirListings(input []string) []int {
	var listings []int
	for i, line := range input {
		if line[:4] == "$ ls" {
			listings = append(listings, i+1) // want to start reading from line after ls
		}
	}
	return listings
}
