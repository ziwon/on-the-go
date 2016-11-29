package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type server struct {
	name string
	host string
}

func main() {
	done := make(chan bool)

	for _, a := range os.Args[1:] {
		params := strings.Split(a, "=")
		s := &server{params[0], params[1]}
		go connect(s, done)
	}

	<-done
}

func connect(s *server, done chan bool) {
	conn, err := net.Dial("tcp", s.host)
	if err != nil {
		log.Fatal(err)
	}

	s.handleConn(conn)
	done <- true
}

func (s server) handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		fmt.Fprintf(os.Stdout, "%s\t%s\n", s.name, input.Text())
	}
}
