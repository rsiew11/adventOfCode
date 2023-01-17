package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const cycleChecks int = 6
const spriteSize int = 3

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

func contains(points *[cycleChecks]int, cycle int) bool {
	for _, point := range points {
		if cycle == point {
			return true
		}
	}
	return false
}

func part1(input []string) int {
	stopPoints := [cycleChecks]int{20, 60, 100, 140, 180, 220}
	cycle := 1
	x := 1
	sum := 0

	for _, line := range input {
		commands := strings.Fields(line)
		command := commands[0]
		switch command {
		case "addx":
			amt, _ := strconv.Atoi(commands[1])
			for i := 0; i < 2; i++ {
				if contains(&stopPoints, cycle) {
					sum += cycle * x
				}
				cycle += 1
			}
			x += amt
			break
		case "noop":
			if contains(&stopPoints, cycle) {
				sum += cycle * x
			}
			cycle += 1
			break
		}
	}
	return sum
}

func updateSprite(points *[spriteSize]int, moveAmt int) {
	points[1] += moveAmt
	points[0] = points[1] - 1
	points[2] = points[1] + 1
}

func draw(monitor *[40 * 6]string) {
	for i := 0; i < len(monitor); i += 40 {
		fmt.Println(monitor[i : i+40])
	}
}

func spriteAtLocation(sprite *[spriteSize]int, cycle int) bool {
	for _, loc := range sprite {
		if loc == cycle {
			return true
		}
	}
	return false
}

func part2(input []string) int {
	monitor := [40 * 6]string{}
	cycle := 0
	sprite := [spriteSize]int{0, 1, 2}

	for _, line := range input {
		commands := strings.Fields(line)
		command := commands[0]

		switch command {

		case "addx":
			amt, _ := strconv.Atoi(commands[1])
			for i := 0; i < 2; i++ {
				if spriteAtLocation(&sprite, cycle%40) {
					monitor[cycle] = "#"
				} else {
					monitor[cycle] = "."
				}
				cycle += 1
			}
			updateSprite(&sprite, amt)
			break
		case "noop":
			if spriteAtLocation(&sprite, cycle%40) {
				monitor[cycle] = "#"
			} else {
				monitor[cycle] = "."
			}
			cycle += 1
			break
		}
	}
	draw(&monitor)
	return 0
}
