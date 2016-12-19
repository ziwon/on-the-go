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
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	const timeout = 10 * time.Second
	timer := time.NewTimer(timeout)

	input := bufio.NewScanner(c)
	token := make(chan string)
	go func() {
		for input.Scan() {
			token <- input.Text()
		}
	}()

	for {
		select {
		case text := <-token:
			timer.Reset(timeout)
			go echo(c, text, 1*time.Second)
		case <-timer.C:
			c.Close()
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
