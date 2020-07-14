package php

// https://www.php.net/manual/zh/function.is-double.php
// is_double
// float64 float32
func IsDouble(v interface{}) bool {
	switch v.(type) {
	case float64, float32:
		return true
	default:
		return false
	}
}
