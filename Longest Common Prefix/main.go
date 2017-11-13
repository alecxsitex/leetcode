package main

import "sort"

func main() {

}

func longestCommonPrefix(strs []string) string {
	var result []rune
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	if len(strs) > 1 {
		sort.Strings(strs)
		a := strs[0]
		b := strs[len(strs)-1]

		for i := 0; i < len(a); i++ {
			if len(b) > i && b[i] == a[i] {
				result = append(result, rune(b[i]))
			} else {
				return string(result)
			}
		}
		return string(result)
	}

	return ""
}
