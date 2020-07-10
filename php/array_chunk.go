// MIT License
//
// Copyright (c) 2020 KING
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
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
