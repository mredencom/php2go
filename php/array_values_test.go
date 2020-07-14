package php

import "testing"

func TestArrayValuesMap(t *testing.T) {
	as := make(map[interface{}]interface{})
	as["n"] = 1
	as["n1"] = "hahh"
	t.Log(ArrayValues(as))
}
