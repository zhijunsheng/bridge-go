package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", ":50000")
	fmt.Println("socket server listeninng on port 50000")
	for {
		conn, _ := listener.Accept()
		fmt.Printf("%s <---> %s\n", conn.LocalAddr(), conn.RemoteAddr())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for col := 0; col < 5; col++ {
		move := fmt.Sprintf("%d,1,%d,3\n", col, col)
		io.WriteString(conn, move)
		time.Sleep(2 * time.Second)
	}
}
