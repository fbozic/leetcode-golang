package main

import (
	"math"
)

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
// 	fmt.Println(lengthOfLongestSubstring("geeksforgeeks"))
// 	fmt.Println(lengthOfLongestSubstring("pwwkew"))
// 	fmt.Println(lengthOfLongestSubstring("dvdf"))
// }

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return n
	}

	latterMap := make(map[string]int) // key is latter, value is last index at which latter was found
	leftIndex := 0                    // left-most element in window
	max := -1.0
	for rightIndex := 0; rightIndex < n; rightIndex++ {
		currentLatter := string(s[rightIndex])
		latterIndex, ok := latterMap[currentLatter]
		if ok && latterIndex >= leftIndex {
			leftIndex = latterIndex + 1
		}

		latterMap[currentLatter] = rightIndex
		max = math.Max(max, float64(rightIndex-leftIndex+1))
	}

	return int(max)
}
