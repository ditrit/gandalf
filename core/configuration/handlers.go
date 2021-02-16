// Package cmd is part of Gandalf
package configuration

import (
	"fmt"
	"strings"
)

// CheckNotEmpty is a helper function to check a string value is not empty
var CheckNotEmpty = func(val interface{}) bool {
	strVal, ok := val.(string)
	ret := ok && strVal != ""
	if !ret {
		fmt.Printf("Error: empty value is not allowed")
	}
	return ret
}

// TrimToLower is a helper function to normalize value
var TrimToLower = func(val interface{}) interface{} {
	strVal, ok := val.(string)
	if ok {
		return strings.ToLower(strings.TrimSpace(strVal))
	}
	return val
}

var CheckTcpHighPort = func(val interface{}) bool {
	intVal, ok := val.(int)
	if ok {
		return intVal >= 1024 && intVal <= 65535
	}
	fmt.Printf("Error: value (%s) is not a TCP high port (between 1024 and 65535) \n", intVal)
	return false
}
