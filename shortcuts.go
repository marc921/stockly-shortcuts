package main

import (
	"fmt"
	"math"
)

/**
Input
The first line contains an integer n (1 ≤ n ≤ 200 000) — the number of Mike's city intersection.

The second line contains n integers a1, a2, ..., an (i ≤ ai ≤ n , , describing shortcuts of Mike's city,
allowing to walk from intersection i to intersection ai using only 1 unit of energy.
Please note that the shortcuts don't allow walking in opposite directions (from ai to i).

Output
In the only line print n integers m1, m2, ..., mn,
where mi denotes the least amount of total energy required to walk from intersection 1 to intersection i.
*/
func main() {
	fmt.Println(Test(3, []int{2, 2, 3}, []int{0, 1, 2}))
	fmt.Println(Test(5, []int{1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4}))
	fmt.Println(Test(7, []int{4, 4, 4, 4, 7, 7, 7}, []int{0, 1, 2, 1, 2, 3, 3}))
	fmt.Println(Test(7, []int{3, 1, 5, 6, 7, 7, 7}, []int{0, 1, 2, 1, 2, 3, 3}))
}

func Test(n int, shortcuts []int, expected []int) bool {
	actual := ShortestPathTree(1, n, shortcuts)
	for i:=0;i<n;i++ {
		if actual[i] != expected[i] {
			fmt.Println(actual, expected)
			return false
		}
	}
	return true
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
	// Push to the queue
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
