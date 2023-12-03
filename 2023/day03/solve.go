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

func findNum(line []string, idx int, validNum map[string]bool) int {
	l, r := idx, idx
	fmt.Println()
	fmt.Println(line)
	fmt.Println("idx: ",idx, "num: ", line[idx])
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
	fmt.Println("parsed",line[l+1:r])
	return atoi(strings.Join(line[l+1:r], ""))
}


// if parsing/starting from a symbol, we can be double counting numbers!!! 
func part1(data []string) int {
	res := 0
	nonSymb := make(map[string]bool)
	validNum := make(map[string]bool)
	dirs := [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}, {-1,-1}, {1,-1}, {1,-1}, {1,1}}
	for _, c := range strings.Split("123456789", "") {
		nonSymb[c] = true
		validNum[c] = true
	}
	nonSymb["."] = true
	var chars [][]string
	for _,row := range data {
		chars = append(chars, strings.Split(row, ""))
	}

	for r, row := range chars {
		for c, char := range row {
			if _, ok := nonSymb[char]; !ok { // if a symbol check all dirs
				for _, d := range dirs { 
					x, y := r+d[0], c+d[1]
					if x < 0 || x >= len(row) || y >= len(chars) || y < 0 {
						continue
					}
					if _, ok := validNum[chars[x][y]]; ok {
						zz := findNum(chars[x], y, validNum)
						res += zz
					}
				}
			}
		}

	}
	return res
}

func part2(data []string) int {
	return 0
}

func main() {
	input := readFile("./input.txt")
	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))
}