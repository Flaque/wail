package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/urfave/cli"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func read(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	line := ""

	// Read infinitely
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}

		char := fmt.Sprintf("%c", input) // "%c" converts rune to string
		fmt.Printf("%c", input)          // Also print stuff
		line += char

		if char == "\n" {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
				log.Println(err)
				return
			}

			line = ""
		}

	}
}

func main() {
	http.HandleFunc("/", read)

	app := cli.NewApp()
	app.Name = "wail"
	app.Usage = "Usage: ./exampleServer | wail"
	app.Action = func(c *cli.Context) error {
		return http.ListenAndServe(":80", nil)
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
