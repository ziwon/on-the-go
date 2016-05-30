package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

/*
╭─ luno@optimus   ~GOPATH/src/github.com/ziwon/on-the-go/go-network/ch03/GetHeadInfo     4134b0 master   ✓ 
╰─ ./GetHeadInfo www.naver.com:80
HTTP/1.1 200 OK
Server: nginx
Date: Wed, 27 Jan 2016 15:34:11 GMT
Content-Type: text/html; charset=UTF-8
Connection: close
Cache-Control: no-cache, no-store, must-revalidate
Pragma: no-cache
P3P: CP="CAO DSP CURa ADMa TAIa PSAa OUR LAW STP PHY ONL UNI PUR FIN COM NAV INT DEM STA PRE"
X-Frame-Options: SAMEORIGIN


╭─ luno@optimus   ~GOPATH/src/github.com/ziwon/on-the-go/go-network/ch03/GetHeadInfo     4134b0 master   ✓ 
╰─ ./GetHeadInfo www.google.com:80
HTTP/1.0 302 Found
Cache-Control: private
Content-Type: text/html; charset=UTF-8
Location: http://www.google.co.kr/?gfe_rd=cr&ei=qOOoVocBjsPwB4TQhLAE
Content-Length: 259
Date: Wed, 27 Jan 2016 15:35:03 GMT
Server: GFE/2.0


╭─ luno@optimus   ~GOPATH/src/github.com/ziwon/on-the-go/go-network/ch03/GetHeadInfo     4134b0 master   ✓ 
╰─ ./GetHeadInfo www.facebook.com:80
HTTP/1.1 302 Found
Location: http://31.13.68.35/unsupportedbrowser
Vary: Accept-Encoding
Content-Type: text/html
X-FB-Debug: cMFdDxUedhzM2usnrvY/caAaAmUxnCPnfwWYxzPUtWXK4sbDvDBZ28FxW9uwEuD02i/o3ZZuN8hpf7qYROYbvg==
Date: Wed, 27 Jan 2016 15:36:35 GMT
Connection: close
Content-Length: 0
*/

/*
func dialTCP(net string, laddr, raddr *TCPAddr, deadline time.Time) (*TCPConn, error) {
	fd, err := internetSocket(net, laddr, raddr, deadline, syscall.SOCK_STREAM, 0, "dial")

	// TCP has a rarely used mechanism called a 'simultaneous connection' in
	// which Dial("tcp", addr1, addr2) run on the machine at addr1 can
	// connect to a simultaneous Dial("tcp", addr2, addr1) run on the machine
	// at addr2, without either machine executing Listen.  If laddr == nil,
	// it means we want the kernel to pick an appropriate originating local
	// address.  Some Linux kernels cycle blindly through a fixed range of
	// local ports, regardless of destination port.  If a kernel happens to
	// pick local port 50001 as the source for a Dial("tcp", "", "localhost:50001"),
	// then the Dial will succeed, having simultaneously connected to itself.
	// This can only happen when we are letting the kernel pick a port (laddr == nil)
	// and when there is no listener for the destination address.
	// It's hard to argue this is anything other than a kernel bug.  If we
	// see this happen, rather than expose the buggy effect to users, we
	// close the fd and try again.  If it happens twice more, we relent and
	// use the result.  See also:
	//	https://golang.org/issue/2690
	//	http://stackoverflow.com/questions/4949858/
	//
	// The opposite can also happen: if we ask the kernel to pick an appropriate
	// originating local address, sometimes it picks one that is already in use.
	// So if the error is EADDRNOTAVAIL, we have to try again too, just for
	// a different reason.
	//
	// The kernel socket code is no doubt enjoying watching us squirm.
	for i := 0; i < 2 && (laddr == nil || laddr.Port == 0) && (selfConnect(fd, err) || spuriousENOTAVAIL(err)); i++ {
		if err == nil {
			fd.Close()
		}
		fd, err = internetSocket(net, laddr, raddr, deadline, syscall.SOCK_STREAM, 0, "dial")
	}

	if err != nil {
		return nil, &OpError{Op: "dial", Net: net, Source: laddr.opAddr(), Addr: raddr.opAddr(), Err: err}
	}
	return newTCPConn(fd), nil
}
*/

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	//result, err := readFully(conn)
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
