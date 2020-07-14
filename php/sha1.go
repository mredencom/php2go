package php

import (
	"crypto/sha1"
	"encoding/hex"
)

// https://www.php.net/manual/zh/function.sha1.php
// sha1() 加密
func Sha1(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
