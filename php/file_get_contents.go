package php

import "io/ioutil"

// https://www.php.net/manual/zh/function.file-get-contents.php
// file_get_contents
// todo 这个只能获取本地地址内容  http 或者 https远程链接还是获取不到的内容的 需要使用网络包获取
func FileGetContents(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
