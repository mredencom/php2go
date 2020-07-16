package php

import "testing"

func TestStripos(t *testing.T) {
	str := "5452315asdHGDafa3sd4"
	t.Log(Stripos(str, "a", 1))
}
