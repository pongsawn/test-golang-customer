package stringx

import "strings"

func IsContains(s []string, c string) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}

func IsContainx(s *[]string, c string) bool {
	for _, v := range *s {
		if v == c {
			return true
		}
	}
	return false
}

func IsContainAny(s []string, c string) bool {
	for _, v := range s {
		if strings.EqualFold(v, c) {
			return true
		}
	}
	return false
}

func IsContainAnySnake(s []string, c string) bool {
	for _, v := range s {
		if strings.EqualFold(v, c) {
			return true
		}
	}
	for _, v := range s {
		if strings.EqualFold(SnakeCase(v), SnakeCase(c)) {
			return true
		}
	}

	return false
}
