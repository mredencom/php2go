package php

// https://www.php.net/manual/zh/function.array-push.php
// array_push()
func ArrayPush(arr *[]interface{}, p ...interface{}) {
	*arr = append(*arr, p...)
}
