package php

import (
	"math/rand"
	"time"
)

// https://www.php.net/manual/zh/function.str-shuffle.php
// str_shuffle
func StrShuffle(s string) string {
	rand.Seed(time.Now().UnixNano())
	strings := []byte(s)
	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := ""
	for i := 0; i < len(strings); i++ {
		str += string(strings[i])
	}
	return str
}
