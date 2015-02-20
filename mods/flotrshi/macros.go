package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
)

var macros map[string]map[string]string

func initmacro(sid string) {
	if macros == nil {
		macros = make(map[string]map[string]string)
	}
	bytes, _ := ioutil.ReadFile("macros."+sid+".json")
	var macrosFile map[string]string
	json.Unmarshal(bytes,&macrosFile)
	macros[sid] = macrosFile

	fmt.Println("Loaded macros for '"+sid+"'")
}

func macro(sid string, msg Message) {
	if (msg.Command == MESSAGE) {
		parts := strings.Split(msg.Text, " ")
		if parts[0] == "!macro" {
			if len(parts) < 2 {
				send(sid, Message{
					Command:MESSAGE, 
					Target:	msg.Target,
					Text:	"fuk u "+msg.Source.Nickname,
				})
				return
			}
			var out string
			if macros[sid][parts[1]] == "" {
				out = "Macro inesistente"
			} else {
				out = macros[sid][parts[1]]
			}

			send(sid, Message{
				Command:MESSAGE,
				Target: msg.Target,
				Text:	msg.Source.Nickname+": "+out,
			})
			return
		}
		if parts[0] == "!addmacro" {
			if len(parts) < 3 {
				send(sid, Message{
					Command:MESSAGE,
					Target:	msg.Target,
					Text:	"fuk u "+msg.Source.Nickname,
				})
				return
			}
			name := parts[1]
			value := parts[2]
			macros[sid][name] = value
			send(sid, Message{
				Command:MESSAGE,
				Target:	msg.Target,
				Text:	"Macro '"+name+"' aggiunta!",
			})
			savemacros(sid)
			return
		}
		if parts[0] == "!macrolist" {
			out := ""
			count := 0
			for key, _ := range macros[sid] {
				out += key + " / "
				count += 1
				if count >= 30 {
					count = 0
					send(sid, Message{
						Command: NOTICE,
						Target:	 msg.Source.Nickname,
						Text:    ""+out,
					})
					out = ""
				}
			}
			send(sid, Message{
				Command: NOTICE,
				Target:	 msg.Source.Nickname,
				Text:    out,
			})
		}
	}
}

func savemacros(sid string) {
	bytes,err := json.Marshal(macros[sid])
	if err != nil { 
		fmt.Printf("CAN'T SAVE MACROS: %s\r\n",err.Error())
		return
	}
	ioutil.WriteFile("macros."+sid+".json",bytes, 0777)
}

