package php

import (
	"crypto/sha1"
	"encoding/hex"
)

// https://www.php.net/manual/zh/function.sha1.php
// sha1 加密
func Sha1(s string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(s))
	return hex.EncodeToString(sha1.Sum(nil))
}
