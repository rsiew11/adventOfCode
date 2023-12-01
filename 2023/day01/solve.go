package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func part1(input []string) int {
	res := 0
	for _, row := range input {
		runes := []rune(row)
		l, lVal := 0, "" 
		for i, val := range runes {
			if unicode.IsDigit(val) {
				l, lVal = i, string(val)
				break
			}
		}
		rVal := ""
		for i := len(runes)-1; i >= l; i-- {
			if unicode.IsDigit(runes[i]) {
				rVal = string(runes[i])
				break
			}
		}
		lineVal, err := strconv.Atoi(lVal+rVal)
		if err != nil {
			return 0
		}
		res += lineVal
	}
	return res
}

func part2(input []string) int {
	validStrs := [9]string{"one", "two", "three", "four", "five", 
						   "six", "seven", "eight", "nine"}
	res := 0

	for _, row := range input {
		runes := []rune(row)
		l, lVal := 0, ""
		for i, val := range runes {
			if unicode.IsDigit(val) {
				l, lVal = i, string(val)
				break
			}
		}
		for i, val := range validStrs {
			loc := strings.Index(row, val)
			if loc != -1 && loc < l {
				l = loc
				lVal = strconv.Itoa(i+1)
			}
		}

		r, rVal := 0, ""
		for i := len(runes)-1; i >= l; i-- {
			if unicode.IsDigit(runes[i]) {
				r, rVal = i, string(runes[i])
				break
			}
		}
		for i, val := range validStrs {
			loc := strings.LastIndex(row, val)
			if loc != -1 && loc > r {
				r = loc
				rVal = strconv.Itoa(i+1)
			}
		}
		lineVal, err := strconv.Atoi(lVal+rVal)
		if err != nil {
			fmt.Println("ERROR")
			return 0
		}
		res += lineVal
	}
	return res
}

func main() {
	data := readFile("./input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}