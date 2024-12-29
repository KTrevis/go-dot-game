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
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		
		if strings.TrimSpace(scanner.Text()) != "" {
			socket.WriteMessage(websocket.TextMessage, []byte(scanner.Text()))
		}

		_, msg, err := socket.ReadMessage()

		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}

		str := string(msg)
		fmt.Println(str)
	}

}
