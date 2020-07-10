package php

import (
	"crypto/md5"
	"encoding/hex"
)
// https://www.php.net/manual/zh/function.md5.php
// md5 加密
func Md5(s string) string {
	md := md5.New()
	md.Write([]byte(s))
	return hex.EncodeToString(md.Sum(nil))
}
