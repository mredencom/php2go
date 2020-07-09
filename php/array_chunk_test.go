package php

import (
	"testing"
)

func TestArrayChunk(t *testing.T) {
	var arr []interface{}
	arr = append(arr, 1)
	arr = append(arr, 45)
	arr = append(arr, 22)
	arr = append(arr, 22)
	arr = append(arr, 89)
	ret := ArrayChunk(arr, 3)
	t.Log(ret)
}
