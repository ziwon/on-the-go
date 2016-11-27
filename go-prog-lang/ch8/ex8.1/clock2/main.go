package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var port, tz string
	flag.StringVar(&port, "port", os.Getenv("PORT"), "Port")
	flag.StringVar(&tz, "TZ", os.Getenv("TZ"), "Timezone")
	flag.Parse()

	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Errorf("Unable to get the location for timezone: %s", tz)
	}

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, loc)
	}
}

func handleConn(c net.Conn, loc *time.Location) {
	defer c.Close()
	for {
		const form = "15:04:05\n"
		now := time.Now()
		_, err := io.WriteString(c, now.In(loc).Format(form))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
