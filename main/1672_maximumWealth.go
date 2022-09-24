package main

func maximumWealth(accounts [][]int) int {
	max := -1
	for _, customerBankAccounts := range accounts {
		currentCustomerBalance := 0
		for _, accountBalance := range customerBankAccounts {
			currentCustomerBalance += accountBalance
		}
		if currentCustomerBalance > max {
			max = currentCustomerBalance
		}
	}
	return max
}
