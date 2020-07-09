package php

// 创建数组
func Array(v ...interface{}) interface{} {
	return v
}

// 获取数组长度
func Count(v []interface{}) int {
	return len(v)
}
