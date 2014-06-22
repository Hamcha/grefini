package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
)

var macros map[string]string

func initmacro() {
	bytes, _ := ioutil.ReadFile("macros.json")
	json.Unmarshal(bytes,&macros)

	fmt.Println("Macro module loaded! (!macro)")
}

func macro(msg Message) {
	if (msg.Command == MESSAGE) {
		parts := strings.Split(msg.Text, " ")
		if parts[0] == "!macro" {
			if len(parts) < 2 {
				send(Message{
					Command:MESSAGE, 
					Target:	msg.Target,
					Text:	"fuk u "+msg.Source.Nickname,
				})
				return
			}
			var out string
			if macros[parts[1]] == "" {
				out = "Macro inesistente"
			} else {
				out = macros[parts[1]]
			}

			send(Message{
				Command:MESSAGE,
				Target: msg.Target,
				Text:	msg.Source.Nickname+": "+out,
			})
			return
		}
		if parts[0] == "!addmacro" {
			if len(parts) < 3 {
				send(Message{
					Command:MESSAGE,
					Target:	msg.Target,
					Text:	"fuk u "+msg.Source.Nickname,
				})
				return
			}
			name := parts[1]
			value := parts[2]
			macros[name] = value
			send(Message{
				Command:MESSAGE,
				Target:	msg.Target,
				Text:	"Macro '"+name+"' aggiunta!",
			})
			savemacros()
			return
		}
		if parts[0] == "!macrolist" {
			out := ""
			count := 0
			for key, _ := range macros {
				out += key + " / "
				count += 1
				if count >= 30 {
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

func savemacros() {
	bytes,err := json.Marshal(macros)
	if err != nil { 
		fmt.Printf("CAN'T SAVE MACROS: %s\r\n",err.Error())
		return
	}
	ioutil.WriteFile("macros.json",bytes, 0777)
}

