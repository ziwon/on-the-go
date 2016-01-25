package main

import (
	"fmt"
	"net"
	"os"
)

type IP []byte
type IPMask []byte

var (
	classAMask = IPv4Mask(0xff, 0, 0, 0)
	classBMask = IPv4Mask(0xff, 0xff, 0, 0)
	classCMask = IPv4Mask(0xff, 0xff, 0xff, 0)
)

func IPv4Mask(a, b, c, d byte) []byte {
	p := make(IP, 4)
	p[0] = a
	p[1] = b
	p[2] = c
	p[3] = d
	return p
}

func DefaultMask2(addr net.IP) IPMask {
	ip := addr.To4()
	if ip == nil {
		return nil
	}

	switch true {
	case ip[0] < 0x80:
		return classAMask
	case ip[0] < 0xC0:
		return classBMask
	default:
		return classCMask
	}
}

func allFF(b []byte) bool {
	for _, c := range b {
		if c != 0xff {
			return false
		}
	}
	return true
}

func Mask2(ip net.IP, mask IPMask) net.IP {
	if len(mask) == 16 && len(ip) == 4 && allFF(mask[:12]) {
		mask = mask[12:]
	}

	n := len(ip)
	if n != len(mask) {
		return nil
	}
	out := make(net.IP, n)
	for i := 0; i < n; i++ {
		out[i] = ip[i] & mask[i]
	}
	return out
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	mask := addr.DefaultMask()
	mask2 := DefaultMask2(addr)

	network := addr.Mask(mask)
	network2 := Mask2(addr, mask2)
	ones, bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		" Default mask length is ", bits,
		"Leading ones count is ", ones,
		"Mask is (hex) ", mask.String(),
		" Network is ", network.String(),
		"Mask2 is (hex) ", mask2,
		" Network is ", network2.String())

	os.Exit(0)
}
