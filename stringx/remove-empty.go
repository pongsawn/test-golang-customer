package stringx

import "strings"

func RemoveEmpty(s []string) []string {

	items := []string{}
	for _, v := range s {
		if strings.TrimSpace(v) != "" {
			items = append(items, v)
		}
	}
	return items
}
