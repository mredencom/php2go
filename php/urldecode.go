package php

import "net/url"

// https://www.php.net/manual/zh/function.urldecode.php
// urldecode()
func Urldecode(s string) string {
	r, err := url.QueryUnescape(s)
	if err != nil {
		return ""
	}
	return r
}
