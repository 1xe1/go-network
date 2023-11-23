package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:5000")
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Connected to server")

    reader := bufio.NewReader(os.Stdin)
    for {
        // Read user input
        fmt.Print("Enter message: ")
        message, _ := reader.ReadString('\n')

		if  message[:len(message)-2] == "quit" {
			break
		}

        // Send the message to the server
        conn.Write([]byte(message))

        // Print the number of bytes sent
        fmt.Printf("Sent %d bytes\n", len(message))

        // Receive and print the server's response
        buffer := make([]byte, 1024)
        n, err := conn.Read(buffer)
        if err != nil {
            fmt.Println("Error reading:", err)
            return
        }
        fmt.Printf("Server response: %s", buffer[:n])
    }
}
