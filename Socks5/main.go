package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:3456")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Println("server.Accept failed, err:", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			fmt.Println("reader.ReadByte failed, err:", err)
			break
		}
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}
