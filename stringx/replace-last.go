package stringx

import "strings"

// replace string from last with n times
func ReplaceLast(s, old, new string, n int) string {
	return rev(strings.Replace(rev(s), rev(old), rev(new), n))
}
