package php

import "testing"

func TestRename(t *testing.T) {
	flag := Rename("testdata/1.txt", "testdata/2.txt")
	if !flag {
		t.Error("失败")
	} else {
		t.Log("成功")
	}
}
