package main

import (
	"fmt"
	"math"
)

func main() {
	x := 2147395599
	fmt.Println(mySqrt(x))
}

func mySqrt(a int) int {
	switch {
	case a == 0:
		return 0
	case a < 4:
		return 1
	case a < 9:
		return 2
	}

	var n = float64(a)
	var x = float64(1)
	for {
		var nx = (x + n/x) / 2
		if math.Abs(x-nx) < 1e-5 {
			break
		}
		x = nx
	}
	return int(x)
}
