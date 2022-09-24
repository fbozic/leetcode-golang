package main

import "fmt"

// func main() {
// 	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
// 	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
// }

func longestCommonPrefix(strs []string) string {
	prefix := ""

	string1 := strs[0]
	for i := 0; i < len(string1); i++ {
		currentChar := string1[i]

		charAtIndex := true
		for j := 1; j < len(strs); j++ {
			charAtIndex = charAtIndex && checkIfCharAtIndex(currentChar, i, strs[j])
		}
		if charAtIndex {
			prefix += fmt.Sprintf("%c", currentChar)
		} else {
			break
		}
	}

	return prefix
}

func checkIfCharAtIndex(char uint8, index int, s string) bool {
	if index > len(s)-1 {
		return false
	}
	return char == s[index]
}
