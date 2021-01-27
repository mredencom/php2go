package php

import (
	"encoding/json"
	"strconv"
	"errors"
)

func Int64val(_var interface{}) (int64, error) {
	switch t := _var.(type) {
	default:
		return 0, errors.New("invalid num")
	case bool:
		if t {
			return 1, nil
		} else {
			return 0, nil
		}
	case nil:
		return 0, errors.New("invalid nil")
	case int:
		return int64(t), nil
	case int16:
		return int64(t), nil
	case int32:
		return int64(t), nil
	case int64:
		return t, nil
	case float32:
		return int64(t), nil
	case float64:
		return int64(t), nil
	case uint:
		return int64(t), nil
	case uint16:
		return int64(t), nil
	case uint32:
		return int64(t), nil
	case uint64:
		return int64(t), nil
	case json.Number:
		num, err := t.Int64()
		return num, err
	case string:
		return strconv.ParseInt(t, 10, 64)
	case uintptr:
		return int64(t), nil
	case byte:
		return int64(t), nil
	}
}

func Intval(_var interface{}) (int32, error) {
	switch t := _var.(type) {
	default:
		return 0, errors.New("invalid num")
	case bool:
		if t {
			return 1, nil
		} else {
			return 0, nil
		}
	case nil:
		return 0, errors.New("invalid nil")
	case int:
		return int32(t), nil
	case int16:
		return int32(t), nil
	case int32:
		return t, nil
	case int64:
		return int32(t), nil
	case float32:
		return int32(t), nil
	case float64:
		return int32(t), nil
	case uint:
		return int32(t), nil
	case uint16:
		return int32(t), nil
	case uint32:
		return int32(t), nil
	case uint64:
		return int32(t), nil
	case json.Number:
		num, err := t.Int64()
		return int32(num), err
	case string:
		num, err := strconv.Atoi(t)
		return int32(num), err
	case uintptr:
		return int32(t), nil
	case byte:
		return int32(t), nil
	}
}
