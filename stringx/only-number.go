package stringx

import (
	"strings"
)

var NUMBERS = []string{`0`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`}

// เอาเฉพาะตัวเลข
func OnlyNumber(s string) string {
	numbers := []string{}
	for _, v := range s {
		n := string(v)
		if IsContains(NUMBERS, n) {
			numbers = append(numbers, n)
		}
	}
	v := strings.Join(numbers, ``)
	return v
}
