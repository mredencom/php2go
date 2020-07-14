package php

// https://www.php.net/manual/zh/function.is-int.php
// is_int
// int32
func IsInt(v interface{}) bool {
	switch v.(type) {
	case int32, int:
		return true
	default:
		return false
	}
}
