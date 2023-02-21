通道模式提供了一种方法来通知 goroutine 停止执行。它使用通道来发出信号，表明 goroutine 可以退出了。

为每个传入的 searcher 启动了一个 goroutine。

select 要么阻塞在 result 通道上的写操作，要么阻塞在等待读取 done 通道上的数据。

因为 done 和 result 通道都是无缓冲通道，r := <- result 将会阻塞在那里，直到第一个 search(s) 返回结果。

close(done) 关闭通道后，select 语句 case <-done 分支将被执行，最终 goroutine 退出，防止泄露。