package validx

import (
	"reflect"
	"strings"
	"time"
)

func IsNilOrEmpty(v any) bool {

	if IsNil(v) {
		return true // skip nil
	}

	// check pointer
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		rv := reflect.ValueOf(v)
		if rv.IsNil() {
			return true
		}
		return IsNilOrEmpty(reflect.Indirect(rv).Interface())
	}

	if x, ok := v.(int); ok && x == 0 {
		return true // skip zero
	}
	if x, ok := v.(int64); ok && x == 0 {
		return true // skip zero
	}
	if x, ok := v.(float64); ok && x == 0 {
		return true // skip zero
	}
	if x, ok := v.(string); ok && strings.TrimSpace(x) == "" {
		return true // skip empty
	}
	if x, ok := v.(bool); ok && !x {
		return true // skip false
	}
	if x, ok := v.(time.Time); ok && x.IsZero() {
		return true // skip zero
	}

	// not nil and not empty
	return false

}
