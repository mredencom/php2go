package php

import (
	"io"
	"os"
)

// https://www.php.net/manual/zh/function.copy.php
// copy
func Copy(source, dest string) bool {
	src, err := os.Open(source)
	defer src.Close()
	if err != nil {
		return false
	}
	dst, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	defer dst.Close()
	if err != nil {
		return false
	}
	_, err = io.Copy(dst, src)
	if err != nil || err != io.EOF {
		return false
	}
	return true
}
