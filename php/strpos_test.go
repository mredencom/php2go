package php

import "testing"

func TestStrpos(t *testing.T) {
	str := "5452315asdafa3sd4"
	t.Log(Strpos(str, "a", 1))
}
