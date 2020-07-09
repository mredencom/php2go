package php

// https://www.php.net/manual/zh/function.array-column.php
// array_column 获取指定的key值
// todo index第三个参数待实现
func ArrayColumn(arrayMap map[string]map[string]interface{}, column string) (ret []interface{}) {
	for _, v := range arrayMap {
		if v, ok := v[column]; ok {
			ret = append(ret, v)
		}
	}
	return
}
