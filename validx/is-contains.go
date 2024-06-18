package validx

import (
	"strings"
	"test-golang/stringx"
)

func IsContains(s []string, c string) bool {
	return stringx.IsContains(s, c)
}

func IsContaini(s []int, c int) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}

func IsContainx(s *[]string, c *string) bool {
	// match case
	for _, v := range *s {
		if v == *c {
			return true
		}
	}
	// lower
	cLower := strings.ToLower(*c)
	if cLower != *c {
		for _, v := range *s {
			if v == cLower {
				return true
			}
		}
	}
	// upper
	cUpper := strings.ToUpper(*c)
	if cUpper != *c {
		for _, v := range *s {
			if v == cUpper {
				return true
			}
		}
	}
	// snake-case
	cSnake := stringx.SnakeCase(*c)
	if cSnake != *c && cSnake != cLower {
		for _, v := range *s {
			if v == cSnake {
				return true
			}
		}
	}
	return false
}
