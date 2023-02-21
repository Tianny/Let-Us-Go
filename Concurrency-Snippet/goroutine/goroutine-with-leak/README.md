每当启动 goroutine 函数时，必须确保它最终会退出。

与变量不同，Go 运行时无法检测 goroutine 函数不再被使用。

如果 goroutine 不退出，即使它什么都不做，调度器仍会定期给它分配执行时间，这就拖慢了程序。即所谓的 goroutine 泄露。