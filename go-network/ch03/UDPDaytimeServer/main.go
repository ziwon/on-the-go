package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {

	var buf [512]byte

	/*
		// ReadFromUDP reads a UDP packet from c, copying the payload into b.
		// It returns the number of bytes copied into b and the return address
		// that was on the packet.
		//
		// ReadFromUDP can be made to time out and return an error with
		// Timeout() == true after a fixed time limit; see SetDeadline and
		// SetReadDeadline.
		func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error) {
			if !c.ok() {
				return 0, nil, syscall.EINVAL
			}
			var addr *UDPAddr
			n, sa, err := c.fd.readFrom(b)
			switch sa := sa.(type) {
			case *syscall.SockaddrInet4:
				addr = &UDPAddr{IP: sa.Addr[0:], Port: sa.Port}
			case *syscall.SockaddrInet6:
				addr = &UDPAddr{IP: sa.Addr[0:], Port: sa.Port, Zone: zoneToString(int(sa.ZoneId))}
			}
			if err != nil {
				err = &OpError{Op: "read", Net: c.fd.net, Source: c.fd.laddr, Addr: c.fd.raddr, Err: err}
			}
			return n, addr, err
		}
	*/
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	daytime := time.Now().String()

	/*
		func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error) {
			if !c.ok() {
				return 0, syscall.EINVAL
			}
			if c.fd.isConnected {
				return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: ErrWriteToConnected}
			}
			if addr == nil {
				return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: nil, Err: errMissingAddress}
			}
			sa, err := addr.sockaddr(c.fd.family)
			if err != nil {
				return 0, &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
			}
			n, err := c.fd.writeTo(b, sa)
			if err != nil {
				err = &OpError{Op: "write", Net: c.fd.net, Source: c.fd.laddr, Addr: addr.opAddr(), Err: err}
			}
			return n, err
		}
	*/
	conn.WriteToUDP([]byte(daytime), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
