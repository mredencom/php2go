package php

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/http"
)

// https://www.php.net/manual/zh/function.sha1-file.php
// sha1_file
func Sha1File(path string) (string, error) {
	rp, err := http.Get(path)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(rp.Body)
	if err != nil {
		return "", err
	}
	hash := sha1.New()
	if _, err = io.Copy(hash, bytes.NewReader(body)); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5Str := hex.EncodeToString(hashInBytes)
	return md5Str, nil
}
