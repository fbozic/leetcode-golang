package main

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int][]int)

	for i, num := range nums {
		_, ok := numsMap[num]
		if ok {
			numsMap[num] = append(numsMap[num], i)
		} else {
			numsMap[num] = []int{i}
		}

		x := target - num
		xIndices, ok := numsMap[x]
		if !ok {
			continue
		}

		if x != num {
			return []int{i, xIndices[0]}
		}
		if x == num && len(xIndices) > 1 {
			return []int{xIndices[0], xIndices[1]}
		}
	}
	return nil
}
