
package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	defer listener.Close()

	log.Println("server listening on port 8081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("connection error: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("client disconnected: %v", err)
			return
		}

		log.Printf("received: %s", message)

		_, err = conn.Write([]byte("ack: " + message))
		if err != nil {
			log.Printf("write error: %v", err)
			return
		}
	}
}