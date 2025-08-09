// Package client implements a simple TCP client that connects to a server
package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Start(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Fprintln(conn, input.Text())
	}
	return nil
}
