package main

import (
	"io"
	"log"
	"net"
	"os"
)

type ConnCtrl struct {
	conn net.Conn
}

func NewConnCtrl(conn net.Conn) *ConnCtrl {
	return &ConnCtrl{
		conn: conn,
	}
}

func (c *ConnCtrl) CloseRead() {
	if conn, ok := c.conn.(*net.TCPConn); ok {
		conn.CloseRead()
	}
}

func (c *ConnCtrl) CloseWrite() {
	if conn, ok := c.conn.(*net.TCPConn); ok {
		conn.CloseWrite()
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	ctrl := NewConnCtrl(conn)

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		ctrl.CloseRead()
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	ctrl.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
