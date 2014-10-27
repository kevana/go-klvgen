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
		log.Fatal("Unable to parse address provided.")
	}

	conn, err = net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Unable to open connection to address provided.")
	}
	/*
		err = conn.Close()
		if err != nil {
			log.Fatal("Unable to close connection ", err)
		}
	*/

	defer conn.Close()
	for {
		// Do things
		b := []byte("PACKET_CONTENTS\n")
		// Options for sending:
		// Write
		n, err := conn.Write(b)
		if err != nil || n != len(b) {
			log.Printf("Failed to send packet ", err.Error())
		}
		// WriteMsgUDP - Fails if already connected
		/*
			n, _, err := conn.WriteMsgUDP(b, nil, addr)
			if err != nil || n != len(b) {
				log.Fatal("Failed to send packet", err)
			}
		*/
		// WriteTo - alias to WriteToUDP
		// WriteToUDP - Fails if already connected
		/*
			n, err := net.UDPConn.WriteToUDP(b, addr)
			if err != nil || n != len(b) {
				log.Fatal("Failed to send packet", err)
			}
		*/

		time.Sleep(1000 * time.Millisecond)
	}

}
