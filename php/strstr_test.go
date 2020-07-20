package php

import "testing"

func TestStrstr(t *testing.T) {
	st := "hello@gmail.com"
	t.Log(Strstr(st, "@"))
}
