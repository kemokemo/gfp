package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	freePort, err := getFreePort()
	if err != nil {
		fmt.Println("free port is not available, ", err)
		return 1
	}

	fmt.Printf("%v", freePort)
	return 0
}

// getFreePort returns the available port of the local environment.
func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
