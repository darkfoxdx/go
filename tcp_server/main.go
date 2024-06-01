package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {
		// Read data from the client
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Process and use the data (here, we'll just print it)
		fmt.Printf("Received: %s\n", string(message))

		// Return message
		data := []byte("Hello, Client!\n")
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("server: failed to write!")
		}
	}
}
