package main

func numberOfSteps(num int) int {
	no := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		no++
	}
	return no
}
