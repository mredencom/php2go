package php

// https://www.php.net/manual/zh/function.chunk-split.php
// chunk_split()
func ChunkSplit(body string, length uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, es := []rune(body), []rune(end)
	l := uint(len(runes))
	if l <= 1 || l < length {
		return body + end
	}
	ns := make([]rune, 0, len(runes)+len(es))
	var i uint
	for i = 0; i < l; i += length {
		if i+length > l {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+length]...)
		}
		ns = append(ns, es...)
	}
	return string(ns)
}
