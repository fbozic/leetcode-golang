package main

// func main() {
// 	fmt.Println(generateParenthesis(2)) // 1
// 	fmt.Println(generateParenthesis(1)) // 1
// 	fmt.Println(generateParenthesis(3)) // 1
// }

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	recursiveGenerateParenthesis(&result, "", n, 0, 0, 1001)
	return result
}

func recursiveGenerateParenthesis(result *[]string, combo string, n, l, r int, runId int) {
	if len(combo) == 2*n {
		*result = append(*result, combo)
		return
	}
	var combo2 string
	// fmt.Println("a", runId, l, r, combo2)
	if l < n {
		combo2 = combo + "("
		// fmt.Println("1", runId, l, r, combo2)
		recursiveGenerateParenthesis(result, combo2, n, l+1, r, runId+1)
	}
	// fmt.Println("x", runId, l, r, combo2)
	if l > r {
		combo2 = combo + ")"
		// fmt.Println("2", runId, l, r, combo2)
		recursiveGenerateParenthesis(result, combo2, n, l, r+1, runId+1)
	}
}
