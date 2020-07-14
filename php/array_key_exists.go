package php

// https://www.php.net/manual/zh/function.array-key-exists.php
// array_key_exists()
func ArrayKeyExists(key interface{}, arr map[interface{}]interface{}) bool {
	if len(arr) == 0 {
		return false
	}
	for k := range arr {
		if key == k {
			return true
		}
	}
	return false
}
