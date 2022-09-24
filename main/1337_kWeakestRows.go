package main

import (
	"sort"
)

// func main() {
// 	mat := [][]int{
// 		{1, 1, 0, 0, 0},
// 		{1, 1, 1, 1, 0},
// 		{1, 0, 0, 0, 0},
// 		{1, 1, 0, 0, 0},
// 		{1, 1, 1, 1, 1},
// 	}
// 	fmt.Println(kWeakestRows(mat, 3))
// }

func kWeakestRows(mat [][]int, k int) []int {
	// key->numberOfSolder, value->list[row]
	soldiersMap := make(map[int][]int)

	for i, row := range mat {
		soldiers := 0
		for _, rowElement := range row {
			if rowElement == 1 {
				soldiers += 1
			} else {
				break
			}
		}

		val, ok := soldiersMap[soldiers]
		if ok {
			soldiersMap[soldiers] = append(val, i)
		} else {
			soldiersMap[soldiers] = []int{i}
		}
	}

	keys := make([]int, 0, len(soldiersMap))
	for soldiers := range soldiersMap {
		keys = append(keys, soldiers)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	result := make([]int, 0, k)
	for _, key := range keys {
		rows := soldiersMap[key]
		sort.Slice(rows, func(i, j int) bool {
			return rows[i] < rows[j]
		})
		for _, row := range rows {
			result = append(result, row)
			if len(result) == k {
				break
			}
		}
		if len(result) == k {
			break
		}
	}
	return result
}
