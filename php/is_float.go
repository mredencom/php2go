package php

// https://www.php.net/manual/zh/function.is-float.php
// is_float
// float32
func IsFloat(v interface{}) bool {
	switch v.(type) {
	case float32:
		return true
	default:
		return false
	}
}
