package main

import "fmt"

type tokenType int

const (
	MatchCharOnce tokenType = iota
	MatchCharZeroOrMore
	MatchAnyOnce
	MatchAnyZeroOrMore
)

type Token struct {
	TokenType tokenType
	Char      string
}

// func main() {
// 	// fmt.Println(myAtoi("42"))
// 	// fmt.Println(myAtoi("     -42"))
// 	// fmt.Println(myAtoi("4193 with words"))
// 	// fmt.Println(myAtoi("9223372036854775808"))
// 	// fmt.Println(isMatch("aasbc", "aa..*.*.*b*b."))
// 	// fmt.Println(isMatch("aabbc", ".*.*b*b."))
// 	// fmt.Println(isMatch("aa", "a"))
// 	// fmt.Println(isMatch("aa", "a*"))
// 	// fmt.Println(isMatch("aab", "c*a*b"))
// 	fmt.Println(isMatch("mississippi", "mis*is*ip*."))
// }

// so wrong approach
// use recursive DP (i have no idea how still)
func isMatch(s string, p string) bool {
	tokens := buildTokenList(p)
	fmt.Println(tokens)
	return run(s, tokens)
}

func run(s string, tokens []Token) bool {
	sIndex := 0
	lenS := len(s)
	tokenIndex := 0
	lenTokens := len(tokens)
	for ; tokenIndex < lenTokens; tokenIndex++ {
		token := tokens[tokenIndex]

		fmt.Println(token)
		fmt.Println(sIndex)
		switch token.TokenType {
		case MatchCharOnce:
			if sIndex >= lenS {
				return false
			}
			if string(s[sIndex]) != token.Char {
				return false
			}
			sIndex++
		case MatchCharZeroOrMore:
			if tokenIndex+1 < lenTokens {
				var nextToken Token
				nextToken = tokens[tokenIndex+1]
				nextIsOkay := false

				for ; sIndex < lenS; sIndex++ {
					if string(s[sIndex]) == nextToken.Char {
						nextIsOkay = true
						break
					}
				}
				if !nextIsOkay {
					return false
				}
			} else {
				for sIndex < lenS {
					if string(s[sIndex]) != token.Char {
						break
					}
					sIndex++
				}
			}
		case MatchAnyOnce:
			if sIndex >= lenS-1 {
				return false
			}
			sIndex++
		case MatchAnyZeroOrMore:
			if tokenIndex+1 < lenTokens {
				var nextToken Token
				nextToken = tokens[tokenIndex+1]
				nextIsOkay := false

				for ; sIndex < lenS; sIndex++ {
					if string(s[sIndex]) == nextToken.Char {
						nextIsOkay = true
						break
					}
				}
				if !nextIsOkay {
					return false
				}
			} else {
				return true
			}
		default:
			panic("i don't know what is going on")
		}
	}

	if tokenIndex == lenTokens && sIndex == lenS {
		return true

	}
	return false
}

func buildTokenList(p string) []Token {
	index := 0
	lenP := len(p)

	var previousToken *Token
	tokens := make([]Token, 0)
	for {
		if index > lenP-1 {
			break
		}

		offset := 1
		currentChar := p[index]
		var token Token
		if index+1 < lenP {
			nextChar := p[index+1]

			if nextChar == '*' {
				offset += 1
				if currentChar == '.' {
					token = Token{
						TokenType: MatchAnyZeroOrMore,
					}
				} else {
					token = Token{
						TokenType: MatchCharZeroOrMore,
						Char:      string(currentChar),
					}
				}
			} else {
				if currentChar == '.' {
					token = Token{
						TokenType: MatchAnyOnce,
					}
				} else {
					token = Token{
						TokenType: MatchCharOnce,
						Char:      string(currentChar),
					}
				}
			}
		} else {
			if currentChar == '.' {
				token = Token{
					TokenType: MatchAnyOnce,
				}
			} else {
				token = Token{
					TokenType: MatchCharOnce,
					Char:      string(currentChar),
				}
			}
		}

		if token.TokenType == MatchAnyZeroOrMore {
			if !(previousToken != nil && previousToken.TokenType == MatchAnyZeroOrMore) {
				tokens = append(tokens, token)
				previousToken = &token
			}
		} else {
			tokens = append(tokens, token)
			previousToken = &token
		}

		index += offset
	}

	return tokens
}
