package php

//https://www.php.net/manual/zh/function.array-diff.php
func ArrayDiff(array []interface{}, arrays ...[]interface{}) []interface{} {
	mapResult := map[interface{}]interface{}{}
	for _, value := range array {
		mapResult[value] = nil
	}
	for _, arrWaitCompare := range arrays {
		for _, valueWaitCompare := range arrWaitCompare {
			if _, ok := mapResult[valueWaitCompare]; ok {
				delete(mapResult, valueWaitCompare)
			}
		}
	}

	// 先make指定长度，否则会进行多次内存分配
	result := make([]interface{}, 0, len(mapResult))

	for valueInKey, _ := range mapResult {
		result = append(result, valueInKey)
		//fmt.Println(&result[0])
	}
	return result
}
