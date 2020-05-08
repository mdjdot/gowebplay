package controllers

import (
	"fmt"

	"golang.org/x/net/websocket"
)

// Echo websocket 处理
func Echo(ws *websocket.Conn) {
	var err error

	for {

		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
