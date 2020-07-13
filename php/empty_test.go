package php

import (
	"errors"
	"testing"
)

func TestEmpty(t *testing.T) {
	t.Log(Empty(""))
	t.Log(Empty(1))
	t.Log(Empty(errors.New("err")))
	t.Log(Empty(0))
}
