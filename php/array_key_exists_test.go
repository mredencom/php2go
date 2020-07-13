package php

import (
	"errors"
	"testing"
)

func TestArrayKeyExists(t *testing.T) {
	ss := make(map[interface{}]interface{})
	ss["1"] = errors.New("hello")
	ss["2"] = "hello world"
	ss["3"] = 1
	ss[1] = 1.1
	ss[1.2] = 1.2
	ss["你好中国"] = "你好中国"
	t.Log(ArrayKeyExists(1, ss))
	t.Log(ArrayKeyExists(11, ss))

}
