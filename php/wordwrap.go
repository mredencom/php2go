package php

// https://www.php.net/manual/zh/function.wordwrap.php
// wordwrap()
func Wordwrap(s string, width uint, br string, cut bool) string {
	sl := len(s)
	bl := len(br)
	ll := int(width)

	if sl == 0 {
		return ""
	}
	if bl == 0 {
		panic("br is not empty")
	}
	if ll == 0 && cut {
		panic("width is not zero and cut true")
	}
	current, ls, lspace := 0, 0, 0
	var ns []byte
	for current = 0; current < sl; current++ {
		if s[current] == br[0] && current+bl < sl && s[current:current+bl] == br {
			ns = append(ns, s[ls:current+bl]...)
			current += bl - 1
			lspace = current + 1
			ls = lspace
		} else if s[current] == ' ' {
			if current-ls >= ll {
				ns = append(ns, s[ls:current]...)
				ns = append(ns, br[:]...)
				ls = current + 1
			}
			lspace = current
		} else if current-ls >= ll && cut && ls >= lspace {
			ns = append(ns, s[ls:current]...)
			ns = append(ns, br[:]...)
			ls = current
			lspace = current
		} else if current-ls >= ll && ls < lspace {
			ns = append(ns, s[ls:lspace]...)
			ns = append(ns, br[:]...)
			lspace++
			ls = lspace
		}
	}

	if ls != current {
		ns = append(ns, s[ls:current]...)
	}
	return string(ns)
}
