package validx

import "strings"

func IsColtypeBool(ct string) bool {
	return strings.Contains(strings.ToUpper(ct), `BOOL`)
}
