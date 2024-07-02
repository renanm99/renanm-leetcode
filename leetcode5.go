package main

import "fmt"

func longestPalindrome(s string) string {

	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		charleft := fmt.Sprintf("%c", s[0])
		charright := fmt.Sprintf("%c", s[1])
		if charleft == charright {
			return fmt.Sprint(s[0:])
		}
		return fmt.Sprint(s[0:1])
	}

	left, right := 0, 0
	biggest := fmt.Sprint(s[0:1])
	palindrome := ""

	for i := 0; i < len(s); i++ {
		left = i
		right = i + 1
		if left >= 0 && right < len(s) {
			charleft := fmt.Sprintf("%c", s[left])
			charright := fmt.Sprintf("%c", s[right])

			for charleft == charright {
				palindrome = fmt.Sprint(s[left : right+1])
				if len(palindrome) > len(biggest) {
					biggest = palindrome
				}
				left--
				right++
				if left < 0 || right >= len(s) {
					break
				}
				charleft = fmt.Sprintf("%c", s[left])
				charright = fmt.Sprintf("%c", s[right])
			}

			left = i
			right = i + 1
			if left-1 >= 0 && right < len(s) {
				left--
				charleft = fmt.Sprintf("%c", s[left])
				charright = fmt.Sprintf("%c", s[right])
				for charleft == charright {
					palindrome = fmt.Sprint(s[left : right+1])
					if len(palindrome) > len(biggest) {
						biggest = palindrome
					}
					left--
					right++
					if left < 0 || right >= len(s) {
						break
					}
					charleft = fmt.Sprintf("%c", s[left])
					charright = fmt.Sprintf("%c", s[right])
				}
			}
		}
	}
	return biggest
}
