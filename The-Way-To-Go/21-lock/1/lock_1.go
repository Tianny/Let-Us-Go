package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 1 // I'm blocked because there is no channel read yet.
	fmt.Println("send")
	go func() {
		<-ch // I will never be called for the main routine is blocked!
		fmt.Println("received")
	}()
	fmt.Println("over")
}
