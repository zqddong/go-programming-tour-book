package server

import (
	"github.com/zqddong/go-programming-tour-book/chatroom/logic"
	"net/http"
)

func RegisterHandle() {
	// 广播消息处理
	go logic.Broadcaster.Start()

	http.HandleFunc("/", HomeHandleFunc)
	http.HandleFunc("/ws", WebSocketHandleFunc)
	http.HandleFunc("/user_list", UserListHandleFunc)
}
