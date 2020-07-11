package php

import "testing"

func TestFileGetContents(t *testing.T) {
	b, _ := FileGetContents("testdata/file_get_contents.txt")
	t.Log(string(b))
}
