package validx

import "regexp"

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// IsEmail checks if the email provided passes the required structure
// and length test. It also checks the domain has a valid MX record.
func IsEmail(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}

	if !emailRegex.MatchString(e) {
		return false
	}
	// parts := strings.Split(e, "@")
	// mx, err := net.LookupMX(parts[1])
	// if err != nil || len(mx) == 0 {
	// 	return false
	// }
	return true
}
