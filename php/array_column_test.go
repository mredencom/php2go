package php

import (
	"strconv"
	"testing"
)

func TestArrayColumn(t *testing.T) {
	s := make(map[string]map[string]interface{})
	ss := make(map[string]interface{})
	ss["hello"] = "hello world"
	for i := 0; i <= 6; i++ {
		s[strconv.Itoa(i)] = ss
	}

	t.Log(ArrayColumn(s, "hello"))

}
