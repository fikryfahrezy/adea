package findmostleast_test

import (
	"testing"

	"github.com/fikryfahrezy/adea/algorithm/findmostleast"
)

func TestFindMostLeast(t *testing.T) {
	testCases := []struct {
		input  string
		expect [2]string
	}{
		{
			input:  "sofware engineer",
			expect: [2]string{"e", "s"},
		},
		{
			input:  "sssss eeeee",
			expect: [2]string{"s", "s"},
		},
		{
			input:  "ssseee fff",
			expect: [2]string{"s", "s"},
		},
		{
			input:  "ssseee ffffffff",
			expect: [2]string{"f", "s"},
		},
		{
			input:  "ffffffff ssseee",
			expect: [2]string{"f", "s"},
		},
		{
			input:  "ffffffff                  ssseee",
			expect: [2]string{"f", "s"},
		},
		{
			input:  "ffffffff                  eeeessseee",
			expect: [2]string{"f", "s"},
		},
		{
			input:  "ffffffff                  eeesss",
			expect: [2]string{"f", "e"},
		},
		{
			input:  "ffffffff                  es es es",
			expect: [2]string{"f", "e"},
		},
		{
			input:  "fse fes fes fes fes fff",
			expect: [2]string{"f", "s"},
		},
	}

	for _, c := range testCases {
		res := findmostleast.FindMostLeast(c.input)

		for i, v := range res {
			if c.expect[i] != v {
				t.Fatalf("input: %v, resulting: %v, expect: %v,", c.input, res, c.expect)
			}
		}
	}
}
