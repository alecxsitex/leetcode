package main

import "fmt"
import "strconv"

/*
1.     1
2.     11
3.     21
4.     1211
5.     111221
6. 	   312211
7.     13112221
8.     1113213211
*/

func main() {
	var gen = 5
	fmt.Println(countAndSay(gen))
}

func countAndSay(n int) string {
	var result string
	if n == 0 || n < 0 {
		return ""
	}

	result = "1"

	for ; n > 1; n-- {
		var cur = ""
		for i := 0; i < len(result); i++ {
			var count = 1
			for (i+1) < len(result) && result[i] == result[i+1] {
				count++
				i++
			}

			cur += strconv.Itoa(count) + string(result[i])

		}
		result = cur
	}
	return result
}
