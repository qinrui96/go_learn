package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintf(c, "%s\n", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "%s\n", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "%s\n", strings.ToLower(shout))
}

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echo(conn, input.Text(), 1*time.Second)
	}
	conn.Close()
}

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
		go handleConn(conn)
	}
}
