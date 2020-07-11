package php

import "testing"

func TestFilePutContents(t *testing.T) {
	err := FilePutContents("testdata/file_put_contents.txt", []byte("你好牛B！\n 🇨🇳"))
	t.Log(err)
}
