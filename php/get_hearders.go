package php

import "net/http"

// https://www.php.net/manual/zh/function.get-headers.php
// get_headers()
func GetHeaders(u string) http.Header {
	var r http.Header
	if u == "" {
		return r
	}
	resp, err := http.Get(u)
	if err != nil {
		return r
	}
	return resp.Header
}
