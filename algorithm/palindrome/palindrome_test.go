package palindrome_test

import (
	"testing"

	"github.com/fikryfahrezy/adea/algorithm/palindrome"
)

func TestPalindrome(t *testing.T) {
	t.Run("Test palindrome for string", func(t *testing.T) {
		testCases := []struct {
			input  string
			expect bool
		}{
			{
				input:  "kasur ini rusak",
				expect: true,
			},
			{
				input:  "rusak ini kasur",
				expect: true,
			},
			{
				input:  "aku kamu",
				expect: false,
			},
			{
				input:  "kasur",
				expect: false,
			},
			{
				input:  "rusak",
				expect: false,
			},
			{
				input:  "katak",
				expect: true,
			},
		}

		for _, v := range testCases {
			res := palindrome.PalindromeString(v.input)

			if res != v.expect {
				t.Fatalf("input: %v, resulting: %v, expect: %v", v.input, res, v.expect)
			}
		}
	})

	t.Run("Test palindrome for number", func(t *testing.T) {
		testCases := []struct {
			input  uint64
			expect bool
		}{
			{
				input:  111,
				expect: true,
			},
			{
				input:  1121,
				expect: false,
			},
			{
				input:  1221,
				expect: true,
			},
		}

		for _, c := range testCases {
			res := palindrome.PalindromeNumber(c.input)

			if res != c.expect {
				t.Fatalf("input: %v, resulting: %v, expect: %v", c.input, res, c.expect)
			}
		}
	})
}
