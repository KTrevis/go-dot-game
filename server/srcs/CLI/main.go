package cli

import (
	"fmt"
	"server/Client"
	"strings"

	"github.com/gorilla/websocket"
)

type CLI struct {
	Socket	*websocket.Conn
	Manager	*client.WebSocketManager
	split	[]string
}

func (this *CLI) sendMessage(msg string) {
	this.Socket.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (this *CLI) Loop() {
	defer fmt.Printf("CLI connection closed")

	m := map[string]func() {
		"account": this.account,
	}

	for {
		_, message, err := this.Socket.ReadMessage()

		if err != nil {
			this.Socket.Close()
			return
		}

		this.split = strings.Fields(string(message))
		
		if f := this.validArg(m); f != nil {
			f()
		}
	}
}
