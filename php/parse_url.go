package php

import "net/url"

// https://www.php.net/manual/zh/function.parse-url.php
// parse_url()
func ParseUrl(rawurl string) (*url.URL, error) {
	return url.Parse(rawurl)
}
