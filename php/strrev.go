package php

// https://www.php.net/manual/zh/function.strrev.php
// strrev
// 反转字符串
func Strrev(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
