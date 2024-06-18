package stringx

func SubStringLeft(s string, length int) string {
	return SubString(s, 0, length)
}

func SubStringRight(s string, length int) string {
	// length = 2
	len := len([]rune(s)) // ทดสอบ = 5
	start := len - length // = 5 - 2 = 3
	return SubString(s, start, len)
}

func SubString(s string, start, length int) string {
	// value := "bird"
	// runes := []rune(value)
	// safeSubstring := string(runes[1:3])
	// safeSubstring = ir
	var r string

	runes := []rune(s)
	len := len(runes)

	if start < 0 {
		start = 0
	}

	if start <= len-1 {
		if start+length > len {
			r = string(runes[start:len])
		} else {
			// เริ่มตั้งแต่ index:start จนถึงก่อน index:start+length
			r = string(runes[start : start+length])
		}
	}

	return r
}
