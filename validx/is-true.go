package validx

func IsTrue(s *bool) bool {
	if s == nil {
		return false
	}
	return *s
}
