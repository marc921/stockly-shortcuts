package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
    argsWithoutProg := os.Args[1:]
	n, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		fmt.Println(err)
	}
	shortcutsStrArr := strings.Split(argsWithoutProg[1], " ")
	shortcuts := make([]int, 0)
	for _, iStr := range(shortcutsStrArr) {
		i, err := strconv.Atoi(iStr)
		if err != nil {
			fmt.Println(err)
		}
		shortcuts = append(shortcuts, i)
	}
	fmt.Println(strings.Trim(fmt.Sprint(ShortestPathTree(1, n, shortcuts)), "[]"))
}

// Assumes start-1 < shortcuts length, expects start in [1, n]
func GetTravelCost(start int, end int, shortcuts []int) int {
	if shortcuts[start-1] == end {
		return 1
	}
	return int(math.Abs(float64(start - end)))
}

// Returns travel costs from start (in [1, n]) to any node in [1, n]
func ShortestPathTree(start int, n int, shortcuts []int) []int {
	travelCosts := make([]int, n)
	travelCosts[start-1] = 0
	for i:=1; i <= n ;i++ {
		if i != start {
			travelCosts[i-1] = n	// cap
		}
	}

	queue := make([]int, 0)
	queue = append(queue, start)

	for len(queue) > 0 {
		// pop
		current := queue[0]
		queue = queue[1:]
		// update costs
		for dest := 1; dest <= n; dest++ {
			if dest != current {
				travelCost := travelCosts[current-1] + GetTravelCost(current, dest, shortcuts)
				if travelCost < travelCosts[dest-1] {
					travelCosts[dest-1] = travelCost
					queue = append(queue, dest)
				}
			}
		}
	}
	return travelCosts
}
