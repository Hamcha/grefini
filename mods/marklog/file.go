package main

import (
	"os"
	"fmt"
)

func doesExist(user string) bool {
	if _, err := os.Stat("logs/"+user); err == nil {
		return true
	}
	return false
}

func log(msg Message) {
	f, err := os.OpenFile("logs/"+msg.Source.Nickname, os.O_RDWR|os.O_APPEND|os.O_CREATE , 0777);
	if err != nil {
		fmt.Println("Can't save log: "+err.Error())
		return
	}
	f.WriteString(msg.Text+"\r\n")
	f.Close()
}

func openf(user string) *os.File {
	fi, _ := os.Open("logs/"+user)
	return fi
}

func closef(f *os.File) {
	f.Close()
}
