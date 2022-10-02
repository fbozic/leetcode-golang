package main

import (
	"math"
)

// func main() {
// 	// fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
// 	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
// }

func trap(height []int) int {
	size := len(height)
	highestLefts := make([]int, size)
	highestRights := make([]int, size)

	highest := height[0]
	highestLefts[0] = height[0]
	for i := 1; i < size; i++ {
		if height[i] > highest {
			highest = height[i]
		}
		highestLefts[i] = highest
	}

	highest = height[size-1]
	highestRights[size-1] = height[size-1]
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > highest {
			highest = height[i]
		}
		highestRights[i] = highest
	}

	volume := 0
	for i := 0; i < len(height); i++ {
		highestLeft := highestLefts[i]
		highestRight := highestRights[i]

		volumeToAdd := int(math.Min(float64(highestRight), float64(highestLeft))) - height[i]
		volume += volumeToAdd
	}
	return volume
}
