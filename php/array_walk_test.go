package php

import (
	"fmt"
	"testing"
)

func TestArrayWalk(t *testing.T) {
	var a []interface{}

	a = append(a, "item:a")
	a = append(a, "item:1")
	a = append(a, "item:2")
	ArrayWalkSlice(&a, test)
}

func test(item interface{}, key interface{}) {

	fmt.Println("item:", item, "\n key:", key)
}
