package timeconv_test

import (
	"testing"

	"github.com/fikryfahrezy/adea/algorithm/timeconv"
)

func TestAMPMto24(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
	}{
		{
			input:  "09:45:30PM",
			expect: "21:45:30",
		},
		{
			input:  "09:45:30AM",
			expect: "09:45:30",
		},
		{
			input:  "07:05:45PM",
			expect: "19:05:45",
		},
		{
			input:  "07:05:45AM",
			expect: "07:05:45",
		},
		{
			input:  "12:00:00AM",
			expect: "00:00:00",
		},
		{
			input:  "12:10:00AM",
			expect: "00:10:00",
		},
		{
			input:  "12:10:00PM",
			expect: "12:10:00",
		},
		{
			input:  "01:10:00PM",
			expect: "13:10:00",
		},
		{
			input:  "04:10:00PM",
			expect: "16:10:00",
		},
	}

	for _, c := range testCases {
		res := timeconv.AMPMto24(c.input)

		if res != c.expect {
			t.Fatalf("input: %v, resulting: %v, expect: %v", c.input, res, c.expect)
		}
	}
}
