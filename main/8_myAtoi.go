package main

import (
	"math"
)

// func main() {
// 	// fmt.Println(myAtoi("42"))
// 	// fmt.Println(myAtoi("     -42"))
// 	// fmt.Println(myAtoi("4193 with words"))
// 	// fmt.Println(myAtoi("9223372036854775808"))
// 	fmt.Println(myAtoi("+"))
// }

func myAtoi(s string) int {
	index := readInWhitespace(s)
	if index == -1 {
		return 0
	}

	prefix := 1
	switch s[index] {
	case '-':
		prefix = -1
		index++
	case '+':
		index++
	}

	number := 0
	if index < len(s) && s[index] >= '0' && s[index] <= '9' {
		number = (number*10 + int(s[index]-'0')) * prefix
		index++
	} else {
		return 0
	}

	for ; index < len(s); index++ {
		currentChar := s[index]
		if currentChar >= '0' && currentChar <= '9' {
			number = number*10 + (int(currentChar-'0') * prefix)
			if number >= math.MaxInt32 {
				return math.MaxInt32
			}
			if number <= math.MinInt32 {
				return math.MinInt32
			}
		} else {
			break
		}
	}

	return number
}

// returns index of first non-whitespace char
func readInWhitespace(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			return i
		}
	}
	return -1
}
