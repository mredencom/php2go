package php

// https://www.php.net/manual/zh/function.array-values.php
// array-values()
func ArrayValues(arr map[interface{}]interface{}) []interface{} {
	r := make([]interface{}, 0)
	for _, v := range arr {
		r = append(r, v)
	}
	return r
}
