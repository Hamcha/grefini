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

type ClientMessage struct {
	ServerId string
	Message  Message
	DateTime int64
}

const MESSAGE = "PRIVMSG"
const NOTICE  = "NOTICE"

var sock net.Conn
var modmap map[string][]func(string, Message)

func initmods() {
	initquote("ponychat")

	initmacro("ponychat")
	initmacro("azzurra")
	initmacro("espernet")

	initviaggi()
	initurl()
	initsed()
	initmarkov()

	inittime("espernet")
	
	fmt.Println("flotrshi-port - All modules loaded!")
}

func handle(sid string, msg Message) {
	mods, ok := modmap[sid]
	if !ok {
		return
	}

	for i := range mods {
		mods[i](sid, msg)
	}
}

func main () {
	// Init flotrshi-port modules
	initmods()

	// Make map
	modmap = make(map[string][]func(string, Message))
	modmap["ponychat"] = []func(string, Message){ quote, macro, ball, meta, viaggi, urldo, sed, mamma, hs, markov }
	modmap["espernet"] = []func(string, Message){ macro, urldo, sed, hs, markov, dotime, mock }
	modmap["azzurra"]  = []func(string, Message){ macro, meta, viaggi, urldo, sed, hs, markov }

	// Connect to the proxy
	var err error
	sock, err = net.Dial("tcp","127.0.0.1:8012")
	if err != nil { panic(err) }
	defer sock.Close()

	in := bufio.NewReader(sock)
	for {
		bytes, _, err := in.ReadLine()
		if err != nil { break }

		var msg ClientMessage
		err = json.Unmarshal(bytes, &msg)
		if err != nil { fmt.Printf("ERROR reading JSON: %s\r\n",err.Error()) }
		// Dispatch to flotrshi-port modules
		handle(msg.ServerId, msg.Message)
	}
}

func send(servid string, msg Message) {
	srvmsg := ClientMessage{
		ServerId: servid,
		Message:  msg,
		DateTime: 0,
	}
	bytes, _ := json.Marshal(srvmsg)
	fmt.Fprintln(sock,string(bytes)+"\r\n")
}
