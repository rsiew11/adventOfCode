package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	hand   []string
	bid    int
	kind   int            // 0,1,2,3,4,5,6 --> high-card, pair, 2-pair, 3, FH, 4, 5
	cntMap map[string]int // for part 2
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

func atoi(a string) int {
	num, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	} else {
		return num
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func assignKind(c map[string]int) int {
	counts := make([]int, 0)
	for _, v := range c {
		counts = append(counts, v)
	}
	sort.Ints(counts)
	switch counts[len(counts)-1] {
	case 5:
		return 6 // 5 of a kind
	case 4:
		return 5 // 4 of a kind
	case 3:
		if counts[len(counts)-2] == 2 {
			return 4 // full house
		} else {
			return 3 // 3 of a kind
		}
	case 2:
		if counts[len(counts)-2] == 2 {
			return 2 // 2-pair
		} else {
			return 1 // pair
		}
	case 1:
		return 0 // high-card
	}
	return 0
}

func formatData(data []string) []hand {
	hands := make([]hand, len(data))
	for i, line := range data {
		h := strings.Split(strings.Split(line, " ")[0], "")
		bid := atoi(strings.Split(line, " ")[1])
		hands[i] = hand{h, bid, 0, nil}
	}
	// assign kind
	for i, h := range hands {
		counts := make(map[string]int)
		for _, c := range h.hand {
			if _, ok := counts[c]; ok {
				counts[c]++
			} else {
				counts[c] = 1
			}
		}
		hands[i].cntMap = counts
		hands[i].kind = assignKind(counts)
	}
	return hands
}

func part1(hands []hand) int {
	values := map[string]int{
		"2": 0, "3": 1, "4": 2, "5": 3, "6": 4, "7": 5,
		"8": 6, "9": 7, "T": 8, "J": 9, "Q": 10, "K": 11, "A": 12}
	sort.Slice(hands, func(a, b int) bool {
		if hands[a].kind < hands[b].kind {
			return true
		} else if hands[a].kind > hands[b].kind {
			return false
		}
		for i, acard := range hands[a].hand {
			bcard := hands[b].hand[i]
			if values[acard] < values[bcard] {
				return true
			} else if values[acard] > values[bcard] {
				return false
			}
		}
		return false
	})
	res := 0
	for i, hand := range hands {
		res += hand.bid * (i + 1)
	}
	return res
}

func findKey(counts map[string]int, val int) string { // find key that is NOT a J
	best, cnt := "", 0
	for k, v := range counts {
		if val == v && k != "J" {
			return k
		}
		if v > cnt && k != "J" {
			cnt, best = v, k
		}
	}
	return best
}

// recursive calls if there exists more than 1 Joker (in the optimal hand) ?
func assignOpt(h hand) int {
	optimal := make([]string, len(h.hand))
	copy(optimal, h.hand)
	for slices.Contains(optimal, "J") {
		j := slices.Index(optimal, "J")
		switch h.kind {
		case 6: // 5 of a kind
			return h.kind // cant do anything here (5 * J)
		case 5: // 4 of a kind
			k := findKey(h.cntMap, 4)
			optimal[j] = k
			h.cntMap[k] += 1
			h.cntMap["J"] -= 1
			h.kind = assignKind(h.cntMap)
		case 4: // full house
			k := findKey(h.cntMap, 3) // -> 4 of a kind
			optimal[j] = k
			h.cntMap[k] += 1
			h.cntMap["J"] -= 1
			h.kind = assignKind(h.cntMap)
		case 3: // 3 of a kind
			k := findKey(h.cntMap, 3) // -> 4 of a kind
			optimal[j] = k
			h.cntMap[k] += 1
			h.cntMap["J"] -= 1
			h.kind = assignKind(h.cntMap)
		case 2: // 2 pair
			k := findKey(h.cntMap, 2) // -> to full house
			optimal[j] = k
			h.cntMap[k] += 1
			h.cntMap["J"] -= 1
			h.kind = assignKind(h.cntMap)
		case 1: // pair
			k := findKey(h.cntMap, 2) // --> to 3 of a kind
			optimal[j] = k
			h.cntMap[k] += 1
			h.cntMap["J"] -= 1
			h.kind = assignKind(h.cntMap)
		case 0: // high card
			k := findKey(h.cntMap, 1) // --> to pair
			optimal[j] = k
			h.cntMap[k] += 1
			h.cntMap["J"] -= 1
			h.kind = assignKind(h.cntMap)
		}
	}
	return h.kind
}

func part2(hands []hand) int {
	values := map[string]int{
		"J": 0, "2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6,
		"8": 7, "9": 8, "T": 9, "Q": 10, "K": 11, "A": 12}

	// find best hand if jokers exist in the hand and update kind
	for i, hand := range hands {
		if _, ok := hand.cntMap["J"]; ok {
			hands[i].kind = assignOpt(hand)
		}
	}
	sort.Slice(hands, func(a, b int) bool {
		if hands[a].kind < hands[b].kind {
			return true
		} else if hands[a].kind > hands[b].kind {
			return false
		}
		for i, acard := range hands[a].hand {
			bcard := hands[b].hand[i]
			if values[acard] < values[bcard] {
				return true
			} else if values[acard] > values[bcard] {
				return false
			}
		}
		return false
	})
	res := 0
	for i, hand := range hands {
		res += hand.bid * (i + 1)
	}
	return res
}

func main() {
	data := readFile("./input.txt")
	m := formatData(data)
	fmt.Println("part1: ", part1(m))
	fmt.Println("part2: ", part2(m))
}
