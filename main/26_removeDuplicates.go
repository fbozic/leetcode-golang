package main

import "fmt"

//
// func main() {
// 	// removeDuplicates([]int{1, 1, 2})
// 	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
// 	fmt.Println(removeDuplicates([]int{1, 1}))
// }

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	deadLetters := 0
	toCompare := nums[0]
	for i := 1; i < len(nums)-deadLetters; {
		current := nums[i]
		if current != toCompare {
			toCompare = current
			i++
		} else {
			// shift to the left
			var garbage []int
			if deadLetters == 0 {
				garbage = make([]int, 0)
			} else {
				garbage = nums[len(nums)-deadLetters:]
			}
			leftPart := nums[:i]
			rightPart := nums[i+1 : len(nums)-deadLetters]
			nums = append(leftPart, rightPart...)
			nums = append(nums, garbage...)
			nums = append(nums, current)
			deadLetters++
		}
	}

	fmt.Println(nums)
	return len(nums) - deadLetters
}
