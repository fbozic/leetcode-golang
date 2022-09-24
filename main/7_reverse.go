package main

import (
	"fmt"
	"math"
)

// func main() {
// 	fmt.Println(reverse(320))
// 	fmt.Println(reverse(-123))
// }

func reverse(x int) int {
	fmt.Println(math.MaxInt32)
	x32 := int32(x)
	var reverseX int32
	reverseX = 0
	for x32 != 0 {
		pop := x32 % 10
		x32 /= 10

		if reverseX > math.MaxInt32/10 || (reverseX == math.MaxInt32/10 && pop > 8) {
			return 0
		}
		if reverseX < math.MinInt32/10 || (reverseX == math.MinInt32/10 && pop < 7) {
			return 0
		}

		reverseX = reverseX*10 + pop
	}

	return int(reverseX)
}

// 2147483647
// 1534236469
