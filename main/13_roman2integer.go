package main

import "fmt"

var romanToIntMapping = map[string]int{
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

// func main() {
// 	fmt.Println(romanToInt("III"))
// 	fmt.Println(romanToInt("LVIII"))
// 	fmt.Println(romanToInt("MCMXCIV"))
// }

func romanToInt(s string) int {
	result := 0

	for i := 0; i < len(s); {
		nextRomanNumber := getNextRomanNumber(i, s)
		result += romanToIntMapping[nextRomanNumber]
		i += len(nextRomanNumber)
	}

	return result
}

func getNextRomanNumber(i int, s string) string {
	romanNumber := ""
	if i+1 <= len(s)-1 {
		potentialRomanNumber := fmt.Sprintf("%c%c", s[i], s[i+1])
		if _, ok := romanToIntMapping[potentialRomanNumber]; ok {
			romanNumber = potentialRomanNumber
		} else {
			romanNumber = fmt.Sprintf("%c", s[i])
		}
	} else {
		romanNumber = fmt.Sprintf("%c", s[i])
	}
	return romanNumber
}
