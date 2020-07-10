package php

import "errors"

type arrayCombineMap map[interface{}]interface{}

// https://www.php.net/manual/zh/function.array-combine.php
// array_combine()
func ArrayCombine(arrK ArraySlice, arrV ArraySlice) (arrayCombineMap, error) {
	if len(arrK) != len(arrV) {
		return nil, errors.New("keys array and values array different lengths")
	}
	if len(arrK) == 0 && len(arrV) == 0 {
		return nil, errors.New("keys or values length equal zero")
	}
	var temp = arrayCombineMap{}
	for k, v := range arrK {
		temp[v] = arrV[k]
	}
	return temp, nil
}
