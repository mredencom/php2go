package php

import "net/url"

// https://www.php.net/manual/zh/function.rawurlencode.php
// rawurlencode()
func Rawurlencode(s string) string {
	return url.PathEscape(s)
}
