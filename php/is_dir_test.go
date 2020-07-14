package php

import "testing"

func TestIsDir(t *testing.T) {
	t.Log(IsDir("testdata"))
	t.Log(IsDir("testdata/file_get_contents.txt"))
}
