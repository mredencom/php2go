package php

import "testing"

func TestIsFile(t *testing.T) {
	t.Log(IsFile("testdata"))
	t.Log(IsFile("testdata/file_get_contents.txt"))
}
