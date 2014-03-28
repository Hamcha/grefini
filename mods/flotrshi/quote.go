package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
	"math/rand"
	"strconv"
)

var quotes []string

func initquote() {
	bytes, _ := ioutil.ReadFile("quotes.json")
	json.Unmarshal(bytes,&quotes)

	fmt.Println("Quote module loaded! (!quote)")
}

func quote(msg Message) {
	if (msg.Command == MESSAGE) {
		parts := strings.Split(msg.Text, " ")
		if parts[0] == "!quote" {
			var n int
			// Get parameter if specified or just random otherwise
			if len(parts) < 2 {
				n = rand.Intn(len(quotes))
			} else {
				var err error
				n, err = strconv.Atoi(parts[1])
				if err != nil {
					send(Message{
						Command:MESSAGE, 
						Target:	msg.Target,
						Text:	"FUK U "+msg.Source.Nickname,
					})
					return
				}
				n = n-1
				if n >= len(quotes) || n < 1 {
					send(Message{
						Command:MESSAGE,
						Target: msg.Target,
						Text:	"Quote inesistente!",
					})
					return
				}
			}
			send(Message{
				Command:MESSAGE,
				Target: msg.Target,
				Text:	"Quote #"+strconv.Itoa(n+1)+": "+quotes[n],
			})
			return
		}
		if parts[0] == "!addquote" {
			if len(parts) < 2 {
				send(Message{
					Command:MESSAGE,
					Target:	msg.Target,
					Text:	"FUK U "+msg.Source.Nickname,
				})
			}
			quote := msg.Text[9:]
			quotes = append(quotes,quote)
			qid := len(quotes)
			send(Message{
				Command:MESSAGE,
				Target:	msg.Target,
				Text:	"Quote #"+strconv.Itoa(qid)+" aggiunta",
			})
			savequote()
			return
		}
	}
}

func savequote() {
	bytes,err := json.Marshal(quotes)
	if err != nil { 
		fmt.Printf("CAN'T SAVE QUOTES: %s\r\n",err.Error())
		return
	}
	ioutil.WriteFile("quotes.json",bytes, 0777)
}
