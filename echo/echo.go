package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go timefunc(conn)
	}
}

func timefunc(conn net.Conn) {
	defer conn.Close()
	for {
		str := time.Now().Format(time.UnixDate + "\n")
		io.WriteString(conn, str)
		time.Sleep(1 * time.Second)
	}

}
