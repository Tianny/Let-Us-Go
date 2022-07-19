package main

import (
	// Go 编译器不允许声明导入某个包却不使用。
	// 下划线让编译器接受这类导入，并且调用对应包内的所有代码文件里定义的 init 函数。
	// 对这个程序来说， 这样做的目的是调用 matchers 包中的 rss.go 代码文件里的 init 函数，注册 RSS 匹配器，以便后用
	_ "LetUsGo/Go-In-Action/chapter2/samples/matchers"
	"LetUsGo/Go-In-Action/chapter2/samples/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
