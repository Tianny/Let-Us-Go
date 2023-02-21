package main

import (
	"fmt"
	"sync"
)

type SlowComplicatedParser interface {
	Parse(string) string
}

var parser SlowComplicatedParser
var once sync.Once

func Parse(dataToParse string) string {
	// Parse 被调用一次后，once.Do 就不会再执行这个闭包
	once.Do(func() {

		parser = initParser()
	})
	return parser.Parse(dataToParse)
}

func initParser() SlowComplicatedParser {
	// do all sorts of setup and loading here
	fmt.Println("do all sort of setup and loading here")
	return nil
}
