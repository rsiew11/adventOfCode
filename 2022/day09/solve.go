package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

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

func (p *Point) moveTail(head Point) {
	dist := max(abs(p.x, head.x), abs(p.y, head.y))
	if dist <= 1 {
		return
	}

	xdiff := head.x - p.x
	if xdiff > 0 {
		p.x++
	} else if xdiff < 0 {
		p.x--
	}

	ydiff := head.y - p.y
	if ydiff > 0 {
		p.y++
	} else if ydiff < 0 {
		p.y--
	}

}

func part1(input []string) int {
	head := Point{0, 0}
	tail := Point{0, 0}
	visited := make(map[Point]bool)
	visited[tail] = true // initial point is visited

	for _, line := range input {
		commands := strings.Fields(line)
		dir := commands[0]
		amt, _ := strconv.Atoi(commands[1])
		for i := 0; i < amt; i++ {
			switch dir {
			case "L":
				head.x -= 1
				break
			case "R":
				head.x += 1
				break
			case "U":
				head.y -= 1
				break
			case "D":
				head.y += 1
				break
			}
			tail.moveTail(head)
			visited[tail] = true
		}
	}
	return len(visited)
}

func part2(input []string) int {
	var points []Point
	for i := 0; i < 10; i++ {
		points = append(points, Point{0, 0})
	}
	head := 0
	tail := 9
	visited := make(map[Point]bool)
	visited[points[tail]] = true // initial point is visited

	for _, line := range input {
		commands := strings.Fields(line)
		dir := commands[0]
		amt, _ := strconv.Atoi(commands[1])
		for i := 0; i < amt; i++ {
			// move head
			switch dir {
			case "L":
				points[head].x -= 1
				break
			case "R":
				points[head].x += 1
				break
			case "U":
				points[head].y -= 1
				break
			case "D":
				points[head].y += 1
				break
			}
			// move body & tail
			for j := 1; j < 10; j++ {
				points[j].moveTail(points[j-1]) // "head" is point closer to head
			}
			visited[points[tail]] = true
		}
	}
	return len(visited)
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}
