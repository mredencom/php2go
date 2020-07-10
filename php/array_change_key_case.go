package php

import "strings"

const (
	CaseLower = iota
	CaseUpper
)

type ArrayMap map[string]interface{}

// https://www.php.net/manual/zh/function.array-change-key-case.php
// array_change_key_case()

func ArrayChangeKeyCase(arr ArrayMap, Case int) ArrayMap {

	var temp = ArrayMap{}
	for k, v := range arr {
		if Case == CaseLower {
			temp[strings.ToLower(k)] = v
		} else if Case == CaseUpper {
			temp[strings.ToUpper(k)] = v
		}
	}
	return temp
}
