package php

import (
	"log"
	"strconv"
	"testing"
)

/**
 * go test -v -run TestArrayDiff ./php
 */
func TestArrayDiff(t *testing.T) {
	array := []interface{}{"green", "red", "blue", "red", "red2"}
	array2 := []interface{}{"green", "blue"}
	array3 := []interface{}{"green", "blue"}

	resultArr := ArrayDiff(array, array2, array3)

	if len(resultArr) == 2 {
		checkMap := map[string]int{"red": 0, "red2": 0}
		for _, value := range resultArr {
			checkMap[value.(string)] += 1
		}
		for key, value := range checkMap {
			if value != 1 {
				log.Fatal("数量不对", key, ":", value)
			}
		}

	} else {
		log.Fatal("array lenght should 2 ,but get " + strconv.Itoa(len(resultArr)))
	}

	resultArr2 := ArrayDiff(array, array)
	if len(resultArr2) != 0 {
		log.Fatal("array diff itself should get a nil arr")
	}

}
