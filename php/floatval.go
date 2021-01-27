package php

import (
	"encoding/json"
	"errors"
	"strconv"
)

func Floatval(_var interface{}) (float64, error) {
	switch t := _var.(type) {
	default:
		return float64(0), errors.New("invalid num")
	case bool:
		if t {
			return float64(1), nil
		} else {
			return float64(0), nil
		}
	case nil:
		return float64(0), errors.New("invalid nil")
	case int:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		// int32 is rune
		return float64(t), nil
	case int64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	case uint:
		return float64(t), nil
	case uint16:
		return float64(t), nil
	case uint32:
		return float64(t), nil
	case uint64:
		return float64(t), nil
	case json.Number:
		num, err := t.Int64()
		return float64(num), err
	case string:
		return strconv.ParseFloat(t, 64)
	case uintptr:
		return float64(t), nil
	case byte:
		return float64(t), nil
	}
}

func Doubleval(_var interface{}) (float64, error) {
	return Floatval(_var)
}
