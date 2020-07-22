package php

// https://www.php.net/manual/zh/function.array-pop.php
// array_pop
func ArrayPop(s []interface{}) []interface{} {
	return s[:len(s)-1]
}
