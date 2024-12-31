package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	socket, _, err := websocket.DefaultDialer.Dial("ws://localhost:81/cli", nil)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	defer socket.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			return
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}

		cmd := scanner.Text()

		if strings.TrimSpace(cmd) == "" {
			continue
		}

		socket.WriteMessage(websocket.TextMessage, []byte(cmd))

		_, msg, err := socket.ReadMessage()

		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}

		str := string(msg)
		fmt.Println(str)
	}

}
