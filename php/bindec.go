package php

import (
	"container/list"
	"math"
)

// https://www.php.net/manual/zh/function.bindec.php
// bindec()
func Bindec(t string) int64 {
	stack := list.New()
	var sum int64
	var bnum, inum float64 = 0, 2
	for _, c := range t {
		stack.PushBack(c)
	}
	for e := stack.Back(); e != nil; e = e.Prev() {
		v := e.Value.(int32)
		sum += int64(v-48) * int64(math.Pow(inum, bnum))
		bnum++
	}
	return sum
}
