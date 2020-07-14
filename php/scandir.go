package php

import (
	"io/ioutil"
	"os"
)

// https://www.php.net/manual/zh/function.scandir.php
// scandir()
func Scandir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}
