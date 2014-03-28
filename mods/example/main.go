package main

import (
	"net"
	"encoding/json"
	"bufio"
	"fmt"
)

type User struct {
	Nickname string
	Username string
	Host	 string
}

type Message struct {
	Source	 User
	Command	 string
	Target	 string
	Text	 string
}

func main () {
	// Connect to the proxy
	sock, err := net.Dial("tcp","127.0.0.1:4314")
	if err != nil { panic(err) }
	defer sock.Close()

	in := bufio.NewReader(sock)
	for {
		bytes, _, err := in.ReadLine()
		if err != nil { break }

		var msg Message
		err = json.Unmarshal(bytes,&msg)
		if err != nil { fmt.Printf("ERROR reading JSON: %s\r\n",err.Error()) }

		handle(sock,msg)
	}
}

func handle(sock net.Conn, msg Message) {
	if (msg.Command == "PRIVMSG") && (msg.Text == "ciao") {
		out := Message{
			Command:"PRIVMSG", 
			Target:msg.Target, 
			Text:"ciao a te "+msg.Source.Nickname,
		}
		bytes, _ := json.Marshal(out)
		fmt.Fprintf(sock,string(bytes)+"\r\n")
	}
}
