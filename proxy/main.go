package main

import (
	"net"
	"fmt"
	"bufio"
)

const proxyAddr = "127.0.0.1:4313"
const modAddr = "127.0.0.1:4314"

var mods []net.Conn

func main() {
	// Create server for bot to connect
	botsrv, err := net.Listen("tcp", proxyAddr)
	channel := make(chan int)
	if err != nil { panic(err) }

	// Block until bot has connected
	conn, err := botsrv.Accept()
	if err != nil {
		fmt.Printf("CAN'T ACCEPT BOT : %s\r\n", err.Error())
		return
	} else {
		fmt.Printf("BOT CONNECTED!\r\n")
	}
	go handleBot(conn,channel)

	// Create server for modules to connect
	modsrv, err := net.Listen("tcp", modAddr)
	if err != nil { panic(err) }

	// Accept loop for modules
	for {
		c, err := modsrv.Accept()
		if err != nil {
			fmt.Printf("CAN'T ACCEPT MOD : %s\r\n", err.Error())
		}
		mods = append(mods, c)
		go handleMod(c,conn)
	}
}

func handleBot(c net.Conn, channel chan int) {
	b := bufio.NewReader(c)
	for {
		bytes, _, err := b.ReadLine()
		if err != nil { panic("BOT DISCONNECTED") }
		broadcast(string(bytes))
	}
	channel <- 0
}

func handleMod(c net.Conn, bot net.Conn) {
	b := bufio.NewReader(c)
	defer c.Close()
	for {
		bytes, _, err := b.ReadLine()
		if err != nil {	break }
		fmt.Fprintf(bot, string(bytes)+"\r\n")
	}
	removeCon(c)
}

func broadcast(message string) {
	for _,c := range mods {
		_, err := fmt.Fprintf(c,message+"\r\n")
		if err != nil { removeCon(c) }
	}
}

func removeCon(c net.Conn) {
	for i,con := range mods {
		if c == con {
			mods = append(mods[:i],mods[i+1:]...)
		}
	}
}
