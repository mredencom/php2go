package php

import "testing"

func TestStrShuffle(t *testing.T) {
	ss := "hello world"

	t.Log(StrShuffle(ss))
}
