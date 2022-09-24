package main

// func main() {
// 	fmt.Println(canConstruct("a", "b"))
// 	fmt.Println(canConstruct("aa", "ab"))
// 	fmt.Println(canConstruct("aa", "aab"))
// }

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := creteMagazineMap(magazine)

	for _, char := range ransomNote {
		count, ok := magazineMap[char]
		if ok && count > 0 {
			magazineMap[char] -= 1
		} else {
			return false
		}
	}
	return true
}

func creteMagazineMap(magazine string) map[int32]int {
	magazineMap := make(map[int32]int)
	for _, char := range magazine {
		_, ok := magazineMap[char]
		if ok {
			magazineMap[char] += 1
		} else {
			magazineMap[char] = 1
		}
	}
	return magazineMap
}
