package main

import (
	"fmt"
	"net"
	"time"
)

// Check function checks the reachability of a given destination and port
func Check(destination string, port string) string {

	// Construct the full address using the destination and port
	address := destination + ":" + port

	// Set the timeout duration for the connection attempt
	timeout := time.Duration(5 * time.Second)

	// Attempt to establish a TCP connection with a timeout
	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable,\n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v", destination,
			conn.LocalAddr(),
			conn.RemoteAddr())
	}

	return status
}
