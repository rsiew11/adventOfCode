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
	fmt.Println(part2(data))
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

func sizeMap(input []string) map[string]int {
	listings := make(map[string]int)
	var path []string
	for _, line := range input {
		fields := strings.Fields(line)
		if fields[1] == "cd" {
			if fields[2] == ".." {
				path = path[:len(path)-1] // pop last part
			} else {
				path = append(path, fields[2])
			}
		} else if fields[1] == "ls" || fields[0] == "dir" {
			continue
		} else {
			size, _ := strconv.Atoi(fields[0])
			for i := 1; i < len(path)+1; i++ {
				pathName := strings.Join(path[:i], "/")
				if _, ok := listings[pathName]; ok {
					listings[pathName] += size
				} else {
					listings[pathName] = size
				}
			}
		}
	}
	return listings
}

func part1(input []string) int {
	sum := 0
	dirs := sizeMap(input)
	for _, v := range dirs {
		if v <= maxSize {
			sum += v
		}
	}
	return sum
}

func part2(input []string) int {
	dirs := sizeMap(input)
	const totalSpace = 70000000
	const requiredSpace = 30000000
	availSpace := totalSpace - dirs["/"]
	toDel := requiredSpace - availSpace

	min := dirs["/"]
	for _, v := range dirs {
		if v >= toDel && v < min {
			min = v
		}
	}
	return min
}
