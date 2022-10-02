package main

import (
	"fmt"
)

// func main() {
// 	// fmt.Println(isWildcardMatch("aa", "a"))
// 	// fmt.Println(isWildcardMatch("aa", "*a"))
// 	// fmt.Println(isWildcardMatch("aa", "a*"))
// 	// fmt.Println(isWildcardMatch("cb", "?a"))
// 	// fmt.Println(isWildcardMatch("abb", "a*?*b"))
// 	// fmt.Println(isWildcardMatch("ab", "a*?*b"))
// 	// fmt.Println(isWildcardMatch("abbbbbbbb", "a*?*b"))
// 	fmt.Println(isWildcardMatch("aaaa", "*a?"))
// 	// fmt.Println(isWildcardMatch("", "******"))
// 	// fmt.Println(isWildcardMatch(
// 	// 	"abbaabbbbababaababababbabbbaaaabbbbaaabbbabaabbbbbabbbbabbabbaaabaaaabbbbbbaaabbabbbbababbbaaabbabbabb",
// 	// 	"***b**a*a*b***b*a*b*bbb**baa*bba**b**bb***b*a*aab*a**",
// 	// ))
//
// 	fmt.Println(
// 		isWildcardMatch(
// 			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
// 			"*aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa*",
// 		),
// 	)
// 	// fmt.Println(isWildcardMatch("a", "a**b***???*?*??****"))
// }

func isWildcardMatch(s string, p string) bool {
	sanitizedP := sanitizeWildcardPatter(p)

	return isWildcardMatchDP(s, sanitizedP)
}

func isWildcardMatchRecurse(s string, p string, mem map[string]bool) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	} else if len(p) == 0 {
		return false
	} else if len(s) == 0 {
		for _, c := range p {
			if c != '*' {
				return false
			}
		}
		return true
	}

	val, ok := mem[fmt.Sprintf("%s %s", s, p)]
	if ok {
		return val
	}

	currentCharS := s[0]
	currentCharP := p[0]
	if currentCharP == currentCharS || currentCharP == '?' {
		subResult := isWildcardMatchRecurse(s[1:], p[1:], mem)
		mem[fmt.Sprintf("%s %s", s, p)] = subResult
		return mem[fmt.Sprintf("%s %s", s, p)]
	}
	if currentCharP == '*' {
		subResult := isWildcardMatchRecurse(s[0:], p[1:], mem) || isWildcardMatchRecurse(s[1:], p[0:], mem)
		mem[fmt.Sprintf("%s %s", s, p)] = subResult
		return mem[fmt.Sprintf("%s %s", s, p)]
	}

	mem[fmt.Sprintf("%s %s", s, p)] = false
	return mem[fmt.Sprintf("%s %s", s, p)]
}

func sanitizeWildcardPatter(p string) string {
	sanitizedP := ""

	var prevC uint8
	for i := 0; i < len(p); i++ {
		if i == 0 {
			sanitizedP += string(p[i])
			prevC = p[i]
		} else if !(prevC == '*' && p[i] == '*') {
			sanitizedP += string(p[i])
			prevC = p[i]
		}
	}

	return sanitizedP
}

func isWildcardMatchDP(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	sizeS := len(s) + 1
	sizeP := len(p) + 1

	mem := make([][]bool, sizeS)
	for i := 0; i < sizeS; i++ {
		mem[i] = make([]bool, sizeP)
	}

	mem[0][0] = true
	for j := 1; j < sizeP; j++ {
		if p[j-1] == '*' {
			mem[0][j] = mem[0][j-1]
		}
	}

	for i := 1; i < sizeS; i++ {
		for j := 1; j < sizeP; j++ {
			currentCharS := s[i-1]
			currentCharP := p[j-1]
			if currentCharP == currentCharS || currentCharP == '?' {
				mem[i][j] = mem[i-1][j-1]
			} else if currentCharP == '*' {
				mem[i][j] = mem[i][j-1] || mem[i-1][j]
			} else {
				mem[i][j] = false
			}
		}
	}

	return mem[sizeS-1][sizeP-1]
}
