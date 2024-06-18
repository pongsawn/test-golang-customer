package stringx

import (
	"strings"
)

func Appendx(s *[]string, c string) {
	c = strings.TrimSpace(c)
	if !IsEmpty(c) {
		if !IsContainx(s, c) {
			*s = append(*s, c)
		}
	}
}
