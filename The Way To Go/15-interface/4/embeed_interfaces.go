package main

type Lock interface {
    Lock()
    Unlock()
}

type File interface {
    Lock
    Close()
}

// 接口 File 包含了 Lock 的所有方法，它还额外有一个 Close() 方法。