package php

import "math"

type ArraySlice []interface{}

// 实现php中的array_chunk
// https://www.php.net/manual/zh/function.array-chunk.php
// todo 没有实现array_chunk第三个参数使用
func ArrayChunk(input ArraySlice, size int) ArraySlice {

	// 求出步长
	limiter := int(math.Ceil(float64(len(input)) / float64(size)))
	ret := make(ArraySlice, limiter)
	for i := 0; i < limiter-1; i++ {
		ret[i] = input[size*i : i*size+size]
	}
	// 取出最后一条
	ret[limiter-1] = input[size*limiter-size:]
	return ret
}
