package php

import "encoding/base64"

// https://www.php.net/manual/zh/function.base64-encode.php
// base64_encode
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// https://www.php.net/manual/zh/function.base64-encode.php
// base64_decode
func Base64Decode(s string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	return string(b), err
}
