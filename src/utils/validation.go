package utils

import "reflect"

// IsPtr check if the value is a pointer type as pointers are not going to be supported by design
func IsPtr[A any](arg A) bool {
	if v := reflect.ValueOf(arg); v.Kind() == reflect.Ptr {
		return true
	}
	return false
}

// IsNilOrZeroValue check if the value is nil for nullable types or zero value for the None nullable types
// as both of them represents the same thing
func IsNilOrZeroValue[A any](arg A) bool {
	// Validate for the nullable types that the value is not null
	// and for the none nullable types the value is not zeroValue
	if v := reflect.ValueOf(arg); ((v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Map ||
		v.Kind() == reflect.Chan ||
		v.Kind() == reflect.Func) && v.IsNil()) || !v.IsValid() {
		return true
	}
	return false
}
