package validx

import "strings"

func IsColtypeUuid(ct string) bool {
	ct = strings.ToUpper(ct)
	return (ct == `UUID`)
}
