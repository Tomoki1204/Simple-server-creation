package main

import(
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for{
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
		fmt.Fprintln(conn, "ok")
	}
}

func main(){
	ln, err:= net.Listen("tcp", ":8030")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	for{
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}
