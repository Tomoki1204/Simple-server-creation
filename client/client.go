package main

import (
	"bufio"
	//"flag"
	"fmt"
	"net"
	"os"
)

func handleReader(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	
	for{
		conn, err := net.Dial("tcp", "127.0.0.1:8030")
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			return
		}
		fmt.Printf("Enter message to send: ")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprint(conn, msg)
		handleReader(conn)
	}
}