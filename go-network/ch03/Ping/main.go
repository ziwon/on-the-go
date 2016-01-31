package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
	return checkError(err)

	c, err := net.DialIP("ip4:icmp", nil, addr)
	return checkError(err)

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
	return checkError(err)

	buffer := make([]uint8, count)
	count, err = c.Read(buffer)
	return checkError(err)

	ID := (uint16(buffer[count-4]) << 8) + uint16(buffer[count-3])

	return icmp.icmpID == ID
}

func checkError(err error) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return false
	} else {
		return true
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}

	result := Ping(os.Args[1], 10)
	fmt.Println("Result", result)
}
