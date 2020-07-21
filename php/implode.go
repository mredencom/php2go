package php

// https://www.php.net/manual/zh/function.implode.php
// implode
func Implode(sep string, ss []string) string {
	return Join(sep, ss)
}
