package php

import "testing"

func TestUnlink(t *testing.T) {
	t.Log(Unlink("testdata/22.txt"))
}
