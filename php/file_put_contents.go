package php

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// https://www.php.net/manual/zh/function.file-put-contents.php
// file_put_contents
func FilePutContents(filename string, data []byte) error {
	if dir := filepath.Dir(filename); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, data, 0644)
}
