package palindrome

import "strconv"

func PalindromeString(s string) bool {
	slen := len(s)
	for i := 0; i < slen/2; i++ {
		if s[i] != s[(slen-1)-i] {
			return false
		}
	}

	return true
}

func PalindromeNumber(i uint64) bool {
	s := strconv.FormatUint(i, 10)

	return PalindromeString(s)
}
