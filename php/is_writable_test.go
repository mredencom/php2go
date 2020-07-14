package php

import "testing"

func TestIsWritable(t *testing.T) {
	t.Log(IsWritable("testdata/file_get_contents.txt"))
}
