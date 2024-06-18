package validx

import "reflect"

func IsNil(v any) bool {

	// https://mangatmodi.medium.com/go-check-nil-interface-the-right-way-d142776edef1

	if v == nil {
		return true
	}

	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		rv := reflect.ValueOf(v)
		if rv.IsNil() {
			return true
		}
		return IsNil(reflect.Indirect(rv).Interface())
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Map, reflect.Chan, reflect.Slice:
		isNil := reflect.ValueOf(v).IsNil()
		return isNil
	}
	return false

}
