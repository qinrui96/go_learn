package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

type retinfo struct {
	num int
	err error
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan retinfo)

	wg := sync.WaitGroup{}
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err = os.Stdin.Read(buf[0:])
	for err == nil {
		wg.Add(1)
		tmp := make([]byte, 1024)
		copy(tmp, buf)
		go func(t []byte) {
			defer wg.Done()
			_, err = conn.Write(buf[0:])
			for err == nil {
				_, err = conn.Read(buf[0:])
				os.Stdout.Write(buf)
			}
		}(tmp)
		_, err = os.Stdin.Read(buf[0:])
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i.num)
	}
	conn.Close()
}
