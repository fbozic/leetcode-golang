package main

// 0,1,2,4,5,6,7 -- sorted
// 4,5,6,7,0,1,2 -> pivot 3
// 1,2,4,5,6,7,0 -> pivot 5
// 6,7,0,1,2,4,5 -> pivot 1
//
// func main() {
// 	// fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
// 	// fmt.Println(searchRange([]int{0, 0, 0, 8, 8, 10}, 0))
// 	fmt.Println(searchRange([]int{1, 2, 3}, 2))
// 	// fmt.Println(searchRange([]int{1}, 1))
// 	// fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
// 	// fmt.Println(searchRange([]int{}, 0))
// 	// fmt.Println(searchRange([]int{0}, 0))
// }

// 5,7,7,8,8,10 - 8 ->
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	leftMost := findLeftMost(nums, target, 0, len(nums)-1)
	rightMost := findRightMost(nums, target, 0, len(nums)-1)
	return []int{leftMost, rightMost}
}

func findLeftMost(nums []int, target, l, r int) int {
	if l > r {
		return -1
	}

	mid := (r + l) / 2
	if nums[mid] == target {
		if mid > 0 && nums[mid-1] == target {
			return findLeftMost(nums, target, l, mid-1)
		} else {
			return mid
		}
	} else if target < nums[mid] {
		return findLeftMost(nums, target, l, mid-1)
	} else {
		return findLeftMost(nums, target, mid+1, r)
	}
}

func findRightMost(nums []int, target, l, r int) int {
	if l > r {
		return -1
	}

	mid := (r + l) / 2
	if nums[mid] == target {
		if mid < len(nums)-1 && nums[mid+1] == target {
			return findRightMost(nums, target, mid+1, r)
		} else {
			return mid
		}
	} else if target < nums[mid] {
		return findRightMost(nums, target, l, mid-1)
	} else {
		return findRightMost(nums, target, mid+1, r)
	}
}
