package php

import "net/url"

// https://www.php.net/manual/zh/function.rawurldecode.php
// rawurldecode()
func Rawurldecode(s string) string {
	r, err := url.PathUnescape(s)
	if err != nil {
		return ""
	}
	return r
}
