package php

// https://www.php.net/manual/zh/function.array-flip.php
// array_flip
func ArrayFlip(arr map[interface{}]interface{}) map[interface{}]interface{} {
	r := make(map[interface{}]interface{})
	for k, v := range arr {
		r[v] = k
	}
	return r
}
