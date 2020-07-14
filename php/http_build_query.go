package php

import "fmt"

// https://www.php.net/manual/zh/function.http-build-query.php
// http_build_query()
func HttpBuildQuery(s map[string]string) (rs string) {
	l := len(s)
	if l == 0 {
		return
	}
	var i int
	for k, v := range s {
		i++
		if i < l {
			rs += fmt.Sprintf("%s=%s&", k, v)
		} else {
			rs += fmt.Sprintf("%s=%s", k, v)
		}
	}
	return
}
