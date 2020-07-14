package php

// https://www.php.net/manual/zh/function.array-merge.php
// array_merge
func ArrayMerge(arr ...interface{}) []interface{} {
	r := make([]interface{}, 0)
	return append(r, arr...)
}
