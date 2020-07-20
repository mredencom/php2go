package php

import "testing"

func TestStrReplace(t *testing.T) {
	str := "This is Tom ,who are you! Tom"
	t.Log(StrReplace(str, "Tom", "tom", 3))
}
