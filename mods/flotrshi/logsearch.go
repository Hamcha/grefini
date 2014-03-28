package main

import (
	"io/ioutil"
	"strings"
	"strconv"
)

func logSearch(msg Message) {
	if (msg.Command == MESSAGE) {
		parts := strings.Split(msg.Text, " ")
		if parts[0] == "!find" {
			if len(parts) < 3 {
				send(Message{
					Command: MESSAGE,
					Target:  msg.Target,
					Text:    "fuk u "+msg.Source.Nickname,
				})
			}
			bytes, err := ioutil.ReadFile("logs/"+parts[1])
			if err != nil {
				send(Message{
					Command: NOTICE,
					Target:  msg.Source.Nickname,
					Text:    "Non ho log di quell'utente :(",
				})
			}
			strlog := string(bytes)
			lines  := strings.Split(strlog,"\n")
			i := 1
			for _,k := range lines {
				if i > 3 { break }
				if strings.Index(strings.ToLower(k),strings.ToLower(strings.Join(parts[2:]," "))) >= 0 {
					send(Message{
						Command: NOTICE,
						Target:  msg.Source.Nickname,
						Text:    strconv.Itoa(i)+". <"+parts[1]+"> "+k,
					})
					i += 1
				}
			}
			if i < 2 {
				send(Message{
					Command: NOTICE,
					Target:  msg.Source.Nickname,
					Text:    "Non ho trovato nulla :(",
				})
			}
		}
	}
}
