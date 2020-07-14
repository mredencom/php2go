package php

// https://www.php.net/manual/zh/function.is-integer.php
// is_integer
// int32
func IsInteger(v interface{}) bool {
	switch v.(type) {
	case int64, int32, int, int16, int8:
		return true
	default:
		return false
	}
}
