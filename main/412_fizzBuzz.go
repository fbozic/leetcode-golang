package main

import "fmt"

// func main() {
// 	fmt.Println(fizzBuzz(3))
// 	fmt.Println(fizzBuzz(5))
// 	fmt.Println(fizzBuzz(15))
// }

func fizzBuzz(n int) []string {
	answer := make([]string, 0, n)
	for i := 0; i < n; i++ {
		target := i + 1
		if target%5 == 0 && target%3 == 0 {
			answer = append(answer, "FizzBuzz")
		} else if target%5 == 0 {
			answer = append(answer, "Buzz")
		} else if target%3 == 0 {
			answer = append(answer, "Fizz")
		} else {
			answer = append(answer, fmt.Sprintf("%d", target))
		}
	}

	return answer
}
