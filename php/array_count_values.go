package php

// https://www.php.net/manual/zh/function.array-count-values.php
// array_count_values
func ArrayCountValues(array []interface{}) map[interface{}]uint {
	result := make(map[interface{}]uint)
	for _, v := range array {
		if c, ok := result[v]; ok {
			result[v] = c + 1
		} else {
			result[v] = 1
		}
	}
	return result
}
