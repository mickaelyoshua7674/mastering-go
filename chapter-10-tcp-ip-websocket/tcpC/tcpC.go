package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port")
		return
	}

	connect := arguments[1]
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(c, "%s\n", text)
		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}