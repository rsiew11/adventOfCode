package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
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

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func strToInt(s string, i int) int {
	val, err := strconv.Atoi(s[1:i-1])
	if err != nil {
		panic(err)
	} else {
		return val
	}
}

func formatGames(input []string) [][3]int {
	games := make([][3]int, len(input)) // [ [r,g,b], [r,g,b] ]
	for i,row := range input {
		allRounds := strings.Split(row, ":")[1]
		for _, round := range strings.Split(allRounds, ";") {
			for _, cube := range strings.Split(round, ",") 	{
				r := strings.Index(cube, "red")
				g := strings.Index(cube, "green")
				b := strings.Index(cube, "blue")
				if r != -1 {
					games[i][0] = max(games[i][0], strToInt(cube, r))
				} else if g != -1 {
					games[i][1] = max(games[i][1], strToInt(cube, g))
				} else if b != -1 {
					games[i][2] = max(games[i][2], strToInt(cube, b))
				}
			}
		}
	}
	return games
}


func part1(games [][3]int) int {
	c := make(map[string]int)
	c["r"], c["g"], c["b"] = 12, 13, 14
	sum := 0
	for id, game := range games {
		if game[0] <= c["r"] && game[1] <= c["g"] && game[2] <= c["b"] {
			sum += id + 1
		}
	}
	return sum
}

func part2(games [][3]int) int {
	sum := 0
	for _, game := range games {
		sum += game[0] * game[1] * game[2]
	}
	return sum
}

func main() {
	data := readFile("./input.txt")
	games := formatGames(data)
	fmt.Println("part1:", part1(games))
	fmt.Println("part2:", part2(games))
}