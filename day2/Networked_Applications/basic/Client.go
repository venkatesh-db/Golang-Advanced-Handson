
package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(input))
		if err != nil {
			log.Fatalf("write error: %v", err)
		}

		response := make([]byte, 1024)
		n, _ := conn.Read(response)

		log.Printf("server response: %s", string(response[:n]))
	}
}
