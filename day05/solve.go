package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const moveindex = 1
const srcIndex = 3
const destIndex = 5

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

func part1(input []string) []string {
	stacks := buildStacks()
	for _, line := range input {
		commands := strings.Fields(line)
		numMoves, _ := strconv.Atoi(commands[moveindex])
		src, _ := strconv.Atoi(commands[srcIndex])
		dest, _ := strconv.Atoi(commands[destIndex])
		src -= 1
		dest -= 1
		resultStack, moved := cutAndReverse(numMoves, stacks[src])
		stacks[src] = resultStack
		stacks[dest] = append(stacks[dest], moved...)
	}
	res := []string{}
	for _, stack := range stacks {
		res = append(res, stack[len(stack)-1])
	}
	return res
}

func part2(input []string) []string {
	stacks := buildStacks()
	for _, line := range input {
		commands := strings.Fields(line)
		numMoves, _ := strconv.Atoi(commands[moveindex])
		src, _ := strconv.Atoi(commands[srcIndex])
		dest, _ := strconv.Atoi(commands[destIndex])
		src -= 1
		dest -= 1
		resultStack, moved := cut(numMoves, stacks[src])
		stacks[src] = resultStack
		stacks[dest] = append(stacks[dest], moved...)
	}
	res := []string{}
	for _, stack := range stacks {
		res = append(res, stack[len(stack)-1])
	}
	return res
}

func cutAndReverse(moves int, stack []string) ([]string, []string) {
	cutPoint := len(stack) - moves
	newSrc := stack[:cutPoint]
	moved := stack[cutPoint:]
	// reverse the moved so it makes it seem like we moved them "in order"
	for i, j := 0, len(moved)-1; i < j; i, j = i+1, j-1 {
		moved[i], moved[j] = moved[j], moved[i]
	}
	return newSrc, moved
}

func cut(moves int, stack []string) ([]string, []string) {
	cutPoint := len(stack) - moves
	newSrc := stack[:cutPoint]
	moved := stack[cutPoint:]
	return newSrc, moved
}

//	[M]     [B]             [N]
//
// [T]     [H]     [V] [Q]         [H]
// [Q]     [N]     [H] [W] [T]     [Q]
// [V]     [P] [F] [Q] [P] [C]     [R]
// [C]     [D] [T] [N] [N] [L] [S] [J]
// [D] [V] [W] [R] [M] [G] [R] [N] [D]
// [S] [F] [Q] [Q] [F] [F] [F] [Z] [S]
// [N] [M] [F] [D] [R] [C] [W] [T] [M]
//
//	1   2   3   4   5   6   7   8   9
func buildStacks() [][]string {
	stacks := make([][]string, 9)
	stacks[0] = []string{"N", "S", "D", "C", "V", "Q", "T"}
	stacks[1] = []string{"M", "F", "V"}
	stacks[2] = []string{"F", "Q", "W", "D", "P", "N", "H", "M"}
	stacks[3] = []string{"D", "Q", "R", "T", "F"}
	stacks[4] = []string{"R", "F", "M", "N", "Q", "H", "V", "B"}
	stacks[5] = []string{"C", "F", "G", "N", "P", "W", "Q"}
	stacks[6] = []string{"W", "F", "R", "L", "C", "T"}
	stacks[7] = []string{"T", "Z", "N", "S"}
	stacks[8] = []string{"M", "S", "D", "J", "R", "Q", "H", "N"}
	return stacks
}
