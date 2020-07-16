package php

import (
	"fmt"
	"testing"
)

func TestArrayWalk(t *testing.T) {
	var a []interface{}

	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	ArrayWalkSlice(&a, test)
	t.Log(a)
}

func test(item interface{}, key interface{}) {
	fmt.Println("item:", item, "\n key:", key)
	fmt.Println(item.(int) + key.(int))
}
