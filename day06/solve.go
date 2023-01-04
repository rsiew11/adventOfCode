package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

func part1(input []string) int {
    // input is one item array since there is only one line
    markers := []rune(input[0])
    for i := 0; i < len(markers)-4; i++ {
        seen := make(map[rune]bool)
        seen[markers[i]] = true
        if _, ok := seen[markers[i+1]]; ok {
            continue
        }
        seen[markers[i+1]] = true
        if _, ok := seen[markers[i+2]]; ok {
            continue
        }
        seen[markers[i+2]] = true
        if _, ok := seen[markers[i+3]]; ok {
            continue
        }
        //seen[markers[i+3]] = true
        return i + 3 + 1 // +1 since we want the char right after the sequence of 4
    }
    return -1
}

func part2(input []string) int {
    markers := []rune(input[0])
    for i := 0; i < len(markers)-14; i++ {
        seen := make(map[rune]bool)
        seen[markers[i]] = true // adding first in sequence into the map
        for j := 1; j<14; j++ {
            if _, ok := seen[markers[i+j]]; ok {
                break
            }
            seen[markers[i+j]] = true
            if j == 13 {
                return i+j+1
            }
        }
    }
    return -1
}

