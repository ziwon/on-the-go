package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s network-type service\n",
			os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]

	// lines, err := queryCS(network, "127.0.0.1", service)
	// query(netdir+"/cs", net+"!"+host+"!"+service, 128)
	/*
		func query(filename, query string, bufSize int) (res []string, err error) {
			file, err := os.OpenFile(filename, os.O_RDWR, 0)
			if err != nil {
				return
			}
			defer file.Close()

			_, err = file.Seek(0, 0)
			if err != nil {
				return
			}
			_, err = file.WriteString(query)
			if err != nil {
				return
			}
			_, err = file.Seek(0, 0)
			if err != nil {
				return
			}
			buf := make([]byte, bufSize)
			for {
				n, _ := file.Read(buf)
				if n <= 0 {
					break
				}
				res = append(res, string(buf[:n]))
			}
			return
		}
	*/
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	fmt.Println("Service port ", port)
	os.Exit(0)
}
