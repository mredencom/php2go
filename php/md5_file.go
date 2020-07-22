package php

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/http"
)

// https://www.php.net/manual/zh/function.md5-file.php
// md5_file
func Md5File(path string) (string, error) {
	rp, err := http.Get(path)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(rp.Body)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	if _, err = io.Copy(hash, bytes.NewReader(body)); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5Str := hex.EncodeToString(hashInBytes)
	return md5Str, nil
}
