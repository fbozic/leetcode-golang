package main

func longestPalindrome(s string) string {
	maxPalindrome := ""
	maxLen := 0

	sLen := len(s)
	for i := 0; i < sLen; i++ {

		left, right := i, i
		for {
			if left >= 0 && right < sLen && s[left] == s[right] {
				if (right - left + 1) > maxLen {
					maxPalindrome = s[left : right+1]
					maxLen = right - left + 1
				}
				left--
				right++
			} else {
				break
			}
		}

		left, right = i, i+1
		for {
			if left >= 0 && right < sLen && s[left] == s[right] {
				if (right - left + 1) > maxLen {
					maxPalindrome = s[left : right+1]
					maxLen = right - left + 1
				}
				left--
				right++
			} else {
				break
			}
		}
	}
	return maxPalindrome
}
