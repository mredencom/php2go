package php

import (
	"fmt"
)
/**
	普通占位符
	占位符     说明                           举例                   输出
	%v      相应值的默认格式。            Printf("%v", people)   {zhangsan}，
	%+v     打印结构体时，会添加字段名     Printf("%+v", people)  {Name:zhangsan}
	%#v     相应值的Go语法表示            Printf("#v", people)   main.Human{Name:"zhangsan"}
	%T      相应值的类型的Go语法表示       Printf("%T", people)   main.Human
	%%      字面上的百分号，并非值的占位符  Printf("%%")            %
	布尔占位符
	占位符       说明                举例                     输出
	%t          true 或 false。     Printf("%t", true)       true
	整数占位符
	占位符     说明                                  举例                       输出
	%b      二进制表示                             Printf("%b", 5)             101
	%c      相应Unicode码点所表示的字符              Printf("%c", 0x4E2D)        中
	%d      十进制表示                             Printf("%d", 0x12)          18
	%o      八进制表示                             Printf("%d", 10)            12
	%q      单引号围绕的字符字面值，由Go语法安全地转义 Printf("%q", 0x4E2D)        '中'
	%x      十六进制表示，字母形式为小写 a-f         Printf("%x", 13)             d
	%X      十六进制表示，字母形式为大写 A-F         Printf("%x", 13)             D
	%U      Unicode格式：U+1234，等同于 "U+%04X"   Printf("%U", 0x4E2D)         U+4E2D

	字符串与字节切片
	占位符     说明                              举例                           输出
	%s      输出字符串表示（string类型或[]byte)   Printf("%s", []byte("Go语言"))  Go语言
	%q      双引号围绕的字符串，由Go语法安全地转义  Printf("%q", "Go语言")         "Go语言"
	%x      十六进制，小写字母，每字节两个字符      Printf("%x", "golang")         676f6c616e67
	%X      十六进制，大写字母，每字节两个字符      Printf("%X", "golang")         676F6C616E67
	其它标记
	占位符      说明                             举例          输出
	+      总打印数值的正负号；对于%q（%+q）保证只输出ASCII编码的字符。
											   Printf("%+q", "中文")  "\u4e2d\u6587"
	-      在右侧而非左侧填充空格（左对齐该区域）
	#      备用格式：为八进制添加前导 0（%#o）      Printf("%#U", '中')      U+4E2D
		   为十六进制添加前导 0x（%#x）或 0X（%#X），为 %p（%#p）去掉前导 0x；
		   如果可能的话，%q（%#q）会打印原始 （即反引号围绕的）字符串；
		   如果是可打印字符，%U（%#U）会写出该字符的
		   Unicode 编码形式（如字符 x 会被打印成 U+0078 'x'）。
	' '    (空格)为数值中省略的正负号留出空白（% d）；
		   以十六进制（% x, % X）打印字符串或切片时，在字节之间用空格隔开
	0      填充前导的0而非空格；对于数字，这会将填充移到正负号之后
 */
// https://www.php.net/manual/zh/function.printf.php
// printf
func Printf(s interface{}) {
	switch s.(type) {
	case byte:
		fmt.Printf("%s", s)
		break
	case int, int64, int32, int16, int8, float64, float32:
		fmt.Printf("%d", s)
		break
	case bool:
		fmt.Printf("%t", s)
		break
	case string:
		fmt.Printf("%s", s)
		break
	case struct{}:
		fmt.Printf("%+v", s)
		break
	default:
		fmt.Println(s)
	}
}
