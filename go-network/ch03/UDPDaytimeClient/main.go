package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	/*
		func sockaddrToUDP(sa syscall.Sockaddr) Addr {
			switch sa := sa.(type) {
			case *syscall.SockaddrInet4:
				return &UDPAddr{IP: sa.Addr[0:], Port: sa.Port}
			case *syscall.SockaddrInet6:
				return &UDPAddr{IP: sa.Addr[0:], Port: sa.Port, Zone: zoneToString(int(sa.ZoneId))}
			}
			return nil
		}

		func dialUDP(net string, laddr, raddr *UDPAddr, deadline time.Time) (*UDPConn, error) {
			fd, err := internetSocket(net, laddr, raddr, deadline, syscall.SOCK_DGRAM, 0, "dial")
			if err != nil {
				return nil, &OpError{Op: "dial", Net: net, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
			}
			return newUDPConn(fd), nil
		}
	*/
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	_, err = conn.Write([]byte("anything"))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
