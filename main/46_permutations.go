package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	permuteRecurse(0, nums, &result)
	return result
}

func permuteRecurse(pivotIndex int, nums []int, result *[][]int) {
	if pivotIndex == len(nums)-1 {
		*result = append(*result, nums)
		return
	}

	for i := pivotIndex; i < len(nums); i++ {
		rotatedNums := rotate(pivotIndex, i, nums)
		permuteRecurse(pivotIndex+1, rotatedNums, result)
	}
}

func rotate(pivotIndex int, index int, nums []int) []int {
	cpy := make([]int, len(nums))
	copy(cpy, nums)

	tmp := nums[pivotIndex]
	cpy[pivotIndex] = cpy[index]
	cpy[index] = tmp

	return cpy
}
