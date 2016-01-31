package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/internal/iana"
	"golang.org/x/net/ipv4"
	"net"
	"os"
	"time"
)

const (
	icmpV4EchoRequest = 8
	icmpV4EchoReply   = 0
	icmpV6EchoRequest = 128
	icmpV6EchoReply   = 129
)

type ICMP struct {
	icmpType  uint8
	icmpCode  uint8
	icmpCheck uint16
	icmpID    uint16
	icmpSeq   uint16
}

func checkSum(msg []byte) uint16 {
	sum := 0

	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n]<<8) + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	return uint16(^sum)
}

func Ping(URL string, Timeout int) bool {
	addr, err := net.ResolveIPAddr("ip", URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return false
	}

	c, err := net.DialIP("ip4:icmp", nil, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return false
	}

	defer c.Close()
	c.SetDeadline(time.Now().Add(time.Duration(Timeout) * time.Second))

	icmp := new(ICMP)
	icmp.icmpType = icmpV4EchoRequest
	icmp.icmpCode = 0
	icmp.icmpID = uint16(os.Getpid() & 0xFFFF)
	icmp.icmpSeq = 1

	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, icmp)
	icmp.icmpCheck = checkSum(buf.Bytes())
	buf.Reset()
	binary.Write(&buf, binary.BigEndian, icmp)

	count, err := c.Write(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return false
	}

	buffer := make([]uint8, count+20)
	count, err = c.Read(buffer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return false
	}

	ID := (uint16(buffer[count-4]) << 8) + uint16(buffer[count-3])
	return icmp.icmpID == ID
}

func Ping2(URL string, Timeout int) bool {
	addr, err := net.ResolveIPAddr("ip", URL)

	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return false
	}

	defer c.Close()
	c.SetDeadline(time.Now().Add(time.Duration(Timeout) * time.Second))

	message := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xFFFF, Seq: 1,
			Data: []byte("Hi"),
		},
	}

	buf, err := message.Marshal(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return false
	}
	if _, err := c.WriteTo(buf, addr); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return false
	}

	rb := make([]byte, 1500)
	n, peer, err := c.ReadFrom(rb)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return false
	}
	rm, err := icmp.ParseMessage(iana.ProtocolICMP, rb[:n])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return false
	}
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		fmt.Printf("got reflection from %v", peer)
	default:
		fmt.Printf("got %+v; want echo reply", rm)
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}

	result := Ping2(os.Args[1], 20)
	fmt.Println("Result: ", result)
}
