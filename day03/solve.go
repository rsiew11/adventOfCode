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

func part1(input []string) int {
    sum := 0
    for _,line := range input {
        mid := len(line)/2
        left := []rune(line[:mid])
        right := []rune(line[mid:])
        common := findCommonFrom2(left, right)
        sum += itemToPriority(common)
    }
    return sum
}

func part2(input []string) int {
    sum := 0
    for i := 0; i<len(input); i += 3 {
        x := []rune(input[i])
        y := []rune(input[i+1])
        z := []rune(input[i+2])
        common := findCommonFrom3(x,y,z)
        sum += itemToPriority(common)
    }
    return sum
}

func findCommonFrom2(l []rune, r []rune) rune {
    seen := make(map[rune]bool)
    for _, char := range l {
        seen[char] = true
    }
    for _, char := range r {
        if _, ok := seen[char]; ok {
            return char
        }
    }
    fmt.Println("no match")
    return '0'
}

func findCommonFrom3(x []rune, y []rune, z []rune) rune {
    seenX := make(map[rune]bool)
    seenY := make(map[rune]bool)
    for _, char := range x {
        seenX[char] = true
    }
    for _, char := range y {
        seenY[char] = true
    }
    for _, char := range z {
        _, okX := seenX[char]
        _, okY := seenY[char]
        if okX && okY {
            return char
        }
    }
    fmt.Println("no match")
    return '0'
}



func itemToPriority(item rune) int {
    switch {
    case 'a' <= item && item <= 'z': // a -> z = 1 -> 26
        return 1 + int(item - 'a')
    case 'A' <= item && item <= 'Z': // A -> Z = 27 -> 52
        return 27 + int(item - 'A')
    default:
        fmt.Println("invalid rune")
        return 0
    }
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


