package php

// https://www.php.net/manual/zh/function.array.php
// array()
func Array(v ...interface{}) interface{} {
	return v
}

// 获取数组长度
func Count(v []interface{}) int {
	return len(v)
}
