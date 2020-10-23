package php

import "strings"

// https://www.php.net/manual/zh/function.strtr.php
// strtr()
func Strtr(haystack string, params ...interface{}) string {
	al := len(params)
	if al == 1 {
		pairs := params[0].(map[string]string)
		length := len(pairs)
		if length == 0 {
			return haystack
		}
		old_new := make([]string, length*2)
		for o, n := range pairs {
			if o == "" {
				return haystack
			}
			old_new = append(old_new, o, n)
		}
		return strings.NewReplacer(old_new...).Replace(haystack)
	} else if al == 2 {
		from := params[0].(string)
		to := params[1].(string)
		tr_len, lt := len(from), len(to)
		if tr_len > lt {
			tr_len = lt
		}
		if tr_len == 0 {
			return haystack
		}

		str := make([]uint8, len(haystack))
		var xlat [256]uint8
		var i int
		var j uint8
		if tr_len == 1 {
			for i = 0; i < len(haystack); i++ {
				if haystack[i] == from[0] {
					str[i] = to[0]
				} else {
					str[i] = haystack[i]
				}
			}
			return string(str)
		}
		for {
			xlat[j] = j
			if j++; j == 0 {
				break
			}
		}
		for i = 0; i < tr_len; i++ {
			xlat[from[i]] = to[i]
		}
		for i = 0; i < len(haystack); i++ {
			str[i] = xlat[haystack[i]]
		}
		return string(str)
	}

	return haystack
}