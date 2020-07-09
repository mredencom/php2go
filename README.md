## 用Go语言实现PHP内置函数

[![Build Status](https://travis-ci.org/php2go/php2go.svg?branch=master)](https://travis-ci.org/php2go/php2go)
[![Go Report Card](https://goreportcard.com/badge/github.com/php2go/php2go)](https://goreportcard.com/report/github.com/php2go/php2go)
[![GoDoc](https://godoc.org/github.com/php2go/php2go/php?status.svg)](https://godoc.org/github.com/php2go/php2go/php)
[![GitHub contributors](https://img.shields.io/github/contributors/php2go/php2go.svg)](https://github.com/php2go/php2go/graphs/contributors)
[![license](https://img.shields.io/github/license/php2go/php2go.svg)](https://github.com/php2go/php2go/blob/master/LICENSE)
[![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/php2go/php2go.svg?colorB=green)](https://github.com/php2go/php2go/archive/master.zip)

这是一个用Go语言开发的辅助库，尤其适用于熟悉PHP内置函数的开发者，将会实现PHP内置函数，让你写go如php般流畅，你值得拥有~

### 下载安装

```shell

go get -u github.com/php2go/php2go/php

```

### 关于命名

PHP下划线命名转为Go大驼峰命名。

### Example:

```go

package main

import (
    "github.com/php2go/php2go/php"
)

func main() {

    php.Echo("Hello ", "world!\n")

}

```

[More](https://github.com/php2go/php2go/blob/master/main.go)

### 项目进度

[TODO List](https://github.com/php2go/php2go/blob/master/TODO.md)

### 贡献代码

[贡献指南](https://github.com/php2go/php2go/blob/master/.github/CONTRIBUTING.md)

## Contributors

[Your contributions are always welcome!](https://github.com/php2go/php2go/graphs/contributors)

## LICENSE

Released under [MIT](https://github.com/php2go/php2go/blob/master/LICENSE) LICENSE
