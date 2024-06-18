package validx

import "strings"

func IsColtypeFloat(ct string) bool {
	ct = strings.ToUpper(ct)
	return ct == `DECIMAL` ||
		ct == `NUMERIC` ||
		ct == `NUMBER` ||
		ct == `FLOAT` ||
		ct == `FLOAT32` ||
		ct == `FLOAT64` ||
		ct == `MONEY` ||
		ct == `DOUBLE` ||
		strings.Contains(ct, `FLOAT`) ||
		strings.Contains(ct, `DECIMAL`) ||
		strings.Contains(ct, `NUMERIC`) ||
		strings.Contains(ct, `NUMBER`)
}
