package php

import "fmt"

// https://www.php.net/manual/zh/function.print.php
// print()
func Print(v ...interface{}) {
	fmt.Print(v...)
}
