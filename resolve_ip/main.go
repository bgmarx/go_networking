package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage %s hostname\n", os.Args[0])
		fmt.Println("usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}

	name := os.Args[1]

	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("resolved address is ", addr.String())
	os.Exit(0)
}
