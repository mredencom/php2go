package php

// https://www.php.net/manual/zh/function.array-fill-keys.php
// array_fill_keys()
func ArrayFillKeys(keys []interface{}, value interface{}) map[interface{}]interface{} {
	r := make(map[interface{}]interface{})
	for _, v := range keys {
		r[v] = value
	}
	return r
}
