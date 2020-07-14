package php

// https://www.php.net/manual/zh/function.array_keys.php
// array_keys()
func ArrayKeys(arr map[interface{}]interface{}) (r []interface{}) {
	if len(arr) == 0 {
		return r
	}
	for k := range arr {
		r = append(r, k)
	}
	return r
}
