package validx

import "strings"

func IsColtypeString(ct string) bool {
	ct = strings.ToUpper(ct)
	return ct == `TEXT` ||
		ct == `VARCHAR` ||
		ct == `NVARCHAR` ||
		ct == `CHAR` ||
		strings.Contains(ct, `CHAR`)
}
