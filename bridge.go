package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":50000")
	fmt.Println("socket server listeninng on port 50000")

	player1, _ := listener.Accept()
	fmt.Printf("player1: %s\n", player1.RemoteAddr())

	player2, _ := listener.Accept()
	fmt.Printf("player2: %s\n", player2.RemoteAddr())

	go io.Copy(player2, player1)
	io.Copy(player1, player2)
}
