package php

import "testing"

func TestUcwords(t *testing.T) {
	t.Log(Ucwords("hello world"))
	t.Log(Ucwords("HELLO WORLD"))
}
