package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("server running on port", ln.Addr())
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		if conn != nil {
			fmt.Println("New Client Connected :", conn.LocalAddr())
		}

		io.WriteString(conn, fmt.Sprint("Welcome to kloudOne ..", "\n", "You are Connected to Server \n", time.Now()))

		conn.Close()
	}
}
