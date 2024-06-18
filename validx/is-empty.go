package validx

import (
	"strings"

	"github.com/google/uuid"
)

func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func IsValue(s string) bool {
	return !IsEmpty(s)
}

func IsEmptyPtr(s *string) bool {
	return s == nil || strings.TrimSpace(*s) == ""
}

func IsValuePtr(s *string) bool {
	return !IsEmptyPtr(s)
}

// IsUUID ตรวจสอบ *uuid
func IsEmptyPtrUUID(uuidx *uuid.UUID) bool {
	if uuidx == nil || *uuidx == uuid.Nil {
		return true
	}
	return false
}
