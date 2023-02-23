#### Channel

三种使用方式： 
- 只能接收（<-chan，only receive）
    ```golang
    func sendMessage(conn net.Conn, ch <-chan string) {
        for msg := range ch {
            fmt.Fprintln(conn, msg)
       }
   }
    ```
- 只能发送（chan<-， only send）
- 正常的既能发送也能接收的 channel

只能接收（<-chan，only receive）和只能发送（chan<-， only send）。
它们没法直接创建，而是通过正常（双向）channel 转换而来（会自动隐式转换）。
它们存在的价值，主要是避免 channel 被乱用。
上面代码中 ch <-chan string 就是为了限制在 sendMessage 函数中只从 channel 读数据，不允许往里写数据。