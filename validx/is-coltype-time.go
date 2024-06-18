package validx

import "strings"

func IsColtypeTime(ct string) bool {
	ct = strings.ToUpper(ct)
	return (ct == `DATE` ||
		ct == `TIME` ||
		ct == `TIMESTAMP` ||
		ct == `TIMESTAMPTZ` ||
		ct == `INTERVAL` ||
		strings.Contains(ct, `TIME`) ||
		strings.Contains(ct, `DATE`))

}
