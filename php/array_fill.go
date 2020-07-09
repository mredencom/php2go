package php

// https://www.php.net/manual/zh/function.array-fill.php
// array_fill
// 数据填充 具体操作请看php链接
func ArrayFill(startIndex int, num uint, value interface{}) map[int]interface{} {
	result := make(map[int]interface{})
	var i uint
	for i = 0; i < num; i++ {
		result[startIndex] = value
		startIndex++
	}
	return result
}
