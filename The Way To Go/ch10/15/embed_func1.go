package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log *Log
}

func (l *Log) Add(s string)  {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
	
}

func (c *Customer) Log() *Log {
	return c.log
	
}

func main() {
	c := &Customer{"Obama", &Log{"1 - yes we can"}}
	c.Log().Add("2 - after ")
	fmt.Println(c.Log())
}