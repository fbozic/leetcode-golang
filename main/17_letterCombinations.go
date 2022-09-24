package main

//
// func main() {
// 	fmt.Println(letterCombinations("79"))
// }

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	buildRecursive(digits, "", &result)
	return result
}

func buildRecursive(digits string, base string, result *[]string) {
	if len(digits) == 0 {
		return
	}

	currentDigit := digits[0]
	for _, latter := range mapping[currentDigit] {
		newBase := base + latter

		if len(digits) == 1 {
			*result = append(*result, newBase)
		} else {
			buildRecursive(digits[1:], newBase, result)
		}
	}
}

var mapping = map[uint8][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}
