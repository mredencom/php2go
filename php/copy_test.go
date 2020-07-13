package php

import "testing"

func TestCopy(t *testing.T) {
	Copy("testdata/file_get_contents.txt", "testdata/file_get_contents1111.txt")
}
