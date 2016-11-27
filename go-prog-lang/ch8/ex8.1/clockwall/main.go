package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type server struct {
	name string
	host string
}

func main() {
	for _, a := range os.Args[1:] {
		params := strings.Split(a, "=")
		s := &server{params[0], params[1]}
		go connect(s)
	}

	time.Sleep(60 * time.Second)
}

func connect(s *server) {
	conn, err := net.Dial("tcp", s.host)
	if err != nil {
		log.Fatal(err)
	}
	s.handleConn(conn)
}

func (s server) handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		fmt.Fprintf(os.Stdout, "%s %s\n", s.name, input.Text())
	}
}
