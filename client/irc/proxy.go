package irc

import (
	"net"
	"bufio"
	"fmt"
	"encoding/json"
)

func proxyConnect(client Client, channel chan int, psocket chan net.Conn) {
	// Try to connect to the proxy
	proxySock, err := net.Dial("tcp",client.Proxyaddr)
	// Panic if can't connect.. we need the proxy!
	if err != nil {
		panic("MISSING PROXY")
	} else {
		fmt.Printf("CONNECTED TO PROXY!\r\n")
	}
	psocket <- proxySock
	// Start receiving proxy messages
	proxyReceive(proxySock,client)
	channel <- 1
}

func proxyReceive(proxySock net.Conn, client Client) {
	// Setup Proxy Reader
	b := bufio.NewReader(proxySock)
	for {
		// Read line by line
		bytes, _, err := b.ReadLine()
		if err != nil { break }
		// Try to parse JSON message
		var msg Message
		err = json.Unmarshal(bytes, &msg)
		if err != nil { 
			fmt.Printf("CAN'T PARSE JSON: %s\r\n", err.Error())
			continue
		}
		// Prepare IRC message and send
		out := prepare(msg)
		fmt.Fprintf(client.Socket,out)
	}

}
