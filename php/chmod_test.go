package php

import "testing"

func TestChmod(t *testing.T) {
	Chmod("testdata/file_get_contents.txt", 0777)
}
