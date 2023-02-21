package main

// 把相同的数据传递给多个函数，但只想得到执行的最快的函数的结果
func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	result := make(chan []string)

	for _, searcher := range searchers {
		go func(searcher func(string) []string) {
			select {
			case result <- searcher(s):
			case <-done:
				return
			}
		}(searcher)
	}

	r := <-result
	close(done)

	return r
}
