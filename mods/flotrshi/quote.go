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
		if parts[0] == "!search" {
			if len(parts) < 2 {
				send(Message{
					Command:MESSAGE,
					Target: msg.Target,
					Text:   "FUK U "+msg.Source.Nickname,
				})
			}
			patt := strings.ToLower(msg.Text[8:])
			lst := make([]int,0)
			for i,v := range quotes {
				if strings.Index(strings.ToLower(v),patt) >= 0 {
					lst = append(lst,i+1)
				}
			}
			if len(lst) == 0 {
				send(Message{
					Command:MESSAGE,
					Target: msg.Target,
					Text:   "Nessuna quote trovata :(",
				})
			} else if len(lst) == 1 {
				send(Message{
					Command:MESSAGE,
					Target: msg.Target,
					Text:   "Quote #"+strconv.Itoa(lst[0])+": "+quotes[lst[0]+1],
				})
			} else {
				out := make([]string,len(lst))
				for i := range lst {
					out[i] = strconv.Itoa(lst[i])
				}
				send(Message{
					Command:MESSAGE,
					Target: msg.Target,
					Text:   "Trovato nelle quote: "+strings.Join(out,","),
				})
			}
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
