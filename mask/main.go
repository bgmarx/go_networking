package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip=addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("invalid address")
		os.Exit(1)
	}

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()

	fmt.Println("address is", addr.String(),
		" Default mask length is", bits,
		" leading ones count is ", ones,
		"mask is (hex)", mask.String(),
		" network is ", network.String())
	os.Exit(0)

}
