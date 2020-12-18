package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":50000")
	fmt.Println("socket server listeninng on port 50000")
	for {
		conn, _ := listener.Accept()
		fmt.Printf("%s <---> %s\n", conn.LocalAddr(), conn.RemoteAddr())
		io.WriteString(conn, `0,0,3,3`+"\n")
	}
}
