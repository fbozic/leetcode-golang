package main

// func main() {
// 	// removeDuplicates([]int{1, 1, 2})
// 	// fmt.Println(strStr("", ""))
// 	// fmt.Println(strStr("sadbutsad", "sadbutsad"))
// 	fmt.Println(strStr("sadbutsad", "sad"))
// }

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	smallestIndex := -1

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			reachedEnd := true
			for j := 1; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					reachedEnd = false
					break
				}
			}
			if reachedEnd {
				if smallestIndex == -1 || i < smallestIndex {
					smallestIndex = i
				}
			}
		}
	}
	return smallestIndex
}
