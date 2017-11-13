package main

import (
	"fmt"
)

func main() {
	var roman = "DCXXI"
	fmt.Println(romanToInt(roman))
}

func romanToInt(s string) int {
	var cost []int
	var result int
	for _, num := range s {
		switch rom := string(num); rom {
		case "I":
			cost = append(cost, 1)
		case "V":
			cost = append(cost, 5)
		case "X":
			cost = append(cost, 10)
		case "L":
			cost = append(cost, 50)
		case "C":
			cost = append(cost, 100)
		case "D":
			cost = append(cost, 500)
		case "M":
			cost = append(cost, 1000)

		}
	}

	for i := 0; i < (len(cost) - 1); i++ {
		if cost[i] < cost[i+1] {
			result = result - cost[i]
		} else {
			result = result + cost[i]
		}
	}
	result = result + cost[(len(cost)-1)]
	return result
}
