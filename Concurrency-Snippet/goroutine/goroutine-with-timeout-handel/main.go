package main

import (
	"errors"
	"fmt"
	"time"
)

func doSomeWork() {
	fmt.Println("doing something")
}

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})

	go func() {
		doSomeWork()
		close(done)
	}()

	select {
	case <-done:
		return result, err
	case <-time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}
