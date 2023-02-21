在子 context 上设置的任何超时都受到在父 context 上设置的超时的限制。
如果父 context 在 2s 内超时，你可以声明子 context 在 3s 内超时；如果父 context 在 2s 后超时，子 context 也会超时。