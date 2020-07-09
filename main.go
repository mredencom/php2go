package main

import (
	"fmt"
	"php2go/php"
)

func main() {
	var arr []interface{}
	arr = append(arr, 1)
	arr = append(arr, 45)
	arr = append(arr, 22)
	arr = append(arr, 22)
	arr = append(arr, 89)
	ret := php.ArrayChunk(arr, 2)
	fmt.Println(ret)
}
