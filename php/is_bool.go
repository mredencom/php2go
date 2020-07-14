package php

// https://www.php.net/manual/zh/function.is-bool.php
// is_bool
func IsBool(v interface{}) bool {
	switch v.(type) {
	case bool:
		return true
	default:
		return false
	}
}
