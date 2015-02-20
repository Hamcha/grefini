package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
	"math/rand"
	"strconv"
)

var quotes map[string][]string

func initquote(sid string) {
	if quotes == nil {
		quotes = make(map[string][]string)
	}
	bytes, _ := ioutil.ReadFile("quotes."+sid+".json")
	var quoteFile []string
	json.Unmarshal(bytes,&quoteFile) 
	quotes[sid] = quoteFile

	fmt.Println("Loaded quotes for '"+sid+"'")
}

func quote(sid string, msg Message) {
	if (msg.Command == MESSAGE) {
		parts := strings.Split(msg.Text, " ")
		if parts[0] == "!quote" {
			var n int
			// Get parameter if specified or just random otherwise
			if len(parts) < 2 {
				n = rand.Intn(len(quotes[sid]))
			} else {
				var err error
				n, err = strconv.Atoi(parts[1])
				if err != nil {
					send(sid, Message{
						Command:MESSAGE,
						Target:	msg.Target,
						Text:	"FUK U "+msg.Source.Nickname,
					})
					return
				}
				n = n-1
				if n >= len(quotes[sid]) || n < 0 {
					send(sid, Message{
						Command:MESSAGE,
						Target: msg.Target,
						Text:	"Quote inesistente!",
					})
					return
				}
			}
			send(sid, Message{
				Command:MESSAGE,
				Target: msg.Target,
				Text:	"Quote #"+strconv.Itoa(n+1)+": "+quotes[sid][n],
			})
			return
		}
		if parts[0] == "!addquote" {
			if len(parts) < 2 {
				send(sid, Message{
					Command:MESSAGE,
					Target:	msg.Target,
					Text:	"FUK U "+msg.Source.Nickname,
				})
			}
			quote := msg.Text[10:]
			quotes[sid] = append(quotes[sid],quote)
			qid := len(quotes[sid])
			send(sid, Message{
				Command:MESSAGE,
				Target:	msg.Target,
				Text:	"Quote #"+strconv.Itoa(qid)+" aggiunta",
			})
			savequote(sid)
			return
		}
		if parts[0] == "!search" {
			if len(parts) < 2 {
				send(sid, Message{
					Command:MESSAGE,
					Target: msg.Target,
					Text:   "FUK U "+msg.Source.Nickname,
				})
			}
			patt := strings.ToLower(msg.Text[8:])
			lst := make([]int,0)
			for i,v := range quotes[sid] {
				if strings.Index(strings.ToLower(v),patt) >= 0 {
					lst = append(lst,i+1)
				}
			}
			if len(lst) == 0 {
				send(sid, Message{
					Command:MESSAGE,
					Target: msg.Target,
					Text:   "Nessuna quote trovata :(",
				})
			} else if len(lst) == 1 {
				send(sid, Message{
					Command:MESSAGE,
					Target: msg.Target,
					Text:   "Quote #"+strconv.Itoa(lst[0])+": "+quotes[sid][lst[0]-1],
				})
			} else {
				out := make([]string,len(lst))
				for i := range lst {
					out[i] = strconv.Itoa(lst[i])
				}
				send(sid, Message{
					Command:MESSAGE,
					Target: msg.Target,
					Text:   "Trovato nelle quote: "+strings.Join(out,","),
				})
			}
		}
	}
}

func savequote(sid string) {
	bytes,err := json.Marshal(quotes[sid])
	if err != nil { 
		fmt.Printf("CAN'T SAVE QUOTES: %s\r\n",err.Error())
		return
	}
	ioutil.WriteFile("quotes."+sid+".json",bytes, 0777)
}
