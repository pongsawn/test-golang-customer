package validx

import "strings"

// IsEqualInt เช็คว่า integer เท่ากันมั้ย
func IsEqualInt(a *uint8, b int) bool {
	aValue := 0
	if a != nil {
		aValue = int(*a)
	}
	return aValue == b
}

// IsEqualStr เช็คว่า string เท่ากันมั้ย
func IsEqualStr(a *string, b string) bool {
	aValue := ""
	if a != nil {
		aValue = *a
	}
	return aValue == b
}

// IsEqualStr เช็คว่า string + trim เท่ากันมั้ย
func IsEqualStrTrim(a *string, b string) bool {
	aValue := ""
	if a != nil {
		aValue = *a
	}
	return strings.TrimSpace(aValue) == strings.TrimSpace(b)
}
