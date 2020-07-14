package php

import "testing"

func TestIsReadable(t *testing.T) {
	t.Log(IsWritable("testdata/file_get_contents.txt"))
}
