package php

import "net/url"

// https://www.php.net/manual/zh/function.urlencode.php
// urlencode()
func Urlencode(s string) string {
	return url.QueryEscape(s)
}
