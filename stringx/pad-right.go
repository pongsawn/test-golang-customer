// https://github.com/pinpt/go-common/blob/v9.1.81/strings/strutil.go#L12

package stringx

import "bytes"

// PadRight returns a new string of a specified length in which the end of the current string is padded with spaces or with a specified Unicode character.
func PadRight(str string, length int, pad byte) string {
	if len(str) >= length {
		return str
	}
	buf := bytes.NewBufferString(str)
	for i := 0; i < length-len(str); i++ {
		buf.WriteByte(pad)
	}
	return buf.String()
}
