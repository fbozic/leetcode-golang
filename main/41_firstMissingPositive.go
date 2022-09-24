package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(firstMissingPositive([]int{0, 10, 2, -10, -20})) // 1
	// fmt.Println(firstMissingPositive([]int{1, 2, 3, 4, 5}))      // 6
	// fmt.Println(firstMissingPositive([]int{1, 2, 4, 5}))         // 3
	// fmt.Println(firstMissingPositive([]int{0, -1}))              // 1
	// fmt.Println(firstMissingPositive([]int{1, 2, 0}))            // 3
	// fmt.Println(firstMissingPositive([]int{1, 2, -1, 4}))        // 3
	// fmt.Println(firstMissingPositive([]int{-1, -2, -60, 40, 4})) // 1
	// fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{1, 1}))
}

func firstMissingPositive(nums []int) int {
	nonPositiveCount := 0
	for i := 0; i < len(nums)-nonPositiveCount; {
		if nums[i] <= 0 {
			// push to back
			nonPositiveCurrent := nums[i]
			nums[i] = nums[len(nums)-nonPositiveCount-1]
			nums[len(nums)-nonPositiveCount-1] = nonPositiveCurrent
			nonPositiveCount++
		} else {
			i++
		}
	}

	if nonPositiveCount > 0 {
		nums = nums[:len(nums)-nonPositiveCount]
	}

	positiveCount := float64(len(nums))
	for i := 0; i < len(nums); i++ {
		currentAbsVal := math.Abs(float64(nums[i]))
		if currentAbsVal <= positiveCount {
			currentMappedIndex := int(currentAbsVal) - 1
			nums[currentMappedIndex] = -int(math.Abs(float64(nums[currentMappedIndex])))
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
