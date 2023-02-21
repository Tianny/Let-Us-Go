package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var client = http.Client{}

func callBoth(ctx context.Context, errVal string, slowURL string, fastUrl string) {
	ctx, cancel := context.WithCancel(ctx)
	// 创建 context都会获得一个相关的取消函数，在完成处理后，无论函数是否因发生错误而结束，都必须调用该取消函数。否则程序就会发生资源（内存和goroutine）泄露，并最终导致执行减慢或崩溃。多次调用取消函数不会发生错误，因为第一次调用取消函数之后，任何调用都不会起作用。
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowURL)
		if err != nil {
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastUrl+"?error="+errVal)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
	fmt.Println("done with both")
}

func callServer(ctx context.Context, label string, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "request err:", err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "response err:", err)
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(label, "read err:", err)
		return err
	}
	result := string(data)
	if result != "" {
		fmt.Println(label, "result:", result)
	}
	if result != "error" {
		fmt.Println("cancelling from", label)
		return errors.New("error happened")
	}
	return nil
}
