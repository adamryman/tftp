package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net"
)

//import "github.com/adamryman/tftp"

func main() {
	pc, err := net.ListenPacket("udp", "0.0.0.0:5040")
	if err != nil {
		log.Fatal(err)
	}
	for {
		buffer := make([]byte, 1024)
		_, addr, err := pc.ReadFrom(buffer)
		if err != nil {
			log.Println(err)
			continue
		}
		go handlePacket(pc, buffer, addr)
	}
}

func handlePacket(pc net.PacketConn, data []byte, addr net.Addr) {
	fmt.Printf("%x\n", data)
	fmt.Println()
	fmt.Printf("%s", addr.String())
}
