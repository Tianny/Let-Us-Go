想确保 parse 只被初始化一次，所以我们从一个闭包中设置 parser 的值，这个闭包被传递到 once 的 Do 方法中。
Parse 被调用一次后，once.Do 就会再执行这个闭包。