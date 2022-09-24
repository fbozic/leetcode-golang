package main

import (
	"math"
)

// func main() {
// 	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
// }

func maxArea(height []int) int {
	max := 0.0
	leftIndex := 0
	rightIndex := len(height) - 1
	for leftIndex < rightIndex {
		area := math.Min(float64(height[leftIndex]), float64(height[rightIndex])) * float64(rightIndex-leftIndex)
		max = math.Max(max, area)

		if height[leftIndex] < height[rightIndex] {
			leftIndex++
		} else {
			rightIndex--
		}
	}

	return int(max)
}
