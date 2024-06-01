package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("Welcome, ", id.String())

	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "uuid: %s\n", id.String())

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(text) == "exit" {
			return
		}
		// send to socket
		fmt.Fprintf(conn, "Message sent: %s\n", text)
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Message from server: %s", message)

	}
}
