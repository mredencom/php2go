package php

// https://www.php.net/manual/zh/function.array-map.php
// array_map
// slice
func ArrayMapSlice(f func(item interface{}) interface{}, arr []interface{}) (r []interface{}) {
	for _, v := range arr {
		r = append(r, f(v))
	}
	return r
}

// https://www.php.net/manual/zh/function.array-map.php
// array_map
// map
func ArrayMapMap(f func(item interface{}) interface{}, arr map[string]interface{}) (r []interface{}) {
	for _, v := range arr {
		r = append(r, f(v))
	}
	return r
}
