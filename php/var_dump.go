package php

import "fmt"

// https://www.php.net/manual/zh/function.var-dump.php
// var_dump()
func VarDump(v ...interface{}) {
	fmt.Println(v...)
}
