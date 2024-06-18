package stringx

import "strings"

// split with trim space and ignore empty string
func Split(s string, sep string) []string {

	strim := strings.TrimSpace(s)
	if strim == `` {
		return []string{}
	}

	// split with trim
	itemx := strings.Split(strim, sep)
	for i := 0; i < len(itemx); i++ {
		itemx[i] = strings.TrimSpace(itemx[i])
	}

	return itemx
}
