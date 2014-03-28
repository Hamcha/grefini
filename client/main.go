package main

import (
	"./irc"
)

func main() {
	var irc irc.Client
	irc.Username = "grefini"
	irc.Nickname = "grefini"
	irc.Altnick  = "grefani"
	irc.Realname = "Gustave le Grand"
	irc.Proxyaddr = "127.0.0.1:4313"
	irc.Channels = []string{"#brony.it","#testbass"}
	irc.Connect("irc.ponychat.net:6667")
}
