package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net"
	"os"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		conn.Close()
	}
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		s := string(buf[0:])
		dieWhenMeetCtrl(s)
		fmt.Printf(spew.Sdump(s))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func dieWhenMeetCtrl(str string) {
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c == 0xff || c == 0xfb {
			fmt.Println("サヨナラ、出会ったから売れし~")
			os.Exit(0)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
