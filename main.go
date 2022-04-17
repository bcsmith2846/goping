package main

import (
	"fmt"
	"net"

)

func main() {
	proto := "udp"
	localhost, _ := net.ResolveUDPAddr("ip4", "127.0.0.1")
	fmt.Printf("Resolved to %s\n", localhost)
	recv, _ := net.ListenUDP("ip4:"+proto,localhost)
	fmt.Printf("Made socket %s\n", recv)

	buf := make([]byte, 1024)
	numRead, _, _ := recv.ReadFrom(buf)
	for i, byte := range buf[:numRead] {
		if i % 8 == 0{
			fmt.Printf("\n")
		}
		fmt.Printf("% 02X", byte)

	}
}
