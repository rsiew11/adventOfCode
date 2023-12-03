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

func atoi(a string) int {
	num, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	} else {
		return num
	}
}

func part1(data []string) int {
	res := 0
	nonSymb := make(map[string]bool)
	validNum := make(map[string]bool)
	dirs := [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}, {-1,-1}, {1,-1}, {-1,1}, {1,1}}
	for _, c := range strings.Split("1234567890", "") {
		nonSymb[c] = true
		validNum[c] = true
	}
	nonSymb["."] = true
	var chars [][]string
	for _,row := range data {
		chars = append(chars, strings.Split(row, ""))
	}
	keep, num := false, ""
	for r, row := range chars {
		for c, char := range row {
			if _, ok := validNum[char]; ok { 
				num += char
				for _, d := range dirs { 
					x, y := r+d[0], c+d[1]
					if x < 0 || x >= len(row) || y >= len(chars) || y < 0 { 
						continue // if invalid
					}
					if _, ok := nonSymb[chars[x][y]]; !ok { // if there is adjacent symbol
						keep = true
					}
				}
			} else {
				if keep == true {
					res += atoi(num)
				}
				keep, num = false, ""
			}
		}
	}
	return res
}

func findNum(line []string, idx int, validNum map[string]bool) int {
	l, r := idx, idx
	// search left
	for l >= 0 {
		if _, ok := validNum[line[l]]; ok {
			l--
		} else {
			break
		}
	}
	// search right
	for r < len(line) {
		if _, ok := validNum[line[r]]; ok {
			r++
		} else {
			break
		}
	}
	return atoi(strings.Join(line[l+1:r], ""))
}


func part2(data []string) int {
	res := 0
	validNum := make(map[string]bool)
	dirs := [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}, {-1,-1}, {1,-1}, {-1,1}, {1,1}}
	for _, c := range strings.Split("1234567890", "") {
		validNum[c] = true
	}
	var chars [][]string
	for _,row := range data {
		chars = append(chars, strings.Split(row, ""))
	}
	for r, row := range chars {
		for c, char := range row {
			if char != "*" { continue }
			adjacent := make(map[int]bool)
			for _, d := range dirs { // from the * check surrounding for nums
				x, y := r+d[0], c+d[1]
				if _, ok := validNum[chars[x][y]]; ok {
					n := findNum(chars[x], y, validNum)
					// fmt.Println("found", n)
					adjacent[n] = true
				}
			}
			product := 1
			if len(adjacent) == 2 {
				for k, _ := range adjacent {
					product *= k
				}
				res += product
			}
		}
	}
	return res
}

func main() {
	input := readFile("./input.txt")
	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}