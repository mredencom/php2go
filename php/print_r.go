package php

import "fmt"

// https://www.php.net/manual/zh/function.print-r.php
// print_r()
func PrintR(v ...interface{}) {
	fmt.Println(v...)
}
