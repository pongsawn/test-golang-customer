package validx

import "strings"

func IsColtypeInt(ct string) bool {
	ct = strings.ToUpper(ct)
	return strings.Contains(ct, `INT`) && (ct != `INTERVAL`)
}
