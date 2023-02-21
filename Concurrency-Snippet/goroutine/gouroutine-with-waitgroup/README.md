当有多个 goroutine 向同一个通道写入时，需要确保被写入的通道只被关闭一次，否则会触发 panic。sync.WaitGroup 正好适用于此。
