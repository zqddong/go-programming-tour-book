package main

import (
	"fmt"
	"github.com/zqddong/go-programming-tour-book/chatroom/gloable"
	"github.com/zqddong/go-programming-tour-book/chatroom/server"
	"log"
	"net/http"
)

var (
	addr   = ":2022"
	banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |

Go 语言编程之旅 —— 一起用 Go 做项目：ChatRoom，start on：%s
`
)

func init() {
	gloable.Init()
}

func main() {
	fmt.Printf("banner+\n", addr)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
