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
}

func (this *CLI) sendMessage(msg string) {
	this.Socket.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (this *CLI) Loop() {
	fmt.Printf("new CLI connection")
	defer fmt.Printf("CLI connection closed")
	for {
		_, message, err := this.Socket.ReadMessage()

		if err != nil {
			this.Socket.Close()
			return
		}

		split := strings.Fields(string(message))

		switch split[0] {
		case "echo":
			this.echo(split)
		default:
			this.sendMessage("unknown command")
		}
	}
}
