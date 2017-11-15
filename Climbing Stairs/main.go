package main

import (
	"fmt"
)

func main() {
	n := 45
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	var stairs []int
	stairs = append(stairs, 0)
	for i := 1; i <= 3; i++ {
		stairs = append(stairs, i)
	}
	if n < 4 {
		return stairs[n]
	}

	for i := 4; i <= n; i++ {
		stairs = append(stairs, (stairs[i-1] + stairs[i-2]))
	}
	return stairs[n]

}
