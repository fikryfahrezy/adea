package timeconv

import (
	"fmt"
	"strconv"
)

func AMPMto24(s string) string {
	sLen := len(s)
	firstTwo := string(s[0]) + string(s[1])
	lastTwo := string(s[sLen-2]) + string(s[sLen-1])
	s = (s[:sLen-2])[2:]

	if lastTwo == "AM" {
		if firstTwo == "12" {
			return "00" + s
		}

		return firstTwo + s
	}

	if firstTwo == "12" {
		return "12" + s
	}

	i, _ := strconv.Atoi(firstTwo)
	i += 12

	return fmt.Sprintf("%d%s", i, s)
}
