package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)

	max := 0
	substr := ""
	head := 0
	for i := 0; i < len(s); i++ {
		letter := fmt.Sprintf("%c", s[i])

		if m[letter] > 0 {
			o := m[letter]
			if head > o {
				substr = fmt.Sprint(s[head:i]) + letter
			} else {
				substr = fmt.Sprint(s[o:i]) + letter
				head = o
			}
			if o == i {
				substr = letter
				head = i
			}

		} else {
			substr = substr + letter
		}
		m[letter] = i + 1

		if len(substr) > max {
			max = len(substr)
		}
	}

	if len(substr) >= len(m) {
		return len(m)
	}

	if len(substr) > max {
		return len(substr)
	}
	return max
}
