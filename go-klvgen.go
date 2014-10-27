package main

import (
	"flag"
	"log"
	"net"
	"time"
)

var (
	addr       *net.UDPAddr
	conn       *net.UDPConn
	addrString = "127.0.0.1:9000"
)

func main() {
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", addrString)
	if err != nil {
		log.Fatal("Unable to parse address provided: %s", err)
	}

	conn, err = net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Unable to open connection %s", err)
	}

	// Give us a target to sink 
	go func() {
		udpAddress, err := net.ResolveUDPAddr("udp",addrString)
		if err != nil {
			log.Fatal("Unable to resolve address %s", err)
		}
		recvConn ,err := net.ListenUDP("udp", udpAddress)
		defer recvConn.Close()
		if err != nil {
			log.Fatal("Unable to listen: %s", err)
		}

		var buf []byte = make([]byte, 1024)
		for {
			n, _, err := recvConn.ReadFromUDP(buf)
			if err != nil {
				log.Fatal("Unable to read: %s", err)
			}
			if true {
				log.Printf("Server: Received message: %s", string(buf[0:n]))
			}
			time.Sleep(1000*time.Millisecond)
		}
	}()

	defer conn.Close()
	for {

		b := []byte("PACKET_CONTENTS\n")

		n, err := conn.Write(b)
		if err != nil || n != len(b) {
			log.Printf("Client: Failed to send packet: %s", err)
		}

		time.Sleep(1000 * time.Millisecond)
	}

}
