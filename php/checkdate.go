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

// https://www.php.net/manual/zh/function.checkdate.php
// checkdate()
func Checkdate(m, d, y int) bool {
	if m < 1 || m > 12 || d < 1 || d > 31 || y < 1 || y > 32767 {
		return false
	}
	switch m {
	case 4, 6, 9, 11:
		if d > 30 {
			return false
		}
	case 2:
		if y%4 == 0 && (y%100 != 0 || y%400 == 0) {
			if d > 29 {
				return false
			}
		} else if d > 28 {
			return false
		}
	}
	return true
}
