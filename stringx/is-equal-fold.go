package stringx

import "strings"

func IsEqualFold(s, c string) bool {
	return strings.EqualFold(strings.TrimSpace(s), strings.TrimSpace(c))
}

func IsEqualFolds(s, c []string) bool {
	if len(s) != len(c) {
		return false
	}
	if len(s) == 0 {
		return true
	}
	for _, v1 := range s {
		found := false
		for _, v2 := range c {
			if strings.EqualFold(v1, v2) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
