package main

import (
	"sort"
)

// func main() {
// 	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
// 	// fmt.Println(threeSum([]int{9, -9, 4, 12, 12, 0, -14, -7, 10, -1, 11, -10, -3, 2, -9, 0, 8, -9, -5, -1, 10, 5, 13, -5, -9, -12, 9, -3, 10, 10, -10, 4, 8, 1, -7, -2, -14, -6, 6, 11, 8, -6, 9, 13, 11, 7, -10, -4, 14, 0, 3, 1, -10, -7, 3, -12, -3, -11, 0, -8, -15, 5, 3, 8, 13, 11, 13, -15, -9, 4, 3, 6, 5, -11, -4, -6, 4, 1, 5, -5, -13, -7, 11, -8, 2, -1, -12, 14, 3, 3, 13, -5, -14, -7, 11, 14, -11, 9, 6, -13, -9, -13, 1, 11, -9, 12, -10, 2, -1, 3, 12, -7, 3, 0, 0, 12, 6, 3, 3, -13, 14, 1, -3}))
// }

func threeSum(nums []int) [][]int {
	cache := make(map[int]int)
	for _, num := range nums {
		_, ok := cache[num]
		if ok {
			cache[num] += 1
		} else {
			cache[num] = 1
		}
	}

	result := make([][]int, 0)
	cacheTriplet := make(map[Triplet]int)
	lenNums := len(nums)
	for i := 0; i < lenNums; i++ {
		for j := i + 1; j < lenNums; j++ {
			n1, n2 := nums[i], nums[j]
			cache[n1]--
			cache[n2]--
			x := -n1 - n2
			xCount, ok := cache[x]
			cache[n1]++
			cache[n2]++
			if xCount < 1 {
				continue
			}

			tripletArray := []int{n1, n2, x}
			sort.Ints(tripletArray)
			triplet := Triplet{
				n1: tripletArray[0],
				n2: tripletArray[1],
				n3: tripletArray[2],
			}
			_, ok = cacheTriplet[triplet]
			if !ok {
				cacheTriplet[triplet] = 1
				result = append(result, tripletArray)
			}
		}
	}

	return result
}

type Triplet struct {
	n1 int
	n2 int
	n3 int
}

// brute force
// func threeSum(nums []int) [][]int {
// 	lenNums := len(nums)
// 	result := make([][]int, 0)
// 	cache := make(map[Triplet]int)
// 	for i := 0; i < lenNums; i++ {
// 		for j := i + 1; j < lenNums; j++ {
// 			for k := j + 1; k < lenNums; k++ {
// 				fmt.Println(nums[i], nums[j], nums[k])
// 				tripletArray := []int{nums[i], nums[j], nums[k]}
// 				sort.Ints(tripletArray)
// 				triplet := Triplet{
// 					n1: tripletArray[0],
// 					n2: tripletArray[1],
// 					n3: tripletArray[2],
// 				}
// 				_, ok := cache[triplet]
// 				if !ok && nums[i]+nums[j]+nums[k] == 0 {
// 					cache[triplet] = 1
// 					result = append(result, tripletArray)
// 				}
// 			}
// 		}
// 	}
// 	return result
// }
//
// type Triplet struct {
// 	n1 int
// 	n2 int
// 	n3 int
// }
