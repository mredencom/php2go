package php

// https://www.php.net/manual/zh/function.array-walk.php
// array_walk
// slice
// todo 待完善
func ArrayWalkSlice(arr *[]interface{}, f func(item interface{}, key interface{})) {
	for k, v := range *arr {
		f(v, k)
	}
}

// https://www.php.net/manual/zh/function.array-walk.php
// array_walk
// map
// todo 待完善
func ArrayWalkMap(arr map[interface{}]interface{}, f func(item interface{}, key interface{})) {
	for k, v := range arr {
		f(v, k)
	}
}
