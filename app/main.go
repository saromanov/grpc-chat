package main

import "net"

func main() {
	lis, err := net.Listen("tcp", ":14200")
	if err != nil {
		panic("unable to listen: %v")
	}
	defer lis.Close()
}
