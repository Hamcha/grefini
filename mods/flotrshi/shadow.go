package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

var counts map[string]int

func initcount() {
	bytes, _ := ioutil.ReadFile("counter.json")
	json.Unmarshal(bytes,&counts)

	fmt.Println("Counter module loaded! (!count)")
}

func count(msg Message) {
	if (msg.Command == MESSAGE) {
		parts := strings.Split(msg.Text, " ")
		if parts[0] == "!count" {
			if len(parts) < 2 {
				send(Message{
					Command:MESSAGE,
					Target:	msg.Target,
					Text:	"fuk u "+msg.Source.Nickname,
				})
				return
			}
			topic := msg.Text[7:]
			var out string
			if _, ok := counts[topic]; !ok {
				out = "Non ne ha mai parlato.. per ora.."
			} else {
				out = "Ne ha solo parlato "+strconv.Itoa(counts[topic])
				if counts[topic] == 1 { out += " volta" } else { out += " volte" }
			}

			send(Message{
				Command:MESSAGE,
				Target: msg.Target,
				Text:	msg.Source.Nickname+": "+out,
			})
			return
		}
		if parts[0] == "!inc" {
			if len(parts) < 2 {
				send(Message{
					Command:MESSAGE,
					Target:	msg.Target,
					Text:	"fuk u "+msg.Source.Nickname,
				})
			}
			name := msg.Text[5:]
			counts[name] += 1
			send(Message{
				Command:MESSAGE,
				Target:	msg.Target,
				Text:	"Ok!",
			})
			savecount()
			return
		}
		if parts[0] == "!countlist" {
			out := ""
			count := 0
			for key, val := range counts {
				out += key + " ("+ strconv.Itoa(val) +") / "
				count += 1
				if count >= 24 {
					count = 0
					send(Message{
						Command: NOTICE,
						Target:	 msg.Source.Nickname,
						Text:    ""+out,
					})
					out = ""
				}
			}
			send(Message{
				Command: NOTICE,
				Target:	 msg.Source.Nickname,
				Text:    out,
			})
		}
	}
}

func savecount() {
	bytes,err := json.Marshal(counts)
	if err != nil {
		fmt.Printf("CAN'T SAVE COUNT: %s\r\n",err.Error())
		return
	}
	ioutil.WriteFile("counter.json",bytes, 0777)
}

