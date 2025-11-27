package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	raddr, err := net.ResolveUDPAddr("udp", ":42069")
	if err != nil {
		log.Fatalf("failed creating a udp listener: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatalf("error creating a connection: %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		data, err := reader.ReadString(byte('\n'))
		if err != nil {
			log.Fatalf("error readingdata, or EOF, addr: %v", err)
		}
		_, err = conn.Write([]byte(data))
		if err != nil {
			log.Fatalf("error writing to conn: %v", err)
		}
	}
}
