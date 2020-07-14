package php

import "testing"

func TestSort(t *testing.T) {
	as := make([]int, 0)
	as = append(as, 1)
	as = append(as, 0)
	as = append(as, 645)
	as = append(as, 3)
	Sort(as)
	t.Log(as)
}
