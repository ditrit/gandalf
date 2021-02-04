// Package cmd is part of Gandalf
package configuration

import "strings"

// CheckNotEmpty is a helper function to check a string value is not empty
var CheckNotEmpty = func(val interface{}) bool {
	strVal, ok := val.(string)
	return ok && strVal != ""
}

// TrimToLower is a helper function to normalize value
var TrimToLower = func(val interface{}) interface{} {
	strVal, ok := val.(string)
	if ok {
		return strings.ToLower(strings.TrimSpace(strVal))
	}
	return val
}
